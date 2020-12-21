// Copyright 2020 The Gitea Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package gitea

import (
	"fmt"
	"net/url"
	"strconv"
)

const defaultPageSize = 10
const maxPageSize = 50

// ListOptions options for using Gitea's API pagination
type ListOptions struct {
	Page     int
	PageSize int
}

func (o ListOptions) getURLQuery() url.Values {
	query := make(url.Values)
	query.Add("page", fmt.Sprintf("%d", o.Page))
	query.Add("limit", fmt.Sprintf("%d", o.PageSize))

	return query
}

func (o ListOptions) setDefaults() {
	if o.Page < 1 {
		o.Page = 1
	}

	if o.PageSize < 0 || o.PageSize > maxPageSize {
		o.PageSize = defaultPageSize
	}
}

// saveSetDefaults respect custom MaxResponseItems settings
func (o ListOptions) saveSetDefaults(c *Client) {
	if o.Page < 1 {
		o.Page = 1
	}

	// TODO: if minimum required gitea version is v1.13.0 return back error and drop "max := 10"
	max := 10 // set max to a low value, this should prevent pagination loops if max was set to low values
	conf, _, err := c.GetGlobalAPISettings()
	if err == nil {
		max = conf.MaxResponseItems
	}

	if o.PageSize < 0 || o.PageSize > max {
		o.PageSize = defaultPageSize
	}
}

// preparePaginatedResponse prepare Response for pagination functions
// if you use this function make sure the endpoint do accept at least the page option!
func (c *Client) preparePaginatedResponse(resp *Response, lo ListOptions, listLength int) error {
	if resp == nil {
		return nil
	}

	// just a safety check if pagination exists
	if listLength > lo.PageSize {
		return fmt.Errorf("API returned more items (%d) as pagination option has set (%d)", listLength, lo.PageSize)
	}

	resp.currentItem = listLength
	resp.maxItems = lo.PageSize
	resp.page = lo.Page

	return nil
}

// Next return true if endpoint support pagination & there is a next page to fetch
func (resp *Response) Next() bool {
	if resp.currentItem == 0 {
		return false
	}

	// check based on page & total amount via header
	total, _ := strconv.ParseInt(resp.Header.Get("x-total-count"), 10, 64)
	if total != 0 && resp.page != 0 {
		return resp.maxItems*resp.page < int(total)
	}

	// if no headers found use pagination options to calculate
	// use fallback witch return ture as long as a page is "full"
	return resp.currentItem == resp.maxItems
}
