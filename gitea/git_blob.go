// Copyright 2019 The Gitea Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package gitea

import (
	"bytes"
	"code.gitea.io/gitea/modules/git"
	"encoding/json"
)

// GitBlobResponse represents a git blob
type BlobResponse struct {
	Content  *BlobContentResponse `json:"content"`
	Encoding string `json:"encoding"`
	URL      string `json:"url"`
	SHA      string `json:"sha"`
	Size     int64  `json:"size"`
}

type BlobContentResponse struct {
	Blob *git.Blob
}

func (bc *BlobContentResponse) MarshalJSON() ([]byte, error) {
	reader, err := bc.Blob.DataAsync()
	if err != nil {
		return nil, err
	}
	defer reader.Close()
	buf := new(bytes.Buffer)
	_, err = buf.ReadFrom(reader)
	if err != nil {
		return nil, err
	}
	return json.Marshal(buf.Bytes())
}
