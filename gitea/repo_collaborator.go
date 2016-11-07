// Copyright 2016 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package gitea

import (
	"bytes"
	"encoding/json"
	"fmt"
)

type AddCollaboratorOption struct {
	Permission *string `json:"permission"`
}

func (c *Client) AddCollaborator(user, repo, collaborator string, opt AddCollaboratorOption) error {
	body, err := json.Marshal(&opt)
	if err != nil {
		return err
	}
	_, err = c.getResponse("PUT", fmt.Sprintf("/repos/%s/%s/collaborators/%s", user, repo, collaborator), nil, bytes.NewReader(body))
	return err
}
