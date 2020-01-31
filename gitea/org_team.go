// Copyright 2019 The Gitea Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package gitea

import (
	"bytes"
	"encoding/json"
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
	// example: ["repo.code","repo.issues","repo.ext_issues","repo.wiki","repo.pulls","repo.releases","repo.ext_wiki"]
	Units []string `json:"units"`
}

// ListOrgTeamsOptions options for listing organization's teams
type ListOrgTeamsOptions struct {
	ListOptions
	Org string
}

// ListOrgTeams lists all teams of an organization
func (c *Client) ListOrgTeams(options ListOrgTeamsOptions) ([]*Team, error) {
	teams := make([]*Team, 0, options.getPageSize())
	return teams, c.getParsedResponse("GET", fmt.Sprintf("/orgs/%s/teams?%s", options.Org, options.getURLQuery().Encode()), nil, nil, &teams)
}

// ListMyTeamsOptions options for current's user teams
type ListMyTeamsOptions struct {
	ListOptions
}

// ListMyTeams lists all the teams of the current user
func (c *Client) ListMyTeams(options *ListMyOrgsOptions) ([]*Team, error) {
	if options == nil {
		options = &ListMyOrgsOptions{}
	}

	teams := make([]*Team, 0, options.getPageSize())
	return teams, c.getParsedResponse("GET", fmt.Sprintf("/user/teams?%s", options.getURLQuery().Encode()), nil, nil, &teams)
}

// GetTeam gets a team by ID
func (c *Client) GetTeam(id int64) (*Team, error) {
	t := new(Team)
	return t, c.getParsedResponse("GET", fmt.Sprintf("/teams/%d", id), nil, nil, t)
}

// CreateTeamOption options for creating a team
type CreateTeamOption struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	// enum: read,write,admin
	Permission string `json:"permission"`
	// example: ["repo.code","repo.issues","repo.ext_issues","repo.wiki","repo.pulls","repo.releases","repo.ext_wiki"]
	Units []string `json:"units"`
}

// CreateTeam creates a team for an organization
func (c *Client) CreateTeam(org string, opt CreateTeamOption) (*Team, error) {
	body, err := json.Marshal(&opt)
	if err != nil {
		return nil, err
	}
	t := new(Team)
	return t, c.getParsedResponse("POST", fmt.Sprintf("/orgs/%s/teams", org), jsonHeader, bytes.NewReader(body), t)
}

// EditTeamOption options for editing a team
type EditTeamOption struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	// enum: read,write,admin
	Permission string `json:"permission"`
	// example: ["repo.code","repo.issues","repo.ext_issues","repo.wiki","repo.pulls","repo.releases","repo.ext_wiki"]
	Units []string `json:"units"`
}

// EditTeam edits a team of an organization
func (c *Client) EditTeam(id int64, opt EditTeamOption) error {
	body, err := json.Marshal(&opt)
	if err != nil {
		return err
	}
	_, err = c.getResponse("PATCH", fmt.Sprintf("/teams/%d", id), jsonHeader, bytes.NewReader(body))
	return err
}

// DeleteTeam deletes a team of an organization
func (c *Client) DeleteTeam(id int64) error {
	_, err := c.getResponse("DELETE", fmt.Sprintf("/teams/%d", id), nil, nil)
	return err
}

// ListTeamMembersOptions options for listing team's members
type ListTeamMembersOptions struct {
	ListOptions
	ID int64
}

// ListTeamMembers lists all members of a team
func (c *Client) ListTeamMembers(options ListTeamMembersOptions) ([]*User, error) {
	members := make([]*User, 0, options.getPageSize())
	return members, c.getParsedResponse("GET", fmt.Sprintf("/teams/%d/members?%s", options.ID, options.getURLQuery().Encode()), nil, nil, &members)
}

// GetTeamMember gets a member of a team
func (c *Client) GetTeamMember(id int64, user string) (*User, error) {
	m := new(User)
	return m, c.getParsedResponse("GET", fmt.Sprintf("/teams/%d/members/%s", id, user), nil, nil, m)
}

// AddTeamMember adds a member to a team
func (c *Client) AddTeamMember(id int64, user string) error {
	_, err := c.getResponse("PUT", fmt.Sprintf("/teams/%d/members/%s", id, user), nil, nil)
	return err
}

// RemoveTeamMember removes a member from a team
func (c *Client) RemoveTeamMember(id int64, user string) error {
	_, err := c.getResponse("DELETE", fmt.Sprintf("/teams/%d/members/%s", id, user), nil, nil)
	return err
}

// ListTeamRepositoriesOptions options for listing team's repositories
type ListTeamRepositoriesOptions struct {
	ListOptions
	ID int64
}

// ListTeamRepositories lists all repositories of a team
func (c *Client) ListTeamRepositories(options ListTeamRepositoriesOptions) ([]*Repository, error) {
	repos := make([]*Repository, 0, options.getPageSize())
	return repos, c.getParsedResponse("GET", fmt.Sprintf("/teams/%d/repos?%s", options.ID, options.getURLQuery().Encode()), nil, nil, &repos)
}

// AddTeamRepository adds a repository to a team
func (c *Client) AddTeamRepository(id int64, org, repo string) error {
	_, err := c.getResponse("PUT", fmt.Sprintf("/teams/%d/repos/%s/%s", id, org, repo), nil, nil)
	return err
}

// RemoveTeamRepository removes a repository from a team
func (c *Client) RemoveTeamRepository(id int64, org, repo string) error {
	_, err := c.getResponse("DELETE", fmt.Sprintf("/teams/%d/repos/%s/%s", id, org, repo), nil, nil)
	return err
}
