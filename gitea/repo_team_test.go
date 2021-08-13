// Copyright 2021 The Gitea Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package gitea

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRepoTeamManagement(t *testing.T) {
	log.Println("== TestRepoTeamManagement ==")
	c := newTestClient()

	// prepare for test
	clean, repo, err := createTestOrgRepo(t, c, "RepoTeamManagement")
	if err != nil {
		return
	}
	defer clean()
	if _, err = createTestOrgTeams(t, c, repo.Owner.UserName, "Admins", AccessModeAdmin, []RepoUnitType{RepoUnitCode, RepoUnitIssues, RepoUnitPulls, RepoUnitReleases}); err != nil {
		return
	}
	if _, err = createTestOrgTeams(t, c, repo.Owner.UserName, "CodeManager", AccessModeWrite, []RepoUnitType{RepoUnitCode}); err != nil {
		return
	}
	if _, err = createTestOrgTeams(t, c, repo.Owner.UserName, "IssueManager", AccessModeWrite, []RepoUnitType{RepoUnitIssues, RepoUnitPulls}); err != nil {
		return
	}

	// test
	teams, _, err := c.GetRepoTeams(repo.Owner.UserName, repo.Name)
	assert.NoError(t, err)
	if !assert.Len(t, teams, 1) {
		return
	}
	assert.EqualValues(t, AccessModeOwner, teams[0].Permission)

	team, _, err := c.CheckRepoTeam(repo.Owner.UserName, repo.Name, "Admins")
	assert.NoError(t, err)
	assert.Nil(t, team)

	resp, err := c.AddRepoTeam(repo.Owner.UserName, repo.Name, "Admins")
	assert.NoError(t, err)
	assert.EqualValues(t, 204, resp.StatusCode)
	resp, err = c.AddRepoTeam(repo.Owner.UserName, repo.Name, "CodeManager")
	assert.NoError(t, err)
	assert.EqualValues(t, 204, resp.StatusCode)
	resp, err = c.AddRepoTeam(repo.Owner.UserName, repo.Name, "IssueManager")
	assert.NoError(t, err)
	assert.EqualValues(t, 204, resp.StatusCode)

	team, _, err = c.CheckRepoTeam(repo.Owner.UserName, repo.Name, "Admins")
	assert.NoError(t, err)
	if assert.NotNil(t, team) {
		assert.EqualValues(t, "Admins", team.Name)
		assert.EqualValues(t, AccessModeAdmin, team.Permission)
	}

	teams, _, err = c.GetRepoTeams(repo.Owner.UserName, repo.Name)
	assert.NoError(t, err)
	assert.Len(t, teams, 4)

	resp, err = c.RemoveRepoTeam(repo.Owner.UserName, repo.Name, "IssueManager")
	assert.NoError(t, err)
	assert.EqualValues(t, 204, resp.StatusCode)

	team, _, err = c.CheckRepoTeam(repo.Owner.UserName, repo.Name, "IssueManager")
	assert.NoError(t, err)
	assert.Nil(t, team)
}
