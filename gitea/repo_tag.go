// Copyright 2019 The Gitea Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package gitea

import (
	"fmt"
)

// Tag represents a repository tag
type Tag struct {
	Name       string      `json:"name"`
	Message    string      `json:"message"`
	ID         string      `json:"id"`
	Commit     *CommitMeta `json:"commit"`
	ZipballURL string      `json:"zipball_url"`
	TarballURL string      `json:"tarball_url"`
}

// AnnotatedTag represents an annotated tag
type AnnotatedTag struct {
	Tag          string                     `json:"tag"`
	SHA          string                     `json:"sha"`
	URL          string                     `json:"url"`
	Message      string                     `json:"message"`
	Tagger       *CommitUser                `json:"tagger"`
	Object       *AnnotatedTagObject        `json:"object"`
	Verification *PayloadCommitVerification `json:"verification"`
}

// AnnotatedTagObject contains meta information of the tag object
type AnnotatedTagObject struct {
	Type string `json:"type"`
	URL  string `json:"url"`
	SHA  string `json:"sha"`
}

// CreateTagOption options when creating a tag
type CreateTagOption struct {
	// required: true
	TagName string `json:"tag_name" binding:"Required"`
	Message string `json:"message"`
	Target  string `json:"target"`
}

// ListRepoTagsOptions options for listing a repository's tags
type ListRepoTagsOptions struct {
	ListOptions
}

// ListRepoTags list all the branches of one repository
func (c *Client) ListRepoTags(user, repo string, opt ListRepoTagsOptions) ([]*Tag, *Response, error) {
	if err := escapeValidatePathSegments(&user, &repo); err != nil {
		return nil, nil, err
	}
	opt.setDefaults()
	tags := make([]*Tag, 0, opt.PageSize)
	resp, err := c.getParsedResponse("GET", fmt.Sprintf("/repos/%s/%s/tags?%s", user, repo, opt.getURLQuery().Encode()), nil, nil, &tags)
	return tags, resp, err
}

// DeleteTag deletes a tag from a repository, if no release refers to it
func (c *Client) DeleteTag(user, repo string, tag string) (*Response, error) {
	if err := escapeValidatePathSegments(&user, &repo, &tag); err != nil {
		return nil, err
	}
	if err := c.checkServerVersionGreaterThanOrEqual(version1_14_0); err != nil {
		return nil, err
	}
	_, resp, err := c.getResponse("DELETE",
		fmt.Sprintf("/repos/%s/%s/tags/%s", user, repo, tag),
		nil, nil)
	return resp, err
}
