// Copyright 2020 The Gitea Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package gitea

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRepoBranches(t *testing.T) {
	log.Println("== TestRepoBranches ==")
	c := newTestClient()
	repoName := "branches"

	repo := prepareBranchTest(t, c, repoName)
	if repo == nil {
		return
	}

	bl, _, err := c.ListRepoBranches(repo.Owner.UserName, repo.Name, ListRepoBranchesOptions{})
	assert.NoError(t, err)
	assert.Len(t, bl, 3)
	assert.EqualValues(t, "feature", bl[0].Name)
	assert.EqualValues(t, "main", bl[1].Name)
	assert.EqualValues(t, "update", bl[2].Name)

	b, _, err := c.GetRepoBranch(repo.Owner.UserName, repo.Name, "update")
	assert.NoError(t, err)
	assert.EqualValues(t, bl[2].Commit.ID, b.Commit.ID)
	assert.EqualValues(t, bl[2].Commit.Added, b.Commit.Added)

	s, _, err := c.DeleteRepoBranch(repo.Owner.UserName, repo.Name, "main")
	assert.NoError(t, err)
	assert.False(t, s)
	s, _, err = c.DeleteRepoBranch(repo.Owner.UserName, repo.Name, "feature")
	assert.NoError(t, err)
	assert.True(t, s)

	bl, _, err = c.ListRepoBranches(repo.Owner.UserName, repo.Name, ListRepoBranchesOptions{})
	assert.NoError(t, err)
	assert.Len(t, bl, 2)

	b, _, err = c.GetRepoBranch(repo.Owner.UserName, repo.Name, "feature")
	assert.Error(t, err)
	assert.Nil(t, b)

	bNew, _, err := c.CreateBranch(repo.Owner.UserName, repo.Name, CreateBranchOption{BranchName: "NewBranch"})
	assert.NoError(t, err)

	b, _, err = c.GetRepoBranch(repo.Owner.UserName, repo.Name, bNew.Name)
	assert.NoError(t, err)
	assert.EqualValues(t, bNew, b)
}

func TestRepoBranchProtection(t *testing.T) {
	log.Println("== TestRepoBranchProtection ==")
	c := newTestClient()
	repoName := "BranchProtection"

	repo := prepareBranchTest(t, c, repoName)
	if repo == nil {
		return
	}
	assert.NotNil(t, repo)

	// ListBranchProtections
	bpl, _, err := c.ListBranchProtections(repo.Owner.UserName, repo.Name, ListBranchProtectionsOptions{})
	assert.NoError(t, err)
	assert.Len(t, bpl, 0)

	// CreateBranchProtection
	bp, _, err := c.CreateBranchProtection(repo.Owner.UserName, repo.Name, CreateBranchProtectionOption{
		BranchName:              "main",
		EnablePush:              true,
		EnablePushWhitelist:     true,
		PushWhitelistUsernames:  []string{"test01"},
		EnableMergeWhitelist:    true,
		MergeWhitelistUsernames: []string{"test01"},
		BlockOnOutdatedBranch:   true,
	})
	assert.NoError(t, err)
	assert.EqualValues(t, "main", bp.BranchName)
	assert.EqualValues(t, false, bp.EnableStatusCheck)
	assert.EqualValues(t, true, bp.EnablePush)
	assert.EqualValues(t, true, bp.EnablePushWhitelist)
	assert.EqualValues(t, []string{"test01"}, bp.PushWhitelistUsernames)

	bp, _, err = c.CreateBranchProtection(repo.Owner.UserName, repo.Name, CreateBranchProtectionOption{
		BranchName:              "update",
		EnablePush:              false,
		EnableMergeWhitelist:    true,
		MergeWhitelistUsernames: []string{"test01"},
	})
	assert.NoError(t, err)
	assert.NotNil(t, bp)

	bpl, _, err = c.ListBranchProtections(repo.Owner.UserName, repo.Name, ListBranchProtectionsOptions{})
	assert.NoError(t, err)
	assert.Len(t, bpl, 2)

	// GetBranchProtection
	bp, _, err = c.GetBranchProtection(repo.Owner.UserName, repo.Name, bpl[0].BranchName)
	assert.NoError(t, err)
	assert.EqualValues(t, bpl[0], bp)

	// EditBranchProtection
	bp, _, err = c.EditBranchProtection(repo.Owner.UserName, repo.Name, bpl[0].BranchName, EditBranchProtectionOption{
		EnablePush:                  OptionalBool(false),
		EnablePushWhitelist:         OptionalBool(false),
		PushWhitelistUsernames:      nil,
		RequiredApprovals:           OptionalInt64(1),
		EnableApprovalsWhitelist:    OptionalBool(true),
		ApprovalsWhitelistUsernames: []string{"test01"},
	})
	assert.NoError(t, err)
	assert.NotEqual(t, bpl[0], bp)
	assert.EqualValues(t, bpl[0].BranchName, bp.BranchName)
	assert.EqualValues(t, bpl[0].EnableMergeWhitelist, bp.EnableMergeWhitelist)
	assert.EqualValues(t, bpl[0].Created, bp.Created)

	// DeleteBranchProtection
	_, err = c.DeleteBranchProtection(repo.Owner.UserName, repo.Name, bpl[1].BranchName)
	assert.NoError(t, err)
	bpl, _, err = c.ListBranchProtections(repo.Owner.UserName, repo.Name, ListBranchProtectionsOptions{})
	assert.NoError(t, err)
	assert.Len(t, bpl, 1)
}

func prepareBranchTest(t *testing.T, c *Client, repoName string) *Repository {
	origRepo, err := createTestRepo(t, repoName, c)
	if !assert.NoError(t, err) {
		return nil
	}

	mainLicense, _, err := c.GetContents(origRepo.Owner.UserName, origRepo.Name, "main", "README.md")
	if !assert.NoError(t, err) || !assert.NotNil(t, mainLicense) {
		return nil
	}

	updatedFile, _, err := c.UpdateFile(origRepo.Owner.UserName, origRepo.Name, "README.md", UpdateFileOptions{
		FileOptions: FileOptions{
			Message:       "update it",
			BranchName:    "main",
			NewBranchName: "update",
		},
		SHA:     mainLicense.SHA,
		Content: "Tk9USElORyBJUyBIRVJFIEFOWU1PUkUKSUYgWU9VIExJS0UgVE8gRklORCBTT01FVEhJTkcKV0FJVCBGT1IgVEhFIEZVVFVSRQo=",
	})
	if !assert.NoError(t, err) || !assert.NotNil(t, updatedFile) {
		return nil
	}

	newFile, _, err := c.CreateFile(origRepo.Owner.UserName, origRepo.Name, "WOW-file", CreateFileOptions{
		Content: "QSBuZXcgRmlsZQo=",
		FileOptions: FileOptions{
			Message:       "creat a new file",
			BranchName:    "main",
			NewBranchName: "feature",
		},
	})
	if !assert.NoError(t, err) || !assert.NotNil(t, newFile) {
		return nil
	}

	return origRepo
}
