// Copyright 2021 The Gitea Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package gitea

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func createTestOrgTeams(t *testing.T, c *Client, org, name string, accessMode AccessMode, units []RepoUnitType) (*Team, error) {
	team, _, e := c.CreateTeam(org, CreateTeamOption{
		Name:                    name,
		Description:             name + "'s team desc",
		Permission:              accessMode,
		CanCreateOrgRepo:        false,
		IncludesAllRepositories: false,
		Units:                   units,
	})
	assert.NoError(t, e)
	assert.NotNil(t, team)
	return team, e
}
