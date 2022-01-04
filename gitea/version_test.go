// Copyright 2020 The Gitea Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package gitea

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVersion(t *testing.T) {
	log.Printf("== TestVersion ==")
	c := newTestClient()
	rawVersion, _, err := c.ServerVersion()
	assert.NoError(t, err)
	assert.True(t, true, rawVersion != "")

	assert.NoError(t, c.checkServerVersionGreaterThanOrEqual(version1_11_0))
	assert.Error(t, c.CheckServerVersionConstraint("< 1.11.0"))

	c.serverVersion = version1_11_0
	assert.Error(t, c.checkServerVersionGreaterThanOrEqual(version1_15_0))
	c.ignoreVersion = true
	assert.NoError(t, c.checkServerVersionGreaterThanOrEqual(version1_15_0))

	c, err = NewClient(getGiteaURL(), newTestClientAuth(), SetGiteaVersion("1.12.123"))
	assert.NoError(t, err)
	assert.NoError(t, c.CheckServerVersionConstraint("=1.12.123"))
}
