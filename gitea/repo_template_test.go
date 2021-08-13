// Copyright 2021 The Gitea Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package gitea

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRepoFromTemplate(t *testing.T) {
	log.Println("== TestRepoFromTemplate ==")
	c := newTestClient()
	repo, err := createTestRepo(t, "TemplateRepo", c)
	assert.NoError(t, err)
	repo, _, err = c.EditRepo(repo.Owner.UserName, repo.Name, EditRepoOption{Template: OptionalBool(true)})
	assert.NoError(t, err)
	_, err = c.SetRepoTopics(repo.Owner.UserName, repo.Name, []string{"abc", "def", "ghi"})
	assert.NoError(t, err)

	newRepo, resp, err := c.CreateRepoFromTemplate(repo.Owner.UserName, repo.Name, CreateRepoFromTemplateOption{
		Owner:       repo.Owner.UserName,
		Name:        "repoFromTemplate",
		Description: "",
		Topics:      true,
		Labels:      true,
	})
	assert.NoError(t, err)
	assert.EqualValues(t, 201, resp.StatusCode)
	assert.False(t, newRepo.Template)

	labels, _, err := c.ListRepoLabels(repo.Owner.UserName, repo.Name, ListLabelsOptions{})
	assert.NoError(t, err)
	assert.Len(t, labels, 7)

	topics, _, _ := c.ListRepoTopics(repo.Owner.UserName, repo.Name, ListRepoTopicsOptions{})
	assert.EqualValues(t, []string{"abc", "def", "ghi"}, topics)

	_, err = c.DeleteRepo(repo.Owner.UserName, "repoFromTemplate")
	assert.NoError(t, err)
}
