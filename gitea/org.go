// Copyright 2015 The Gogs Authors. All rights reserved.
// Copyright 2019 The Gitea Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package gitea

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// Organization represents an organization
type Organization struct {
	ID          int64  `json:"id"`
	UserName    string `json:"username"`
	FullName    string `json:"full_name"`
	AvatarURL   string `json:"avatar_url"`
	Description string `json:"description"`
	Website     string `json:"website"`
	Location    string `json:"location"`
	Visibility  string `json:"visibility"`
}

// ListMyOrgsOptions options for listing current user's organizations
type ListMyOrgsOptions struct {
	ListOptions
}

// ListMyOrgs list all of current user's organizations
func (c *Client) ListMyOrgs(options *ListMyOrgsOptions) ([]*Organization, error) {
	if options == nil {
		options = &ListMyOrgsOptions{}
	}

	orgs := make([]*Organization, 0, options.getPerPage())
	return orgs, c.getParsedResponse("GET", fmt.Sprintf("/user/orgs?%s", options.getURLQuery()), nil, nil, &orgs)
}

// ListUserOrgsOptions options for listing an user's organizations
type ListUserOrgsOptions struct {
	ListOptions
	User string
}

// ListUserOrgs list all of some user's organizations
func (c *Client) ListUserOrgs(options ListUserOrgsOptions) ([]*Organization, error) {
	orgs := make([]*Organization, 0, options.getPerPage())
	return orgs, c.getParsedResponse("GET", fmt.Sprintf("/users/%s/orgs?%s", options.User, options.getURLQuery()), nil, nil, &orgs)
}

// GetOrg get one organization by name
func (c *Client) GetOrg(orgname string) (*Organization, error) {
	org := new(Organization)
	return org, c.getParsedResponse("GET", fmt.Sprintf("/orgs/%s", orgname), nil, nil, org)
}

// CreateOrgOption options for creating an organization
type CreateOrgOption struct {
	UserName    string `json:"username"`
	FullName    string `json:"full_name"`
	Description string `json:"description"`
	Website     string `json:"website"`
	Location    string `json:"location"`
	// possible values are `public` (default), `limited` or `private`
	// enum: public,limited,private
	Visibility string `json:"visibility"`
}

// CreateOrg creates an organization
func (c *Client) CreateOrg(opt CreateOrgOption) (*Organization, error) {
	body, err := json.Marshal(&opt)
	if err != nil {
		return nil, err
	}
	org := new(Organization)
	return org, c.getParsedResponse("POST", "/orgs", jsonHeader, bytes.NewReader(body), org)
}

// EditOrgOption options for editing an organization
type EditOrgOption struct {
	FullName    string `json:"full_name"`
	Description string `json:"description"`
	Website     string `json:"website"`
	Location    string `json:"location"`
	// possible values are `public`, `limited` or `private`
	// enum: public,limited,private
	Visibility string `json:"visibility"`
}

// EditOrg modify one organization via options
func (c *Client) EditOrg(orgname string, opt EditOrgOption) error {
	body, err := json.Marshal(&opt)
	if err != nil {
		return err
	}
	_, err = c.getResponse("PATCH", fmt.Sprintf("/orgs/%s", orgname), jsonHeader, bytes.NewReader(body))
	return err
}

// DeleteOrg deletes an organization
func (c *Client) DeleteOrg(orgname string) error {
	_, err := c.getResponse("DELETE", fmt.Sprintf("/orgs/%s", orgname), nil, nil)
	return err
}
