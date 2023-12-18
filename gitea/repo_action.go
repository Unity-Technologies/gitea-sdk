// Copyright 2023 The Gitea Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package gitea

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type CreateOrUpdateSecretOption struct {
	Description string `json:"description,omitempty"`
	Data        string `json:"data"`
}

// CreateOrUpdateRepoSecret create or update a secret value in a repository
func (c *Client) CreateOrUpdateRepoSecret(user, repo, secret string, opt CreateOrUpdateSecretOption) (*Response, error) {
	if err := escapeValidatePathSegments(&user, &repo); err != nil {
		return nil, err
	}
	body, err := json.Marshal(opt)
	if err != nil {
		return nil, err
	}
	status, resp, err := c.getStatusCode("PUT", fmt.Sprintf("/repos/%s/%s/actions/secrets/%s", user, repo, secret), jsonHeader, bytes.NewReader(body))
	if err != nil {
		return resp, err
	}
	if status == http.StatusCreated || status == http.StatusNoContent {
		return resp, nil
	}
	return resp, fmt.Errorf("unexpected Status: %d", status)
}

// DeleteRepoSecret detele a secret in a repository
func (c *Client) DeleteRepoSecret(user, repo, secret string) (*Response, error) {
	if err := escapeValidatePathSegments(&user, &repo); err != nil {
		return nil, err
	}
	_, resp, err := c.getResponse("DELETE", fmt.Sprintf("/repos/%s/%s/actions/secrets/%s", user, repo, secret), nil, nil)
	return resp, err
}
