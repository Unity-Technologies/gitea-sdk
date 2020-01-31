// Copyright 2020 The Gitea Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package gitea

import (
	"fmt"
	"net/url"
)

// ListOptions options for using Gitea's API pagination
type ListOptions struct {
	Page     int
	PageSize int
}

func (o ListOptions) getURLQuery() url.Values {
	o.setDefaults()

	query := make(url.Values)
	query.Add("page", fmt.Sprintf("%d", o.Page))
	query.Add("limit", fmt.Sprintf("%d", o.PageSize))

	return query
}

func (o ListOptions) setDefaults() {
	if o.Page < 1 {
		o.Page = 1
	}

	if o.PageSize < 0 || o.PageSize > 50 {
		o.PageSize = 10
	}
}
