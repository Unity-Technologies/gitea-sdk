// Copyright 2021 The Gitea Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package gitea

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRepoCollaborator(t *testing.T) {
	log.Println("== TestRepoCollaborator ==")
	c := newTestClient()

	repo, _ := createTestRepo(t, "RepoCollaborators", c)
	createTestUser(t, "ping", c)
	defer c.AdminDeleteUser("ping")

	collaborators, _, err := c.ListCollaborators(repo.Owner.UserName, repo.Name, ListCollaboratorsOptions{})
	assert.NoError(t, err)
	assert.Len(t, collaborators, 0)

	mode := AccessModeAdmin
	resp, err := c.AddCollaborator(repo.Owner.UserName, repo.Name, "ping", AddCollaboratorOption{Permission: &mode})
	assert.NoError(t, err)
	assert.EqualValues(t, 204, resp.StatusCode)

	reviewers, _, err := c.GetReviewers(repo.Owner.UserName, repo.Name)
	assert.NoError(t, err)
	if assert.Len(t, reviewers, 1) {
		assert.EqualValues(t, "ping", reviewers[0].UserName)
	}

	assignees, _, err := c.GetAssignees(repo.Owner.UserName, repo.Name)
	assert.NoError(t, err)
	if assert.Len(t, assignees, 2) {
		assert.EqualValues(t, "ping", assignees[0].UserName)
	}

	collaborators, _, err = c.ListCollaborators(repo.Owner.UserName, repo.Name, ListCollaboratorsOptions{})
	assert.NoError(t, err)
	assert.Len(t, collaborators, 1)

	resp, err = c.DeleteCollaborator(repo.Owner.UserName, repo.Name, "ping")
	assert.NoError(t, err)
	assert.EqualValues(t, 204, resp.StatusCode)

	collaborators, _, err = c.ListCollaborators(repo.Owner.UserName, repo.Name, ListCollaboratorsOptions{})
	assert.NoError(t, err)
	assert.Len(t, collaborators, 0)
}
