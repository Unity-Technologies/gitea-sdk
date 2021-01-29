// Copyright 2020 The Gitea Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package gitea

import (
	"fmt"
	"time"
)

// CronTask represents a Cron task
type CronTask struct {
	Name      string    `json:"name"`
	Schedule  string    `json:"schedule"`
	Next      time.Time `json:"next"`
	Prev      time.Time `json:"prev"`
	ExecTimes int64     `json:"exec_times"`
}

// ListCronTaskOptions list options for ListCronTasks
type ListCronTaskOptions struct {
	ListOptions
}

// ListCronTasks list available cron tasks
// response support Next()
func (c *Client) ListCronTasks(opt ListCronTaskOptions) ([]*CronTask, *Response, error) {
	if err := c.checkServerVersionGreaterThanOrEqual(version1_13_0); err != nil {
		return nil, nil, err
	}
	if err := opt.saveSetDefaults(c); err != nil {
		return nil, nil, err
	}
	ct := make([]*CronTask, 0, opt.PageSize)
	resp, err := c.getParsedResponse("GET", fmt.Sprintf("/admin/cron?%s", opt.getURLQuery().Encode()), jsonHeader, nil, &ct)
	if err = c.preparePaginatedResponse(resp, &opt.ListOptions, len(ct)); err != nil {
		return ct, resp, err
	}
	return ct, resp, err
}

// RunCronTasks run a cron task
func (c *Client) RunCronTasks(task string) (*Response, error) {
	if err := c.checkServerVersionGreaterThanOrEqual(version1_13_0); err != nil {
		return nil, err
	}
	_, resp, err := c.getResponse("POST", fmt.Sprintf("/admin/cron/%s", task), jsonHeader, nil)
	return resp, err
}
