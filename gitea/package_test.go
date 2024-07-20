// Copyright 2023 The Gitea Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package gitea

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// create an org with a single package for testing purposes
func createTestPackage(t *testing.T, c *Client) error {
	_, _ = c.DeletePackage("PackageOrg", "generic", "MyPackage", "v1")
	_, _ = c.DeleteOrg("PackageOrg")
	_, _, _ = c.CreateOrg(CreateOrgOption{Name: "PackageOrg"})

	client := &http.Client{
		Timeout: time.Second * 10,
	}

	reader := bytes.NewReader([]byte("Hello world!"))

	url := fmt.Sprintf("%s/api/packages/PackageOrg/generic/MyPackage/v1/file1.txt", os.Getenv("GITEA_SDK_TEST_URL"))
	req, err := http.NewRequest(http.MethodPut, url, reader)
	if err != nil {
		log.Println(err)
		return err
	}

	req.SetBasicAuth(os.Getenv("GITEA_SDK_TEST_USERNAME"), os.Getenv("GITEA_SDK_TEST_PASSWORD"))
	response, err := client.Do(req)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	return nil
}

func TestListPackages(t *testing.T) {
	log.Println("== TestListPackages ==")
	c := newTestClient()
	err := createTestPackage(t, c)
	assert.NoError(t, err)

	packagesList, _, err := c.ListPackages("PackageOrg", ListPackagesOptions{
		ListOptions{
			Page:     1,
			PageSize: 1000,
		},
	})
	assert.NoError(t, err)
	assert.Len(t, packagesList, 1)
}

func TestGetPackage(t *testing.T) {
	log.Println("== TestGetPackage ==")
	c := newTestClient()
	err := createTestPackage(t, c)
	assert.NoError(t, err)

	pkg, _, err := c.GetPackage("PackageOrg", "generic", "MyPackage", "v1")
	assert.NoError(t, err)
	assert.NotNil(t, pkg)
	assert.True(t, pkg.Name == "MyPackage")
	assert.True(t, pkg.Version == "v1")
	assert.NotEmpty(t, pkg.CreatedAt)
}

func TestDeletePackage(t *testing.T) {
	log.Println("== TestDeletePackage ==")
	c := newTestClient()
	err := createTestPackage(t, c)
	assert.NoError(t, err)

	_, err = c.DeletePackage("PackageOrg", "generic", "MyPackage", "v1")
	assert.NoError(t, err)

	// no packages should be listed following deletion
	packagesList, _, err := c.ListPackages("PackageOrg", ListPackagesOptions{
		ListOptions{
			Page:     1,
			PageSize: 1000,
		},
	})
	assert.NoError(t, err)
	assert.Len(t, packagesList, 0)
}

func TestListPackageFiles(t *testing.T) {
	log.Println("== TestListPackageFiles ==")
	c := newTestClient()
	err := createTestPackage(t, c)
	assert.NoError(t, err)

	packageFiles, _, err := c.ListPackageFiles("PackageOrg", "generic", "MyPackage", "v1")
	assert.NoError(t, err)
	assert.Len(t, packageFiles, 1)
	assert.True(t, packageFiles[0].Name == "file1.txt")
}
