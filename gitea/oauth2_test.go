// Copyright 2020 The Gitea Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package gitea

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOauth2(t *testing.T) {
	log.Println("== TestOauth2Application ==")
	c := newTestClient()

	user := createTestUser(t, "oauth2_user", c)
	c.SetSudo(user.UserName)

	newApp, _, err := c.CreateOauth2(CreateOauth2Option{Name: "test", RedirectURIs: []string{"http://test/test"}})
	assert.NoError(t, err)
	assert.NotNil(t, newApp)
	assert.EqualValues(t, "test", newApp.Name)

	a, _, err := c.ListOauth2(ListOauth2Option{})
	assert.NoError(t, err)
	assert.Len(t, a, 1)
	assert.EqualValues(t, newApp.Name, a[0].Name)

	b, _, err := c.GetOauth2(newApp.ID)
	assert.NoError(t, err)
	assert.EqualValues(t, newApp.Name, b.Name)

	b, _, err = c.UpdateOauth2(newApp.ID, CreateOauth2Option{Name: newApp.Name, RedirectURIs: []string{"https://test/login"}})
	assert.NoError(t, err)
	assert.EqualValues(t, newApp.Name, b.Name)
	assert.EqualValues(t, "https://test/login", b.RedirectURIs[0])
	assert.EqualValues(t, newApp.ID, b.ID)
	assert.NotEqual(t, newApp.ClientSecret, b.ClientSecret)

	_, err = c.DeleteOauth2(newApp.ID)
	assert.NoError(t, err)
}
