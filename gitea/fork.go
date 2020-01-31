// Copyright 2016 The Gitea Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package gitea

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// ListForksOptions options for listing repository's forks
type ListForksOptions struct {
	ListOptions
	User string
	Repo string
}

// ListForks list a repository's forks
func (c *Client) ListForks(options ListForksOptions) ([]*Repository, error) {
	forks := make([]*Repository, options.getPageSize())
	err := c.getParsedResponse("GET",
		fmt.Sprintf("/repos/%s/%s/forks?%s", options.User, options.Repo, options.getURLQuery().Encode()),
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
	body, err := json.Marshal(form)
	if err != nil {
		return nil, err
	}
	fork := new(Repository)
	err = c.getParsedResponse("POST",
		fmt.Sprintf("/repos/%s/%s/forks", user, repo),
		jsonHeader, bytes.NewReader(body), &fork)
	return fork, err
}
