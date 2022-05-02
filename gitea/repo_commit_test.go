// Copyright 2020 The Gitea Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package gitea

import (
	"encoding/base64"
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

func TestGetCommitDiffOrPatch(t *testing.T) {
	log.Println("== TestGetCommitDiffOrPatch ==")
	c := newTestClient()

	repo, err := createTestRepo(t, "TestGetCommitDiffOrPatch", c)
	assert.NoError(t, err)

	// Add a new simple small commit to the repository.
	fileResponse, _, err := c.CreateFile(repo.Owner.UserName, repo.Name, "NOT_A_LICENSE", CreateFileOptions{
		Content: base64.StdEncoding.EncodeToString([]byte("But is it?\n")),
		FileOptions: FileOptions{
			Message: "Ensure people know it's not a license!",
			Committer: Identity{
				Name:  "Sup3rCookie",
				Email: "Sup3rCookie@example.com",
			},
		},
	})
	assert.NoError(t, err)

	// Test the diff output.
	diffOutput, _, err := c.GetCommitDiff(repo.Owner.UserName, repo.Name, fileResponse.Commit.SHA)
	assert.NoError(t, err)
	assert.EqualValues(t, "diff --git a/NOT_A_LICENSE b/NOT_A_LICENSE\nnew file mode 100644\nindex 0000000..f27a20a\n--- /dev/null\n+++ b/NOT_A_LICENSE\n@@ -0,0 +1 @@\n+But is it?\n", string(diffOutput))

	// Test the patch output.
	patchOutput, _, err := c.GetCommitPatch(repo.Owner.UserName, repo.Name, fileResponse.Commit.SHA)
	assert.NoError(t, err)
	// Use contains, because we cannot include the first part, because of dates + non-static CommitID..
	assert.Contains(t, string(patchOutput), "Subject: [PATCH] Ensure people know it's not a license!\n\n---\n NOT_A_LICENSE | 1 +\n 1 file changed, 1 insertion(+)\n create mode 100644 NOT_A_LICENSE\n\ndiff --git a/NOT_A_LICENSE b/NOT_A_LICENSE\nnew file mode 100644\nindex 0000000..f27a20a\n--- /dev/null\n+++ b/NOT_A_LICENSE\n@@ -0,0 +1 @@\n+But is it?\n")
}
