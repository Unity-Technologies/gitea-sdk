// Copyright 2021 The Gitea Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package gitea

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserSettings(t *testing.T) {
	log.Println("== TestUserSettings ==")
	c := newTestClient()

	userConf, _, err := c.GetUserSettings()
	assert.NoError(t, err)
	assert.NotNil(t, userConf)
	assert.EqualValues(t, UserSettings{
		Theme:        "auto",
		HideEmail:    false,
		HideActivity: false,
	}, *userConf)

	userConf, _, err = c.UpdateUserSettings(UserSettingsOptions{
		FullName:  OptionalString("Admin User on Test"),
		Language:  OptionalString("de_de"),
		HideEmail: OptionalBool(true),
	})
	assert.NoError(t, err)
	assert.NotNil(t, userConf)
	assert.EqualValues(t, UserSettings{
		FullName:     "Admin User on Test",
		Theme:        "auto",
		Language:     "de_de",
		HideEmail:    true,
		HideActivity: false,
	}, *userConf)

	_, _, err = c.UpdateUserSettings(UserSettingsOptions{
		FullName:  OptionalString(""),
		Language:  OptionalString(""),
		HideEmail: OptionalBool(false),
	})
	assert.NoError(t, err)
}
