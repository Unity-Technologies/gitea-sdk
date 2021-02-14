// Copyright 2020 The Gitea Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package gitea

import (
	"fmt"
)

// AdminListUnadoptedRepositories lists the unadopted repositories that match the provided names
func (c *Client) AdminListUnadoptedRepositories(opt ListOptions) ([]string, *Response, error) {
	if err := c.checkServerVersionGreaterThanOrEqual(version1_14_0); err != nil {
		return nil, nil, err
	}
	var repoNames []string
	resp, err := c.getParsedResponse("GET", fmt.Sprintf("/admin/unadopted?%s", opt.getURLQuery().Encode()), jsonHeader, nil, &repoNames)
	return repoNames, resp, err
}

// AdminAdoptRepository will adopt an unadopted repository
func (c *Client) AdminAdoptRepository(owner, repo string) (*Response, error) {
	if err := c.checkServerVersionGreaterThanOrEqual(version1_14_0); err != nil {
		return nil, err
	}
	_, resp, err := c.getResponse("POST", fmt.Sprintf("/admin/unadopted/%s/%s", owner, repo), jsonHeader, nil)
	return resp, err
}

// AdminDeleteUnadoptedRepository will delete an unadopted repository
func (c *Client) AdminDeleteUnadoptedRepository(owner, repo string) (*Response, error) {
	if err := c.checkServerVersionGreaterThanOrEqual(version1_14_0); err != nil {
		return nil, err
	}
	_, resp, err := c.getResponse("DELETE", fmt.Sprintf("/admin/unadopted/%s/%s", owner, repo), jsonHeader, nil)
	return resp, err
}
