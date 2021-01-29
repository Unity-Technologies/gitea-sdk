// Copyright 2020 The Gitea Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package gitea

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestLabels test label related func
func TestLabels(t *testing.T) {
	log.Println("== TestLabels ==")
	c := newTestClient()
	repo, err := createTestRepo(t, "LabelTestsRepo", c)
	assert.NoError(t, err)

	createOpts := CreateLabelOption{
		Name:        " ",
		Description: "",
		Color:       "",
	}
	err = createOpts.Validate()
	assert.Error(t, err)
	assert.EqualValues(t, "invalid color format", err.Error())
	createOpts.Color = "12345f"
	err = createOpts.Validate()
	assert.Error(t, err)
	assert.EqualValues(t, "empty name not allowed", err.Error())
	createOpts.Name = "label one"

	labelOne, _, err := c.CreateLabel(repo.Owner.UserName, repo.Name, createOpts)
	assert.NoError(t, err)
	assert.EqualValues(t, createOpts.Name, labelOne.Name)
	assert.EqualValues(t, createOpts.Color, labelOne.Color)

	labelTwo, _, err := c.CreateLabel(repo.Owner.UserName, repo.Name, CreateLabelOption{
		Name:        "blue",
		Color:       "#0000FF",
		Description: "CMYB(100%, 100%, 0%, 0%)",
	})
	assert.NoError(t, err)
	_, _, err = c.CreateLabel(repo.Owner.UserName, repo.Name, CreateLabelOption{
		Name:        "gray",
		Color:       "808080",
		Description: "CMYB(0%, 0%, 0%, 50%)",
	})
	assert.NoError(t, err)
	_, _, err = c.CreateLabel(repo.Owner.UserName, repo.Name, CreateLabelOption{
		Name:        "green",
		Color:       "#98F76C",
		Description: "CMYB(38%, 0%, 56%, 3%)",
	})
	assert.NoError(t, err)

	labels, resp, err := c.ListRepoLabels(repo.Owner.UserName, repo.Name, ListLabelsOptions{ListOptions: ListOptions{PageSize: 3}})
	assert.NoError(t, err)
	assert.Len(t, labels, 3)
	assert.True(t, resp.Next())
	assert.Contains(t, labels, labelTwo)
	assert.NotContains(t, labels, labelOne)
}
