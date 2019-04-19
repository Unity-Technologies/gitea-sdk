// Copyright 2016 The Gitea Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package gitea

import (
	"fmt"
)

// ListForks list a repository's forks
func (c *Client) ListForks(user, repo string) ([]*Repository, error) {
	forks := make([]*Repository, 10)
	err := c.getParsedResponse("GET",
		fmt.Sprintf("/repos/%s/%s/forks", user, repo),
		nil, nil, &forks)
	return forks, err
}

// CreateForkOption options for creating a fork
type CreateForkOption struct {
	// organization name, if forking into an organization
	Organization *string `json:"organization"`
}

// CreateFork create a fork of a repository
func (c *Client) CreateFork(user, repo string, form CreateForkOption) (*Repository, error) {
	fork := new(Repository)
	err := c.getParsedResponse("POST",
		fmt.Sprintf("/repos/%s/%s/forks", user, repo),
		jsonHeader, form, &fork)
	return fork, err
}
