// Copyright 2020 The Gitea Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package gitea

import (
	"fmt"
	"net/url"
	"time"
)

// NotificationThread expose Notification on API
type NotificationThread struct {
	ID         int64                `json:"id"`
	Repository *Repository          `json:"repository"`
	Subject    *NotificationSubject `json:"subject"`
	Unread     bool                 `json:"unread"`
	Pinned     bool                 `json:"pinned"`
	UpdatedAt  time.Time            `json:"updated_at"`
	URL        string               `json:"url"`
}

// NotificationSubject contains the notification subject (Issue/Pull/Commit)
type NotificationSubject struct {
	Title            string `json:"title"`
	URL              string `json:"url"`
	LatestCommentURL string `json:"latest_comment_url"`
	Type             string `json:"type" binding:"In(Issue,Pull,Commit)"`
}

// NotifyStatus notification status type
type NotifyStatus string

const (
	// NotifyStatusUnread was not read
	NotifyStatusUnread NotifyStatus = "unread"
	// NotifyStatusRead was already read by user
	NotifyStatusRead NotifyStatus = "read"
	// NotifyStatusPinned notification is pinned by user
	NotifyStatusPinned NotifyStatus = "pinned"
)

// ListNotificationOptions represents the filter options
type ListNotificationOptions struct {
	ListOptions
	Since  time.Time
	Before time.Time
	Status []NotifyStatus
}

// MarkNotificationOptions represents the filter & modify options
type MarkNotificationOptions struct {
	LastReadAt time.Time
	Status     []NotifyStatus
	ToStatus   NotifyStatus
}

// QueryEncode encode options to url query
func (opt *ListNotificationOptions) QueryEncode() string {
	query := opt.getURLQuery()
	if !opt.Since.IsZero() {
		query.Add("since", opt.Since.Format(time.RFC3339))
	}
	if !opt.Before.IsZero() {
		query.Add("before", opt.Before.Format(time.RFC3339))
	}
	for _, s := range opt.Status {
		query.Add("status-types", string(s))
	}
	return query.Encode()
}

// Validate the CreateUserOption struct
func (opt ListNotificationOptions) Validate(c *Client) error {
	if len(opt.Status) != 0 {
		return c.CheckServerVersionConstraint(">=1.12.3")
	}
	return nil
}

// QueryEncode encode options to url query
func (opt *MarkNotificationOptions) QueryEncode() string {
	query := make(url.Values)
	if !opt.LastReadAt.IsZero() {
		query.Add("last_read_at", opt.LastReadAt.Format(time.RFC3339))
	}
	for _, s := range opt.Status {
		query.Add("status-types", string(s))
	}

	if len(opt.ToStatus) != 0 {
		query.Add("to-status", string(opt.ToStatus))
	}
	return query.Encode()
}

// Validate the CreateUserOption struct
func (opt MarkNotificationOptions) Validate(c *Client) error {
	if len(opt.Status) != 0 || len(opt.ToStatus) != 0 {
		return c.CheckServerVersionConstraint(">=1.12.3")
	}
	return nil
}

// CheckNotifications list users's notification threads
func (c *Client) CheckNotifications() (int64, error) {
	if err := c.CheckServerVersionConstraint(">=1.12.0"); err != nil {
		return 0, err
	}
	new := struct {
		New int64 `json:"new"`
	}{}

	return new.New, c.getParsedResponse("GET", "/notifications/new", jsonHeader, nil, &new)
}

// GetNotification get notification thread by ID
func (c *Client) GetNotification(id int64) (*NotificationThread, error) {
	if err := c.CheckServerVersionConstraint(">=1.12.0"); err != nil {
		return nil, err
	}
	thread := new(NotificationThread)
	return thread, c.getParsedResponse("GET", fmt.Sprintf("/notifications/threads/%d", id), nil, nil, thread)
}

// ReadNotification mark notification thread as read by ID
func (c *Client) ReadNotification(id int64, status *NotifyStatus) error {
	if err := c.CheckServerVersionConstraint(">=1.12.0"); err != nil {
		return err
	}
	link := fmt.Sprintf("/notifications/threads/%d", id)
	if status != nil {
		link += fmt.Sprintf("?to-status=%s", *status)
	}
	_, err := c.getResponse("PATCH", link, nil, nil)
	return err
}

// ListNotifications list users's notification threads
func (c *Client) ListNotifications(opt ListNotificationOptions) ([]*NotificationThread, error) {
	if err := opt.Validate(c); err != nil {
		return nil, err
	}
	if err := c.CheckServerVersionConstraint(">=1.12.0"); err != nil {
		return nil, err
	}
	link, _ := url.Parse("/notifications")
	link.RawQuery = opt.QueryEncode()
	threads := make([]*NotificationThread, 0, 10)
	return threads, c.getParsedResponse("GET", link.String(), nil, nil, &threads)
}

// ReadNotifications mark notification threads as read
func (c *Client) ReadNotifications(opt MarkNotificationOptions) error {
	if err := opt.Validate(c); err != nil {
		return err
	}
	if err := c.CheckServerVersionConstraint(">=1.12.0"); err != nil {
		return err
	}
	link, _ := url.Parse("/notifications")
	link.RawQuery = opt.QueryEncode()
	_, err := c.getResponse("PUT", link.String(), nil, nil)
	return err
}

// ListRepoNotifications list users's notification threads on a specific repo
func (c *Client) ListRepoNotifications(owner, reponame string, opt ListNotificationOptions) ([]*NotificationThread, error) {
	if err := opt.Validate(c); err != nil {
		return nil, err
	}
	if err := c.CheckServerVersionConstraint(">=1.12.0"); err != nil {
		return nil, err
	}
	link, _ := url.Parse(fmt.Sprintf("/repos/%s/%s/notifications", owner, reponame))
	link.RawQuery = opt.QueryEncode()
	threads := make([]*NotificationThread, 0, 10)
	return threads, c.getParsedResponse("GET", link.String(), nil, nil, &threads)
}

// ReadRepoNotifications mark notification threads as read on a specific repo
func (c *Client) ReadRepoNotifications(owner, reponame string, opt MarkNotificationOptions) error {
	if err := opt.Validate(c); err != nil {
		return err
	}
	if err := c.CheckServerVersionConstraint(">=1.12.0"); err != nil {
		return err
	}
	link, _ := url.Parse(fmt.Sprintf("/repos/%s/%s/notifications", owner, reponame))
	link.RawQuery = opt.QueryEncode()
	_, err := c.getResponse("PUT", link.String(), nil, nil)
	return err
}
