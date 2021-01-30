// Copyright 2015 The Gogs Authors. All rights reserved.
// Copyright 2019 The Gitea Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package gitea

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/url"
)

// AdminListOrgsOptions options for listing admin's organizations
type AdminListOrgsOptions struct {
	ListOptions
}

var adminListOrgsLink, _ = url.Parse("/admin/orgs")

// AdminListOrgs lists all orgs
// response support Next()
func (c *Client) AdminListOrgs(opt AdminListOrgsOptions) ([]*Organization, *Response, error) {
	orgs := make([]*Organization, 0, mustPositive(opt.PageSize))
	resp, err := c.getParsedPaginatedResponse("GET", adminListOrgsLink, &opt, &orgs)
	if err = c.preparePaginatedResponse(resp, &opt.ListOptions, len(orgs)); err != nil {
		return orgs, resp, err
	}
	return orgs, resp, err
}

// AdminCreateOrg create an organization
func (c *Client) AdminCreateOrg(user string, opt CreateOrgOption) (*Organization, *Response, error) {
	body, err := json.Marshal(&opt)
	if err != nil {
		return nil, nil, err
	}
	org := new(Organization)
	resp, err := c.getParsedResponse("POST", fmt.Sprintf("/admin/users/%s/orgs", user), jsonHeader, bytes.NewReader(body), org)
	return org, resp, err
}
