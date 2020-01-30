// Copyright 2020 The Gitea Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package gitea

// ListOptions options for using Gitea's API pagination
type ListOptions struct {
	Page     int
	PageSize int
}
