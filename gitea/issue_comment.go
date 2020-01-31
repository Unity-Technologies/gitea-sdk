// Copyright 2016 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package gitea

import (
	"bytes"
	"encoding/json"
	"fmt"
	"time"
)

// Comment represents a comment on a commit or issue
type Comment struct {
	ID               int64     `json:"id"`
	HTMLURL          string    `json:"html_url"`
	PRURL            string    `json:"pull_request_url"`
	IssueURL         string    `json:"issue_url"`
	Poster           *User     `json:"user"`
	OriginalAuthor   string    `json:"original_author"`
	OriginalAuthorID int64     `json:"original_author_id"`
	Body             string    `json:"body"`
	Created          time.Time `json:"created_at"`
	Updated          time.Time `json:"updated_at"`
}

// ListIssueCommentsOptions options for listing issue's comments
type ListIssueCommentsOptions struct {
	ListOptions
	Owner string
	Repo  string
	Index int64
}

// ListIssueComments list comments on an issue.
func (c *Client) ListIssueComments(options ListIssueCommentsOptions) ([]*Comment, error) {
	comments := make([]*Comment, 0, options.getPageSize())
	return comments, c.getParsedResponse("GET", fmt.Sprintf("/repos/%s/%s/issues/%d/comments?%s", options.Owner, options.Repo, options.Index, options.getURLQuery().Encode()), nil, nil, &comments)
}

// ListRepoIssueCommentsOptions options for listing repository's issue's comments
type ListRepoIssueCommentsOptions struct {
	ListOptions
	Owner string
	Repo  string
}

// ListRepoIssueComments list comments for a given repo.
func (c *Client) ListRepoIssueComments(options ListRepoIssueCommentsOptions) ([]*Comment, error) {
	comments := make([]*Comment, 0, options.getPageSize())
	if err := c.CheckServerVersionConstraint(">=1.12.0"); err != nil {
		return comments, err
	}
	return comments, c.getParsedResponse("GET", fmt.Sprintf("/repos/%s/%s/issues/comments?%s", options.Owner, options.Repo, options.getURLQuery().Encode()), nil, nil, &comments)
}

// CreateIssueCommentOption options for creating a comment on an issue
type CreateIssueCommentOption struct {
	Body string `json:"body"`
}

// CreateIssueComment create comment on an issue.
func (c *Client) CreateIssueComment(owner, repo string, index int64, opt CreateIssueCommentOption) (*Comment, error) {
	body, err := json.Marshal(&opt)
	if err != nil {
		return nil, err
	}
	comment := new(Comment)
	return comment, c.getParsedResponse("POST", fmt.Sprintf("/repos/%s/%s/issues/%d/comments", owner, repo, index), jsonHeader, bytes.NewReader(body), comment)
}

// EditIssueCommentOption options for editing a comment
type EditIssueCommentOption struct {
	Body string `json:"body"`
}

// EditIssueComment edits an issue comment.
func (c *Client) EditIssueComment(owner, repo string, commentID int64, opt EditIssueCommentOption) (*Comment, error) {
	body, err := json.Marshal(&opt)
	if err != nil {
		return nil, err
	}
	comment := new(Comment)
	return comment, c.getParsedResponse("PATCH", fmt.Sprintf("/repos/%s/%s/issues/comments/%d", owner, repo, commentID), jsonHeader, bytes.NewReader(body), comment)
}

// DeleteIssueComment deletes an issue comment.
func (c *Client) DeleteIssueComment(owner, repo string, commentID int64) error {
	_, err := c.getResponse("DELETE", fmt.Sprintf("/repos/%s/%s/issues/comments/%d", owner, repo, commentID), nil, nil)
	return err
}
