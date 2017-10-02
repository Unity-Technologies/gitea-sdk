// Copyright 2015 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package gitea

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// CreateUserOption create user options
// swagger:parameters adminCreateUser
type CreateUserOption struct {
	// in: body
	SourceID int64 `json:"source_id"`
	// in: body
	LoginName string `json:"login_name"`
	// in: body
	Username string `json:"username" binding:"Required;AlphaDashDot;MaxSize(35)"`
	// in: body
	FullName string `json:"full_name" binding:"MaxSize(100)"`
	// in: body
	Email string `json:"email" binding:"Required;Email;MaxSize(254)"`
	// in: body
	Password string `json:"password" binding:"Required;MaxSize(255)"`
	// in: body
	SendNotify bool `json:"send_notify"`
}

// AdminCreateUser create a user
func (c *Client) AdminCreateUser(opt CreateUserOption) (*User, error) {
	body, err := json.Marshal(&opt)
	if err != nil {
		return nil, err
	}
	user := new(User)
	return user, c.getParsedResponse("POST", "/admin/users", jsonHeader, bytes.NewReader(body), user)
}

// EditUserOption edit user options
// swagger:parameters adminEditUser
type EditUserOption struct {
	// in: body
	SourceID int64 `json:"source_id"`
	// in: body
	LoginName string `json:"login_name"`
	// in: body
	FullName string `json:"full_name" binding:"MaxSize(100)"`
	// in: body
	Email string `json:"email" binding:"Required;Email;MaxSize(254)"`
	// in: body
	Password string `json:"password" binding:"MaxSize(255)"`
	// in: body
	Website string `json:"website" binding:"MaxSize(50)"`
	// in: body
	Location string `json:"location" binding:"MaxSize(50)"`
	// in: body
	Active *bool `json:"active"`
	// in: body
	Admin *bool `json:"admin"`
	// in: body
	AllowGitHook *bool `json:"allow_git_hook"`
	// in: body
	AllowImportLocal *bool `json:"allow_import_local"`
	// in: body
	MaxRepoCreation *int `json:"max_repo_creation"`
}

// AdminEditUser modify user informations
func (c *Client) AdminEditUser(user string, opt EditUserOption) error {
	body, err := json.Marshal(&opt)
	if err != nil {
		return err
	}
	_, err = c.getResponse("PATCH", fmt.Sprintf("/admin/users/%s", user), jsonHeader, bytes.NewReader(body))
	return err
}

// AdminDeleteUser delete one user according name
func (c *Client) AdminDeleteUser(user string) error {
	_, err := c.getResponse("DELETE", fmt.Sprintf("/admin/users/%s", user), nil, nil)
	return err
}

// AdminCreateUserPublicKey create one user with options
func (c *Client) AdminCreateUserPublicKey(user string, opt CreateKeyOption) (*PublicKey, error) {
	body, err := json.Marshal(&opt)
	if err != nil {
		return nil, err
	}
	key := new(PublicKey)
	return key, c.getParsedResponse("POST", fmt.Sprintf("/admin/users/%s/keys", user), jsonHeader, bytes.NewReader(body), key)
}
