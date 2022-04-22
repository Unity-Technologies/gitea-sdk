// Copyright 2020 The Gitea Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package gitea

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRepoTransfer(t *testing.T) {
	log.Printf("== TestRepoTransfer ==")
	c := newTestClient()

	org, _, err := c.AdminCreateOrg(c.username, CreateOrgOption{Name: "TransferOrg"})
	assert.NoError(t, err)
	repo, err := createTestRepo(t, "ToMove", c)
	assert.NoError(t, err)

	newRepo, _, err := c.TransferRepo(c.username, repo.Name, TransferRepoOption{NewOwner: org.UserName})
	assert.NoError(t, err) // admin transfer repository will execute immediately but not set as pendding.
	assert.NotNil(t, newRepo)
	assert.EqualValues(t, "ToMove", newRepo.Name)

	repo, err = createTestRepo(t, "ToMove", c)
	assert.NoError(t, err)
	_, resp, err := c.TransferRepo(c.username, repo.Name, TransferRepoOption{NewOwner: org.UserName})
	assert.EqualValues(t, 422, resp.StatusCode)
	assert.Error(t, err)

	_, err = c.DeleteRepo(repo.Owner.UserName, repo.Name)
	assert.NoError(t, err)
	_, err = c.DeleteRepo(newRepo.Owner.UserName, newRepo.Name)
	assert.NoError(t, err)
	_, err = c.DeleteOrg(org.UserName)
	assert.NoError(t, err)
}
