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

	type test struct {
		name               string
		confidentialClient *bool
	}
	boolTrue := true
	boolFalse := false

	testCases := []test{
		{"ConfidentialClient unset should fallback to false", nil},
		{"ConfidentialClient true", &boolTrue},
		{"ConfidentialClient false", &boolFalse},
	}

	for _, testCase := range testCases {
		createOptions := CreateOauth2Option{
			Name:         "test",
			RedirectURIs: []string{"http://test/test"},
		}
		if testCase.confidentialClient != nil {
			createOptions.ConfidentialClient = *testCase.confidentialClient
		}

		newApp, _, err := c.CreateOauth2(createOptions)
		assert.NoError(t, err, testCase.name)
		assert.NotNil(t, newApp, testCase.name)
		assert.EqualValues(t, "test", newApp.Name, testCase.name)
		if testCase.confidentialClient != nil {
			assert.EqualValues(t, *testCase.confidentialClient, newApp.ConfidentialClient, testCase.name)
		} else {
			assert.EqualValues(t, false, newApp.ConfidentialClient, testCase.name)
		}

		a, _, err := c.ListOauth2(ListOauth2Option{})
		assert.NoError(t, err, testCase.name)
		assert.Len(t, a, 1, testCase.name)
		assert.EqualValues(t, newApp.Name, a[0].Name, testCase.name)
		assert.EqualValues(t, newApp.ConfidentialClient, a[0].ConfidentialClient, testCase.name)

		b, _, err := c.GetOauth2(newApp.ID)
		assert.NoError(t, err, testCase.name)
		assert.EqualValues(t, newApp.Name, b.Name, testCase.name)
		assert.EqualValues(t, newApp.ConfidentialClient, b.ConfidentialClient, testCase.name)

		b, _, err = c.UpdateOauth2(newApp.ID, CreateOauth2Option{
			Name:               newApp.Name,
			ConfidentialClient: !newApp.ConfidentialClient,
			RedirectURIs:       []string{"https://test/login"},
		})
		assert.NoError(t, err, testCase.name)
		assert.EqualValues(t, newApp.Name, b.Name, testCase.name)
		assert.EqualValues(t, "https://test/login", b.RedirectURIs[0], testCase.name)
		assert.EqualValues(t, newApp.ID, b.ID, testCase.name)
		assert.NotEqual(t, newApp.ClientSecret, b.ClientSecret, testCase.name)
		assert.NotEqual(t, newApp.ConfidentialClient, b.ConfidentialClient, testCase.name)

		_, err = c.DeleteOauth2(newApp.ID)
		assert.NoError(t, err, testCase.name)
	}
}
