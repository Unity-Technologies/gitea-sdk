// Copyright 2015 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package gitea

import (
	"fmt"
)

// AdminListOrgs list all organizations
func (c *Client) AdminListOrgs() ([]*Organization, error) {
	orgs := make([]*Organization, 0, 5)
	return orgs, c.getParsedResponse("GET", "/admin/orgs", nil, nil, &orgs)
}

// AdminCreateOrg create an organization
func (c *Client) AdminCreateOrg(user string, opt CreateOrgOption) (*Organization, error) {
	org := new(Organization)
	return org, c.getParsedResponse("POST", fmt.Sprintf("/admin/users/%s/orgs", user),
		jsonHeader, opt, org)
}
