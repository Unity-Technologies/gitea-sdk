// Copyright 2020 The Gitea Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package gitea

import (
	"fmt"
)

// ListUnadoptedRepositoriesOptions options for listing unadopted repositories
type ListUnadoptedRepositoriesOptions struct {
	ListOptions
}

// ListUnadoptedRepositories lists the unadopted repositories that match the provided names
func (c *Client) ListUnadoptedRepositories(opt ListUnadoptedRepositoriesOptions) ([]string, *Response, error) {
	if err := c.checkServerVersionGreaterThanOrEqual(version1_14_0); err != nil {
		return nil, nil, err
	}
	var repoNames []string
	resp, err := c.getParsedResponse("GET", fmt.Sprintf("/admin/unadopted?%s", opt.getURLQuery().Encode()), jsonHeader, nil, &repoNames)
	return repoNames, resp, err
}

// AdoptRepository will adopt an unadopted repository
func (c *Client) AdoptRepository(owner, repo string) (*Response, error) {
	if err := c.checkServerVersionGreaterThanOrEqual(version1_14_0); err != nil {
		return nil, err
	}
	_, resp, err := c.getResponse("POST", fmt.Sprintf("/admin/unadopted/%s/%s", owner, repo), jsonHeader, nil)
	return resp, err
}

// DeleteUnadoptedRepository will delete an unadopted repository
func (c *Client) DeleteUnadoptedRepository(owner, repo string) (*Response, error) {
	if err := c.checkServerVersionGreaterThanOrEqual(version1_14_0); err != nil {
		return nil, err
	}
	_, resp, err := c.getResponse("DELETE", fmt.Sprintf("/admin/unadopted/%s/%s", owner, repo), jsonHeader, nil)
	return resp, err
}
