// Copyright 2020 The Gitea Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package gitea

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestListRepoCommits(t *testing.T) {
	log.Println("== TestListRepoCommits ==")
	c := newTestClient()

	repo, err := createTestRepo(t, "ListRepoCommits", c)
	assert.NoError(t, err)

	l, _, err := c.ListRepoCommits(repo.Owner.UserName, repo.Name, ListCommitOptions{})
	assert.NoError(t, err)
	assert.Len(t, l, 1)
	assert.EqualValues(t, "Initial commit\n", l[0].RepoCommit.Message)
	assert.EqualValues(t, "gpg.error.not_signed_commit", l[0].RepoCommit.Verification.Reason)
	assert.EqualValues(t, 100, l[0].Stats.Additions)
}
