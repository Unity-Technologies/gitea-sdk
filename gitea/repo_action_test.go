// Copyright 2024 The Gitea Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package gitea

import (
	"log"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateRepoActionSecret(t *testing.T) {
	log.Println("== TestCreateRepoActionSecret ==")
	c := newTestClient()

	user := createTestUser(t, "repo_action_user", c)
	c.SetSudo(user.UserName)
	newRepo, _, err := c.CreateRepo(CreateRepoOption{
		Name: "test",
	})
	assert.NoError(t, err)
	assert.NotNil(t, newRepo)

	// create secret
	resp, err := c.CreateRepoActionSecret(newRepo.Owner.UserName, newRepo.Name, CreateSecretOption{Name: "test", Data: "test"})
	assert.NoError(t, err)
	assert.Equal(t, http.StatusCreated, resp.StatusCode)

	// update secret
	resp, err = c.CreateRepoActionSecret(newRepo.Owner.UserName, newRepo.Name, CreateSecretOption{Name: "test", Data: "test2"})
	assert.NoError(t, err)
	assert.Equal(t, http.StatusNoContent, resp.StatusCode)

	// list secrets
	secrets, _, err := c.ListRepoActionSecret(newRepo.Owner.UserName, newRepo.Name, ListRepoActionSecretOption{})
	assert.NoError(t, err)
	assert.Len(t, secrets, 1)
}
