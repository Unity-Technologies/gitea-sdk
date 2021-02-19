// Copyright 2021 The Gitea Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package gitea

import (
	"log"
	"os/exec"
	"path"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdminAdopt(t *testing.T) {
	log.Println("== TestAdminAdopt ==")
	c := newTestClient()
	createTestUser(t, "adoptuser", c)
	createTestAdoptRepos(t)

	repos, _, err := c.AdminListUnadoptedRepositories(ListOptions{})
	assert.NoError(t, err)
	assert.Len(t, repos, 2)

	// cleanup after test
	/*
		_, err = c.DeleteRepo("adoptuser", "repo1")
		assert.NoError(t, err)
		_, err = c.DeleteRepo("adoptuser", "repo1")
		assert.NoError(t, err)
		_, err = c.AdminDeleteUser("adoptuser")

	*/
}

func createTestAdoptRepos(t *testing.T) {
	repoRoot := getRepoRoot()
	out, err := exec.Command("mkdir", "-p", path.Join(repoRoot, "adoptuser")).Output()
	assert.NoError(t, err)
	assert.EqualValues(t, "", out)

	repo1 := filepath.Join(repoRoot, "adoptuser", "repo1.git")
	_, err = exec.Command("git", "init", "--bare", repo1).Output()
	assert.NoError(t, err)

	repo2 := filepath.Join(repoRoot, "adoptuser", "repo2.git")
	_, err = exec.Command("git", "init", "--bare", repo2).Output()
	assert.NoError(t, err)

	/*
	  - git init --bare /tmp/data/adoptuser/repo1.git
	  - git init --bare /tmp/data/adoptuser/repo2.git
	*/
}
