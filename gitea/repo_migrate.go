// Copyright 2020 The Gitea Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package gitea

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// GitServiceType represents a git service
type GitServiceType string

const (
	// GitServicePlain represents a plain git service
	GitServicePlain GitServiceType = "git"
	//GitServiceGithub represents github.com
	GitServiceGithub GitServiceType = "github"
	// GitServiceGitea represents a gitea service
	GitServiceGitea GitServiceType = "gitea"
	// GitServiceGitlab represents a gitlab service
	GitServiceGitlab GitServiceType = "gitlab"
	// GitServiceGogs represents a gogs service
	GitServiceGogs GitServiceType = "gogs"
)

// MigrateRepoOption options for migrating a repository from an external service
type MigrateRepoOption struct {
	RepoName  string `json:"repo_name"`
	RepoOwner string `json:"repo_owner"`
	// deprecated use RepoOwner
	RepoOwnerID  int64          `json:"uid"`
	CloneAddr    string         `json:"clone_addr"`
	Service      GitServiceType `json:"service"`
	AuthUsername string         `json:"auth_username"`
	AuthPassword string         `json:"auth_password"`
	AuthToken    string         `json:"auth_token"`
	Mirror       bool           `json:"mirror"`
	Private      bool           `json:"private"`
	Description  string         `json:"description"`
	Wiki         bool           `json:"wiki"`
	Milestones   bool           `json:"milestones"`
	Labels       bool           `json:"labels"`
	Issues       bool           `json:"issues"`
	PullRequests bool           `json:"pull_requests"`
	Releases     bool           `json:"releases"`
}

// Validate the MigrateRepoOption struct
func (opt *MigrateRepoOption) Validate() error {
	// check user options
	if len(opt.CloneAddr) == 0 {
		return fmt.Errorf("CloneAddr required")
	}
	if len(opt.RepoName) == 0 {
		return fmt.Errorf("RepoName required")
	} else if len(opt.RepoName) > 100 {
		return fmt.Errorf("RepoName to long")
	}
	if len(opt.Description) > 255 {
		return fmt.Errorf("Description to long")
	}
	switch opt.Service {
	case GitServiceGithub:
		if len(opt.AuthToken) == 0 {
			return fmt.Errorf("github require token authentication")
		}
	}
	return nil
}

// MigrateRepo migrates a repository from other Git hosting sources for the authenticated user.
//
// To migrate a repository for a organization, the authenticated user must be a
// owner of the specified organization.
func (c *Client) MigrateRepo(opt MigrateRepoOption) (*Repository, error) {
	if err := opt.Validate(); err != nil {
		return nil, err
	}

	if err := c.CheckServerVersionConstraint(">=1.13.0"); err != nil {
		if len(opt.AuthToken) != 0 {
			// old gitea instances dont understand AuthToken
			opt.AuthUsername = opt.AuthToken
			opt.AuthPassword, opt.AuthToken = "", ""
		}
		if len(opt.RepoOwner) != 0 {
			u, err := c.GetUserInfo(opt.RepoOwner)
			if err != nil {
				return nil, err
			}
			opt.RepoOwnerID = u.ID
		} else if opt.RepoOwnerID == 0 {
			// fallback to current user as migration initiator
			u, err := c.GetMyUserInfo()
			if err != nil {
				return nil, err
			}
			opt.RepoOwnerID = u.ID
		}
	}

	body, err := json.Marshal(&opt)
	if err != nil {
		return nil, err
	}
	repo := new(Repository)
	return repo, c.getParsedResponse("POST", "/repos/migrate", jsonHeader, bytes.NewReader(body), repo)
}
