// Copyright 2020 The Gitea Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package gitea

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRepoWatch(t *testing.T) {
	log.Printf("== TestRepoWatch ==")
	c := newTestClient()
	rawVersion, _, err := c.ServerVersion()
	assert.NoError(t, err)
	assert.True(t, true, rawVersion != "")

	repo1, _ := createTestRepo(t, "TestRepoWatch_1", c)
	repo2, _ := createTestRepo(t, "TestRepoWatch_2", c)
	assert.NotEqual(t, repo1, repo2)

	// GetWatchedRepos
	wl, _, err := c.GetWatchedRepos("test01")
	assert.NoError(t, err)
	assert.NotNil(t, wl)
	maxcount := len(wl)

	// GetMyWatchedRepos
	wl, _, err = c.GetMyWatchedRepos()
	assert.NoError(t, err)
	assert.Len(t, wl, maxcount)

	// CheckRepoWatch
	isWatching, _, err := c.CheckRepoWatch(repo1.Owner.UserName, repo1.Name)
	assert.NoError(t, err)
	assert.True(t, isWatching)

	// UnWatchRepo
	_, err = c.UnWatchRepo(repo1.Owner.UserName, repo1.Name)
	assert.NoError(t, err)
	isWatching, _, _ = c.CheckRepoWatch(repo1.Owner.UserName, repo1.Name)
	assert.False(t, isWatching)

	// WatchRepo
	_, err = c.WatchRepo(repo1.Owner.UserName, repo1.Name)
	assert.NoError(t, err)
	isWatching, _, _ = c.CheckRepoWatch(repo1.Owner.UserName, repo1.Name)
	assert.True(t, isWatching)
}
