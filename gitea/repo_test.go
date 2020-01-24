// Copyright 2020 The Gitea Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package gitea

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateRepo(t *testing.T) {
	log.Printf("== TestCreateRepo ==")
	user, err := testClient.GetMyUserInfo()
	assert.NoError(t, err)

	var repoName = "test1"
	_, err = testClient.GetRepo(user.UserName, repoName)
	if err != nil {
		repo, err := testClient.CreateRepo(CreateRepoOption{
			Name: repoName,
		})
		assert.NoError(t, err)
		assert.NotNil(t, repo)
	}

	err = testClient.DeleteRepo(user.UserName, repoName)
	assert.NoError(t, err)
}
