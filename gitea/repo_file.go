// Copyright 2014 The Gogs Authors. All rights reserved.
// Copyright 2019 The Gitea Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package gitea

import (
	"fmt"
)

// GetFile downloads a file of repository, ref can be branch/tag/commit.
// e.g.: ref -> master, tree -> macaron.go(no leading slash)
func (c *Client) GetFile(user, repo, ref, tree string) ([]byte, error) {
	return c.getResponse("GET", fmt.Sprintf("/repos/%s/%s/raw/%s/%s", user, repo, ref, tree), nil, nil)
}

// FileOptions options for all file APIs
type FileOptions struct {
	Message       string   `json:"message" binding:"Required"`
	BranchName    string   `json:"branch"`
	NewBranchName string   `json:"new_branch"`
	Author        Identity `json:"author"`
	Committer     Identity `json:"committer"`
}

// CreateFileOptions options for creating files
type CreateFileOptions struct {
	FileOptions
	Content string `json:"content"`
}

// DeleteFileOptions options for deleting files (used for other File structs below)
type DeleteFileOptions struct {
	FileOptions
	SHA string `json:"sha" binding:"Required"`
}

// UpdateFileOptions options for updating files
type UpdateFileOptions struct {
	DeleteFileOptions
	Content  string `json:"content"`
	FromPath string `json:"from_path" binding:"MaxSize(500)"`
}

// FileLinksResponse contains the links for a repo's file
type FileLinksResponse struct {
	Self    string `json:"url"`
	GitURL  string `json:"git_url"`
	HTMLURL string `json:"html_url"`
}

// FileContentResponse contains information about a repo's file stats and content
type FileContentResponse struct {
	Name        string             `json:"name"`
	Path        string             `json:"path"`
	SHA         string             `json:"sha"`
	Size        int64              `json:"size"`
	URL         string             `json:"url"`
	HTMLURL     string             `json:"html_url"`
	GitURL      string             `json:"git_url"`
	DownloadURL string             `json:"download_url"`
	Type        string             `json:"type"`
	Links       *FileLinksResponse `json:"_links"`
}

// FileCommitResponse contains information generated from a Git commit for a repo's file.
type FileCommitResponse struct {
	CommitMeta
	HTMLURL   string        `json:"html_url"`
	Author    *CommitUser   `json:"author"`
	Committer *CommitUser   `json:"committer"`
	Parents   []*CommitMeta `json:"parents"`
	Message   string        `json:"message"`
	Tree      *CommitMeta   `json:"tree"`
}

// FileResponse contains information about a repo's file
type FileResponse struct {
	Content      *FileContentResponse       `json:"content"`
	Commit       *FileCommitResponse        `json:"commit"`
	Verification *PayloadCommitVerification `json:"verification"`
}

// FileDeleteResponse contains information about a repo's file that was deleted
type FileDeleteResponse struct {
	Content      interface{}                `json:"content"` // to be set to nil
	Commit       *FileCommitResponse        `json:"commit"`
	Verification *PayloadCommitVerification `json:"verification"`
}
