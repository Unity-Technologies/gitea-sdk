// Copyright 2021 The Gitea Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package gitea

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func createTestOrgRepo(t *testing.T, c *Client, name string) (func(), *Repository, error) {
	_, _, err := c.GetOrg(name)
	if err == nil {
		_, _ = c.DeleteOrg(name)
	}
	_, _, err = c.CreateOrg(CreateOrgOption{
		Name:                      name,
		Visibility:                VisibleTypePublic,
		RepoAdminChangeTeamAccess: true,
	})
	if !assert.NoError(t, err) {
		return nil, nil, err
	}

	_, _, err = c.GetRepo(name, name)
	if err == nil {
		_, _ = c.DeleteRepo(name, name)
	}

	repo, _, err := c.CreateOrgRepo(name, CreateRepoOption{
		Name:        name,
		Description: "A test Repo: " + name,
		AutoInit:    true,
		Gitignores:  "C,C++",
		License:     "MIT",
		Readme:      "Default",
		IssueLabels: "Default",
		Private:     false,
	})
	assert.NoError(t, err)
	assert.NotNil(t, repo)

	return func() {
		_, _ = c.DeleteRepo(name, name)
		_, _ = c.DeleteOrg(name)
	}, repo, err
}
