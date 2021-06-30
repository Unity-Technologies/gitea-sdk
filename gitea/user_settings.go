// Copyright 2021 The Gitea Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package gitea

// UserSettings represents user settings
type UserSettings struct {
	FullName      string `json:"full_name"`
	Website       string `json:"website"`
	Description   string `json:"description"`
	Location      string `json:"location"`
	Language      string `json:"language"`
	Theme         string `json:"theme"`
	DiffViewStyle string `json:"diff_view_style"`
	// Privacy
	HideEmail    bool `json:"hide_email"`
	HideActivity bool `json:"hide_activity"`
}

// UserSettingsOptions represents options to change user settings
type UserSettingsOptions struct {
	FullName      *string `json:"full_name" binding:"MaxSize(100)"`
	Website       *string `json:"website" binding:"OmitEmpty;ValidUrl;MaxSize(255)"`
	Description   *string `json:"description" binding:"MaxSize(255)"`
	Location      *string `json:"location" binding:"MaxSize(50)"`
	Language      *string `json:"language"`
	Theme         *string `json:"theme"`
	DiffViewStyle *string `json:"diff_view_style"`
	// Privacy
	HideEmail    *bool `json:"hide_email"`
	HideActivity *bool `json:"hide_activity"`
}
