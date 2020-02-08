// Copyright 2020 The Gitea Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package gitea

import (
	"fmt"
	"net/url"
)

// DeleteOrgMembership remove a member from an organization
func (c *Client) DeleteOrgMembership(org, user string) error {
	_, err := c.getResponse("DELETE", fmt.Sprintf("/orgs/%s/members/%s", url.PathEscape(org), url.PathEscape(user)), nil, nil)
	return err
}

// ListIssueOption list issue options
type ListOrgMembershipOption struct {
	ListOptions
}

// ListOrgMembership list an organization's members
func (c *Client) ListOrgMembership(org string, opt ListOrgMembershipOption) ([]*User, error) {
	opt.setDefaults()
	users := make([]*User, 0, opt.PageSize)

	link, _ := url.Parse(fmt.Sprintf("/orgs/%s/members", url.PathEscape(org)))
	link.RawQuery = opt.getURLQuery().Encode()
	return users, c.getParsedResponse("GET", link.String(), jsonHeader, nil, &users)
}
