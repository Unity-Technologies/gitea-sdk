// Copyright 2019 The Gitea Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package gitea

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// GitHook represents a Git repository hook
type GitHook struct {
	Name     string `json:"name"`
	IsActive bool   `json:"is_active"`
	Content  string `json:"content,omitempty"`
}

// ListRepoGitHooksOptions options for listing repository's githooks
type ListRepoGitHooksOptions struct {
	ListOptions
}

// ListRepoGitHooks list all the Git hooks of one repository
func (c *Client) ListRepoGitHooks(user, repo string, opt ListRepoGitHooksOptions) ([]*GitHook, *Response, error) {
	opt.setDefaults()
	gitHooks := make([]*GitHook, 0, opt.PageSize)
	resp, err := c.getParsedResponse("GET", fmt.Sprintf("/repos/%s/%s/hooks/git?%s", user, repo, opt.getURLQuery().Encode()), nil, nil, &gitHooks)
	if err != nil {
		return nil, nil, err
	}
	return gitHooks, resp, nil
}

// GetRepoGitHook get a Git hook of a repository
func (c *Client) GetRepoGitHook(user, repo, id string) (*GitHook, *Response, error) {
	gitHook := new(GitHook)
	resp, err := c.getParsedResponse("GET", fmt.Sprintf("/repos/%s/%s/hooks/git/%s", user, repo, id), nil, nil, gitHook)
	if err != nil {
		return nil, nil, err
	}
	return gitHook, resp, nil
}

// EditGitHookOption options when modifying one Git hook
type EditGitHookOption struct {
	Content string `json:"content"`
}

// EditRepoGitHook modify one Git hook of a repository
func (c *Client) EditRepoGitHook(user, repo, id string, opt EditGitHookOption) (*Response, error) {
	body, err := json.Marshal(&opt)
	if err != nil {
		return nil, err
	}
	_, resp, err := c.getResponse("PATCH", fmt.Sprintf("/repos/%s/%s/hooks/git/%s", user, repo, id), jsonHeader, bytes.NewReader(body))
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// DeleteRepoGitHook delete one Git hook from a repository
func (c *Client) DeleteRepoGitHook(user, repo, id string) (*Response, error) {
	_, resp, err := c.getResponse("DELETE", fmt.Sprintf("/repos/%s/%s/hooks/git/%s", user, repo, id), nil, nil)
	if err != nil {
		return nil, err
	}
	return resp, err
}
