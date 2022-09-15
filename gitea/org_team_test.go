// Copyright 2021 The Gitea Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package gitea

import (
	"log"
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

func TestTeamSearch(t *testing.T) {
	log.Println("== TestTeamSearch ==")
	c := newTestClient()

	orgName := "TestTeamsOrg"
	// prepare for test
	_, _, err := c.CreateOrg(CreateOrgOption{
		Name:                      orgName,
		Visibility:                VisibleTypePublic,
		RepoAdminChangeTeamAccess: true,
	})
	defer func() {
		_, _ = c.DeleteOrg(orgName)
	}()

	assert.NoError(t, err)

	if _, err = createTestOrgTeams(t, c, orgName, "Admins", AccessModeAdmin, []RepoUnitType{RepoUnitCode, RepoUnitIssues, RepoUnitPulls, RepoUnitReleases}); err != nil {
		return
	}

	teams, _, err := c.SearchOrgTeams(orgName, &SearchTeamsOptions{
		Query: "Admins",
	})
	assert.NoError(t, err)
	if assert.Len(t, teams, 1) {
		assert.Equal(t, "Admins", teams[0].Name)
	}
}
