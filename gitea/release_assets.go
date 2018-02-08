// Copyright 2018 The Gitea Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package gitea

import "time"

// Asset represent an attachment of issue/comment/release.
type Asset struct {
	ID                 int64  `json:"id"`
	UUID               string `json:"uuid"`
	URL                string `json:"url"`
	BrowserDownloadURL string `json:"browser_download_url"`
	Name               string `json:"name"`
	ContentType        string `json:"content_type"`
	Size               int64  `json:"size"`
	DownloadCount      int64  `json:"download_count"`
	// swagger:strfmt date-time `json:"url"`
	CreatedAt time.Time `json:"created_at"`
	// swagger:strfmt date-time `json:"url"`
	UpdatedAt time.Time `json:"updated_at"`
	Uploader  *User     `json:"uploader"`
}
