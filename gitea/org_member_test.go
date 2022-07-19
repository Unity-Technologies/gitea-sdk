// Copyright 2020 The Gitea Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package gitea

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

/**
// DeleteOrgMembership remove a member from an organization
func (c *Client) DeleteOrgMembership(org, user string) error {}


*/
func TestOrgMembership(t *testing.T) {
	log.Println("== TestOrgMembership ==")
	c := newTestClient()

	user := createTestUser(t, "org_mem_user", c)
	c.SetSudo(user.UserName)
	newOrg, _, err := c.CreateOrg(CreateOrgOption{Name: "MemberOrg"})
	assert.NoError(t, err)
	assert.NotNil(t, newOrg)

	// Check func
	check, _, err := c.CheckPublicOrgMembership(newOrg.UserName, user.UserName)
	assert.NoError(t, err)
	assert.False(t, check)
	check, _, err = c.CheckOrgMembership(newOrg.UserName, user.UserName)
	assert.NoError(t, err)
	assert.True(t, check)

	perm, _, err := c.GetOrgPermissions(newOrg.UserName, user.UserName)
	assert.NoError(t, err)
	assert.NotNil(t, perm)
	assert.True(t, perm.IsOwner)

	_, err = c.SetPublicOrgMembership(newOrg.UserName, user.UserName, true)
	assert.NoError(t, err)
	check, _, err = c.CheckPublicOrgMembership(newOrg.UserName, user.UserName)
	assert.NoError(t, err)
	assert.True(t, check)

	u, _, err := c.ListOrgMembership(newOrg.UserName, ListOrgMembershipOption{})
	assert.NoError(t, err)
	assert.Len(t, u, 1)
	assert.EqualValues(t, user.UserName, u[0].UserName)
	u, _, err = c.ListPublicOrgMembership(newOrg.UserName, ListOrgMembershipOption{})
	assert.NoError(t, err)
	assert.Len(t, u, 1)
	assert.EqualValues(t, user.UserName, u[0].UserName)

	_, err = c.DeleteOrgMembership(newOrg.UserName, user.UserName)
	assert.Error(t, err)

	c.sudo = ""
	_, err = c.AdminDeleteUser(user.UserName)
	assert.Error(t, err)
	_, err = c.DeleteOrg(newOrg.UserName)
	assert.NoError(t, err)
	_, err = c.AdminDeleteUser(user.UserName)
	assert.NoError(t, err)
}
