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
