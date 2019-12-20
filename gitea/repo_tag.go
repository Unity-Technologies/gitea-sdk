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
	ID         string      `json:"id"`
	Commit     *CommitMeta `json:"commit"`
	ZipballURL string      `json:"zipball_url"`
	TarballURL string      `json:"tarball_url"`
}

// ListRepoTagsOptions options for listing a repository's tags
type ListRepoTagsOptions struct {
	ListOptions
	User string
	Repo string
}

// ListRepoTags list all the branches of one repository
func (c *Client) ListRepoTags(options ListRepoTagsOptions) ([]*Tag, error) {
	tags := make([]*Tag, 0, options.getPerPage())
	return tags, c.getParsedResponse("GET", fmt.Sprintf("/repos/%s/%s/tags?%s", options.User, options.Repo, options.getURLQuery()), nil, nil, &tags)
}
