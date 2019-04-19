// Copyright 2015 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package gitea

import (
	"fmt"
)

// AdminCreateRepo create a repo
func (c *Client) AdminCreateRepo(user string, opt CreateRepoOption) (*Repository, error) {
	repo := new(Repository)
	return repo, c.getParsedResponse("POST", fmt.Sprintf("/admin/users/%s/repos", user),
		jsonHeader, opt, repo)
}
