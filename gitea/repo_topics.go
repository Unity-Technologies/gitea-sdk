// Copyright 2019 The Gitea Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package gitea

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// ListRepoTopicsOptions options for listing repo's topics
type ListRepoTopicsOptions struct {
	ListOptions
}

// topicsList represents a list of repo's topics
type topicsList struct {
	Topics []string `json:"topics"`
}

// ListRepoTopics list all repository's topics
// response support Next()
func (c *Client) ListRepoTopics(user, repo string, opt ListRepoTopicsOptions) ([]string, *Response, error) {
	if err := opt.saveSetDefaults(c); err != nil {
		return nil, nil, err
	}
	list := new(topicsList)
	resp, err := c.getParsedResponse("GET", fmt.Sprintf("/repos/%s/%s/topics?%s", user, repo, opt.getURLQuery().Encode()), nil, nil, list)
	if err != nil {
		return nil, resp, err
	}
	if err = c.preparePaginatedResponse(resp, &opt.ListOptions, len(list.Topics)); err != nil {
		return list.Topics, resp, err
	}
	return list.Topics, resp, nil
}

// SetRepoTopics replaces the list of repo's topics
func (c *Client) SetRepoTopics(user, repo string, list []string) (*Response, error) {

	l := topicsList{Topics: list}

	body, err := json.Marshal(&l)
	if err != nil {
		return nil, err
	}
	_, resp, err := c.getResponse("PUT", fmt.Sprintf("/repos/%s/%s/topics", user, repo), jsonHeader, bytes.NewReader(body))
	return resp, err
}

// AddRepoTopic adds a topic to a repo's topics list
func (c *Client) AddRepoTopic(user, repo, topic string) (*Response, error) {
	_, resp, err := c.getResponse("PUT", fmt.Sprintf("/repos/%s/%s/topics/%s", user, repo, topic), nil, nil)
	return resp, err
}

// DeleteRepoTopic deletes a topic from repo's topics list
func (c *Client) DeleteRepoTopic(user, repo, topic string) (*Response, error) {
	_, resp, err := c.getResponse("DELETE", fmt.Sprintf("/repos/%s/%s/topics/%s", user, repo, topic), nil, nil)
	return resp, err
}
