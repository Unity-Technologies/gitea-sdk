// Copyright 2016 The Gogs Authors. All rights reserved.
// Copyright 2018 The Gitea Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package gitea

import (
	"fmt"
)

// Team represents a team in an organization
type Team struct {
	ID           int64         `json:"id"`
	Name         string        `json:"name"`
	Description  string        `json:"description"`
	Organization *Organization `json:"organization"`
	// enum: none,read,write,admin,owner
	Permission string `json:"permission"`
	// enum: repo.code,repo.issues,repo.ext_issues,repo.wiki,repo.pulls,repo.releases,repo.ext_wiki
	Units []string `json:"units"`
}

// CreateTeamOption options for creating a team
type CreateTeamOption struct {
	// required: true
	Name        string `json:"name" binding:"Required;AlphaDashDot;MaxSize(30)"`
	Description string `json:"description" binding:"MaxSize(255)"`
	// enum: read,write,admin
	Permission string `json:"permission"`
	// enum: repo.code,repo.issues,repo.ext_issues,repo.wiki,repo.pulls,repo.releases,repo.ext_wiki
	Units []string `json:"units"`
}

// EditTeamOption options for editing a team
type EditTeamOption struct {
	// required: true
	Name        string `json:"name" binding:"Required;AlphaDashDot;MaxSize(30)"`
	Description string `json:"description" binding:"MaxSize(255)"`
	// enum: read,write,admin
	Permission string `json:"permission"`
	// enum: repo.code,repo.issues,repo.ext_issues,repo.wiki,repo.pulls,repo.releases,repo.ext_wiki
	Units []string `json:"units"`
}

// ListOrgTeams list all teams of an organization
func (c *Client) ListOrgTeams(orgname string) ([]*Team, error) {
	teams := make([]*Team, 0, 0)
	return teams, c.getParsedResponse("GET", fmt.Sprintf("/orgs/%s/teams", orgname), nil, nil, &teams)
}

// CreateTeam creates a new team
func (c *Client) CreateTeam(orgname string, opt CreateTeamOption) (*Team, error) {
	team := new(Team)
	return team, c.getParsedResponse("POST", fmt.Sprintf("/orgs/%s/teams", orgname), jsonHeader, opt, team)
}

// GetTeam gets a team by team ID
func (c *Client) GetTeam(teamID int64) (*Team, error) {
	team := new(Team)
	return team, c.getParsedResponse("GET", fmt.Sprintf("/teams/%d", teamID), nil, nil, team)
}

// DeleteTeam delete a team by team ID
func (c *Client) DeleteTeam(teamID int64) error {
	_, err := c.getResponse("DELETE", fmt.Sprintf("/teams/%d", teamID), nil, nil)
	return err
}

// EditTeam modify a team via options
func (c *Client) EditTeam(teamID int64, opt EditTeamOption) error {
	_, err := c.getResponse("PATCH", fmt.Sprintf("/teams/%d", teamID), jsonHeader, opt)
	return err
}

// ListTeamMembers list all members of a team
func (c *Client) ListTeamMembers(teamID int64) ([]*User, error) {
	users := make([]*User, 0, 0)
	return users, c.getParsedResponse("GET", fmt.Sprintf("/teams/%d/members", teamID), nil, nil, &users)
}

// AddTeamMember adds a member to a team
func (c *Client) AddTeamMember(teamID int64, username string) error {
	_, err := c.getResponse("PUT", fmt.Sprintf("/teams/%d/members/%s", teamID, username), nil, nil)
	return err
}

// RemoveTeamMember removes a member from a team
func (c *Client) RemoveTeamMember(teamID int64, username string) error {
	_, err := c.getResponse("DELETE", fmt.Sprintf("/teams/%d/members/%s", teamID, username), nil, nil)
	return err
}

// ListTeamRepos list all members of a team
func (c *Client) ListTeamRepos(teamID int64) ([]*Repository, error) {
	repos := make([]*Repository, 0, 0)
	return repos, c.getParsedResponse("GET", fmt.Sprintf("/teams/%d/repos", teamID), nil, nil, &repos)
}

// AddTeamRepo adds a repository to a team
func (c *Client) AddTeamRepo(teamID int64, org, repo string) error {
	_, err := c.getResponse("PUT", fmt.Sprintf("/teams/%d/repos/%s/%s", teamID, org, repo), nil, nil)
	return err
}

// RemoveTeamRepo removes a repository from a team
func (c *Client) RemoveTeamRepo(teamID int64, org, repo string) error {
	_, err := c.getResponse("DELETE", fmt.Sprintf("/teams/%d/repos/%s/%s", teamID, org, repo), nil, nil)
	return err
}
