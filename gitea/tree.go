// Copyright 2018 The Gitea Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package gitea

import (
	"fmt"
)

// TreeEntry represents the subtree structure within a tree.
type GitTreeEntry struct {
	Path		string 		`json:"path"`
	Mode		string 		`json:"mode"`
	Type		string 		`json:"type"`
	Size		int64 		`json:"size,omitempty"`
	SHA		string		`json:"sha"`
	URL		string		`json:"url"`
}

// Tree represents the tree structure.
type GitTreeResponse struct {
	SHA		string		`json:"sha"`
	URL		string		`json:"url"`
	Entries		[]GitTreeEntry	`json:"tree,omitempty"`
	Truncated 	bool		`json:"truncated"`
}

// GetTree gets information on a tree given the owner, repo, and sha or ref.
func (c *Client) GetTree(user string, repo string, tree string, recursive bool) (*GitTreeResponse, error) {
	t := new(GitTreeResponse)
	var Path = fmt.Sprintf("/repos/%s/%s/trees/%s", user, repo, tree)
	if recursive {
		Path += "?recursive=1"
	}
	err := c.getParsedResponse("GET", Path, nil, nil, t)
	return t, err
}
