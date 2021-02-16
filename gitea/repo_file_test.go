// Copyright 2020 The Gitea Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package gitea

import (
	"bytes"
	"encoding/base64"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFileCreateUpdateGet(t *testing.T) {
	log.Println("== TestFileCRUD ==")
	c := newTestClient()

	repo, err := createTestRepo(t, "ChangeFiles", c)
	assert.NoError(t, err)
	assert.NotNil(t, repo)

	raw, _, err := c.GetFile(repo.Owner.UserName, repo.Name, "master", "README.md")
	assert.NoError(t, err)
	assert.EqualValues(t, "IyBDaGFuZ2VGaWxlcwoKQSB0ZXN0IFJlcG86IENoYW5nZUZpbGVz", base64.StdEncoding.EncodeToString(raw))

	newFile, _, err := c.CreateFile(repo.Owner.UserName, repo.Name, "A", CreateFileOptions{
		FileOptions: FileOptions{
			Message: "create file A",
		},
		Content: "ZmlsZUEK",
	})
	assert.NoError(t, err)
	raw, _, _ = c.GetFile(repo.Owner.UserName, repo.Name, "master", "A")
	assert.EqualValues(t, "ZmlsZUEK", base64.StdEncoding.EncodeToString(raw))

	updatedFile, _, err := c.UpdateFile(repo.Owner.UserName, repo.Name, "A", UpdateFileOptions{
		FileOptions: FileOptions{
			Message: "add a new line",
		},
		SHA:     newFile.Content.SHA,
		Content: "ZmlsZUEKCmFuZCBhIG5ldyBsaW5lCg==",
	})
	assert.NoError(t, err)
	assert.NotNil(t, updatedFile)

	file, _, err := c.GetContents(repo.Owner.UserName, repo.Name, "master", "A")
	assert.NoError(t, err)
	assert.EqualValues(t, updatedFile.Content.SHA, file.SHA)
	assert.EqualValues(t, &updatedFile.Content.Content, &file.Content)

	_, err = c.DeleteFile(repo.Owner.UserName, repo.Name, "A", DeleteFileOptions{
		FileOptions: FileOptions{
			Message: "Delete File A",
		},
		SHA: updatedFile.Content.SHA,
	})
	assert.NoError(t, err)
	_, resp, err := c.GetFile(repo.Owner.UserName, repo.Name, "master", "A")
	assert.EqualValues(t, "404 Not Found", err.Error())
	assert.EqualValues(t, 404, resp.StatusCode )

	licence, _, err := c.GetContents(repo.Owner.UserName, repo.Name, "", "LICENSE")
	assert.NoError(t, err)
	licenceRaw, _, err := c.GetFile(repo.Owner.UserName, repo.Name, "", "LICENSE")
	updatedFile, _, err = c.UpdateFile(repo.Owner.UserName, repo.Name, "LICENSE", UpdateFileOptions{
		FileOptions: FileOptions{
			Message:       "Overwrite",
			BranchName:    "master",
			NewBranchName: "overwrite_licence",
		},
		SHA:     licence.SHA,
		Content: "Tk9USElORyBJUyBIRVJFIEFOWU1PUkUKSUYgWU9VIExJS0UgVE8gRklORCBTT01FVEhJTkcKV0FJVCBGT1IgVEhFIEZVVFVSRQo=",
	})
	assert.NoError(t, err)
	assert.NotNil(t, updatedFile)
	licenceRawNew, _, err := c.GetFile(repo.Owner.UserName, repo.Name, "overwrite_licence", "LICENSE")
	assert.NoError(t, err)
	assert.NotNil(t, licence)
	assert.False(t, bytes.Equal(licenceRaw, licenceRawNew))
}
