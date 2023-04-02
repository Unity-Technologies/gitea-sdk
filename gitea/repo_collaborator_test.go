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
	createTestUser(t, "pong", c)
	defer func() {
		_, err := c.AdminDeleteUser("ping")
		assert.NoError(t, err)
		_, err = c.AdminDeleteUser("pong")
		assert.NoError(t, err)
	}()

	collaborators, _, err := c.ListCollaborators(repo.Owner.UserName, repo.Name, ListCollaboratorsOptions{})
	assert.NoError(t, err)
	assert.Len(t, collaborators, 0)

	mode := AccessModeAdmin
	resp, err := c.AddCollaborator(repo.Owner.UserName, repo.Name, "ping", AddCollaboratorOption{Permission: &mode})
	assert.NoError(t, err)
	assert.EqualValues(t, 204, resp.StatusCode)

	permissonPing, resp, err := c.CollaboratorPermission(repo.Owner.UserName, repo.Name, "ping")
	assert.NoError(t, err)
	assert.EqualValues(t, 200, resp.StatusCode)
	assert.EqualValues(t, AccessModeAdmin, permissonPing.Permission)
	assert.EqualValues(t, "ping", permissonPing.User.UserName)

	mode = AccessModeRead
	_, err = c.AddCollaborator(repo.Owner.UserName, repo.Name, "pong", AddCollaboratorOption{Permission: &mode})
	assert.NoError(t, err)

	permissonPong, resp, err := c.CollaboratorPermission(repo.Owner.UserName, repo.Name, "pong")
	assert.NoError(t, err)
	assert.EqualValues(t, 200, resp.StatusCode)
	assert.EqualValues(t, AccessModeRead, permissonPong.Permission)
	assert.EqualValues(t, "pong", permissonPong.User.UserName)

	collaborators, _, err = c.ListCollaborators(repo.Owner.UserName, repo.Name, ListCollaboratorsOptions{})
	assert.NoError(t, err)
	assert.Len(t, collaborators, 2)
	assert.EqualValues(t, []string{"ping", "pong"}, userToStringSlice(collaborators))

	reviewers, _, err := c.GetReviewers(repo.Owner.UserName, repo.Name)
	assert.NoError(t, err)
	assert.Len(t, reviewers, 3)
	assert.EqualValues(t, []string{"ping", "pong", "test01"}, userToStringSlice(reviewers))

	assignees, _, err := c.GetAssignees(repo.Owner.UserName, repo.Name)
	assert.NoError(t, err)
	assert.Len(t, assignees, 2)
	assert.EqualValues(t, []string{"ping", "test01"}, userToStringSlice(assignees))

	resp, err = c.DeleteCollaborator(repo.Owner.UserName, repo.Name, "ping")
	assert.NoError(t, err)
	assert.EqualValues(t, 204, resp.StatusCode)

	collaborators, _, err = c.ListCollaborators(repo.Owner.UserName, repo.Name, ListCollaboratorsOptions{})
	assert.NoError(t, err)
	assert.Len(t, collaborators, 1)

	permissonNotExists, resp, err := c.CollaboratorPermission(repo.Owner.UserName, repo.Name, "user_that_not_exists")
	assert.Error(t, err)
	assert.EqualValues(t, 404, resp.StatusCode)
	assert.Nil(t, permissonNotExists)
}
