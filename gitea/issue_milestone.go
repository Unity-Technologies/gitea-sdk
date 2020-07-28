// Copyright 2016 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package gitea

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/url"
	"reflect"
	"strings"
	"time"
)

// Milestone milestone is a collection of issues on one repository
type Milestone struct {
	ID           int64      `json:"id"`
	Title        string     `json:"title"`
	Description  string     `json:"description"`
	State        StateType  `json:"state"`
	OpenIssues   int        `json:"open_issues"`
	ClosedIssues int        `json:"closed_issues"`
	Closed       *time.Time `json:"closed_at"`
	Deadline     *time.Time `json:"due_on"`
}

// ListMilestoneOption list milestone options
type ListMilestoneOption struct {
	ListOptions
	// open, closed, all
	State StateType
	Name  string
}

// QueryEncode turns options into querystring argument
func (opt *ListMilestoneOption) QueryEncode() string {
	query := opt.getURLQuery()
	if opt.State != "" {
		query.Add("state", string(opt.State))
	}
	if len(opt.Name) != 0 {
		query.Add("name", opt.Name)
	}
	return query.Encode()
}

// ListRepoMilestones list all the milestones of one repository
func (c *Client) ListRepoMilestones(owner, repo string, opt ListMilestoneOption) ([]*Milestone, error) {
	opt.setDefaults()
	milestones := make([]*Milestone, 0, opt.PageSize)

	link, _ := url.Parse(fmt.Sprintf("/repos/%s/%s/milestones", owner, repo))
	link.RawQuery = opt.QueryEncode()
	return milestones, c.getParsedResponse("GET", link.String(), nil, nil, &milestones)
}

// GetMilestone get one milestone by repo name and milestone id / name
func (c *Client) GetMilestone(owner, repo string, value interface{}) (*Milestone, error) {
	id, err := getMileIDbyStringOrInt64(c, owner, repo, value)
	if err != nil {
		return nil, err
	}
	milestone := new(Milestone)
	return milestone, c.getParsedResponse("GET", fmt.Sprintf("/repos/%s/%s/milestones/%d", owner, repo, id), nil, nil, milestone)
}

// CreateMilestoneOption options for creating a milestone
type CreateMilestoneOption struct {
	Title       string     `json:"title"`
	Description string     `json:"description"`
	State       StateType  `json:"state"`
	Deadline    *time.Time `json:"due_on"`
}

// Validate the CreateMilestoneOption struct
func (opt CreateMilestoneOption) Validate() error {
	if len(strings.TrimSpace(opt.Title)) == 0 {
		return fmt.Errorf("title is empty")
	}
	return nil
}

// CreateMilestone create one milestone with options
func (c *Client) CreateMilestone(owner, repo string, opt CreateMilestoneOption) (*Milestone, error) {
	if err := opt.Validate(); err != nil {
		return nil, err
	}
	body, err := json.Marshal(&opt)
	if err != nil {
		return nil, err
	}
	milestone := new(Milestone)
	err = c.getParsedResponse("POST", fmt.Sprintf("/repos/%s/%s/milestones", owner, repo), jsonHeader, bytes.NewReader(body), milestone)

	// make creating closed milestones need gitea >= v1.13.0
	// this make it backwards compatible
	if err == nil && opt.State == StateClosed && milestone.State != StateClosed {
		closed := StateClosed
		return c.EditMilestone(owner, repo, milestone.ID, EditMilestoneOption{
			State: &closed,
		})
	}

	return milestone, err
}

// EditMilestoneOption options for editing a milestone
type EditMilestoneOption struct {
	Title       string     `json:"title"`
	Description *string    `json:"description"`
	State       *StateType `json:"state"`
	Deadline    *time.Time `json:"due_on"`
}

// Validate the EditMilestoneOption struct
func (opt EditMilestoneOption) Validate() error {
	if len(opt.Title) != 0 && len(strings.TrimSpace(opt.Title)) == 0 {
		return fmt.Errorf("title is empty")
	}
	return nil
}

// EditMilestone modify milestone with options by milestone id / name
func (c *Client) EditMilestone(owner, repo string, value interface{}, opt EditMilestoneOption) (*Milestone, error) {
	id, err := getMileIDbyStringOrInt64(c, owner, repo, value)
	if err != nil {
		return nil, err
	}
	if err := opt.Validate(); err != nil {
		return nil, err
	}
	body, err := json.Marshal(&opt)
	if err != nil {
		return nil, err
	}
	milestone := new(Milestone)
	return milestone, c.getParsedResponse("PATCH", fmt.Sprintf("/repos/%s/%s/milestones/%d", owner, repo, id), jsonHeader, bytes.NewReader(body), milestone)
}

// DeleteMilestone delete one milestone by milestone id / name
func (c *Client) DeleteMilestone(owner, repo string, value interface{}) error {
	id, err := getMileIDbyStringOrInt64(c, owner, repo, value)
	if err != nil {
		return err
	}
	_, err = c.getResponse("DELETE", fmt.Sprintf("/repos/%s/%s/milestones/%d", owner, repo, id), nil, nil)
	return err
}

func getMileIDbyStringOrInt64(c *Client, owner, repo string, value interface{}) (int64, error) {
	vv := reflect.ValueOf(value)
	if vv.Kind() != reflect.String || vv.Kind() != reflect.Int64 {
		return 0, fmt.Errorf("only string and int64 supported")
	}
	if vv.Kind() == reflect.String {
		id, err := c.ResolveMileIDbyName(owner, repo, value.(string))
		if err != nil {
			return 0, err
		}
		return id, nil
	} else {
		return value.(int64), nil
	}
}

// ResolveMileIDbyName take a milestone name and return the id if it exist
func (c *Client) ResolveMileIDbyName(owner, repo, name string) (int64, error) {
	i := 0
	for {
		i++
		miles, err := c.ListRepoMilestones(owner, repo, ListMilestoneOption{
			ListOptions: ListOptions{
				Page: i,
			},
			State: "all",
			Name:  name,
		})
		if err != nil || len(miles) == 0 {
			return 0, err
		}
		for _, m := range miles {
			if strings.ToLower(strings.TrimSpace(m.Title)) == strings.ToLower(strings.TrimSpace(name)) {
				return m.ID, nil
			}
		}
	}
}
