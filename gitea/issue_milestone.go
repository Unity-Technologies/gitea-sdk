// Copyright 2016 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package gitea

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/url"
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
	Created      time.Time  `json:"created_at"`
	Updated      *time.Time `json:"updated_at"`
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
func (c *Client) ListRepoMilestones(owner, repo string, opt ListMilestoneOption) ([]*Milestone, *Response, error) {
	opt.setDefaults()
	milestones := make([]*Milestone, 0, opt.PageSize)

	link, _ := url.Parse(fmt.Sprintf("/repos/%s/%s/milestones", owner, repo))
	link.RawQuery = opt.QueryEncode()
	resp, err := c.getParsedResponse("GET", link.String(), nil, nil, &milestones)
	return milestones, resp, err
}

// GetMilestone get one milestone by repo name and milestone id / name
func (c *Client) GetMilestone(owner, repo string, value interface{}) (*Milestone, *Response, error) {
	id, converted, err := milestoneValueToString(value)
	if err != nil {
		return nil, nil, err
	}
	if !converted && c.CheckServerVersionConstraint(">=1.13") != nil {
		// backwards compatibility mode
		m, resp, err := c.resolveMilestoneByName(owner, repo, id)
		return m, resp, err
	}

	milestone := new(Milestone)
	resp, err := c.getParsedResponse("GET", fmt.Sprintf("/repos/%s/%s/milestones/%s", owner, repo, id), nil, nil, milestone)
	return milestone, resp, err
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
func (c *Client) CreateMilestone(owner, repo string, opt CreateMilestoneOption) (*Milestone, *Response, error) {
	if err := opt.Validate(); err != nil {
		return nil, nil, err
	}
	body, err := json.Marshal(&opt)
	if err != nil {
		return nil, nil, err
	}
	milestone := new(Milestone)
	resp, err := c.getParsedResponse("POST", fmt.Sprintf("/repos/%s/%s/milestones", owner, repo), jsonHeader, bytes.NewReader(body), milestone)

	// make creating closed milestones need gitea >= v1.13.0
	// this make it backwards compatible
	if err == nil && opt.State == StateClosed && milestone.State != StateClosed {
		closed := StateClosed
		return c.EditMilestone(owner, repo, milestone.ID, EditMilestoneOption{
			State: &closed,
		})
	}

	return milestone, resp, err
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

// EditMilestone modify milestone with options by id / name
func (c *Client) EditMilestone(owner, repo string, value interface{}, opt EditMilestoneOption) (*Milestone, *Response, error) {
	id, converted, err := milestoneValueToString(value)
	if err != nil {
		return nil, nil, err
	}
	if !converted && c.CheckServerVersionConstraint(">=1.13") != nil {
		// backwards compatibility mode
		m, _, err := c.resolveMilestoneByName(owner, repo, id)
		if err != nil {
			return nil, nil, err
		}
		id = fmt.Sprint(m.ID)
	}

	if err := opt.Validate(); err != nil {
		return nil, nil, err
	}
	body, err := json.Marshal(&opt)
	if err != nil {
		return nil, nil, err
	}
	milestone := new(Milestone)
	resp, err := c.getParsedResponse("PATCH", fmt.Sprintf("/repos/%s/%s/milestones/%s", owner, repo, id), jsonHeader, bytes.NewReader(body), milestone)
	return milestone, resp, err
}

// DeleteMilestone delete one milestone by id / name
func (c *Client) DeleteMilestone(owner, repo string, value interface{}) (*Response, error) {
	id, converted, err := milestoneValueToString(value)
	if err != nil {
		return nil, err
	}
	if !converted && c.CheckServerVersionConstraint(">=1.13") != nil {
		// backwards compatibility mode
		m, _, err := c.resolveMilestoneByName(owner, repo, id)
		if err != nil {
			return nil, err
		}
		id = fmt.Sprint(m.ID)
	}
	_, resp, err := c.getResponse("DELETE", fmt.Sprintf("/repos/%s/%s/milestones/%s", owner, repo, id), nil, nil)
	return resp, err
}

// milestoneValueToString return string of int/int6/string and if it was converted
func milestoneValueToString(value interface{}) (string, bool, error) {
	ii, ok := value.(int64)
	if ok {
		return fmt.Sprint(ii), true, nil
	}
	o, ok := value.(int)
	if ok {
		return fmt.Sprint(o), true, nil
	}
	s, ok := value.(string)
	if ok {
		return s, false, nil
	}
	return "", false, fmt.Errorf("unsuported type: %T", value)
}

// resolveMilestoneByName is a fallback method to find milestone id by name
func (c *Client) resolveMilestoneByName(owner, repo, name string) (*Milestone, *Response, error) {
	for i := 1; ; i++ {
		miles, resp, err := c.ListRepoMilestones(owner, repo, ListMilestoneOption{
			ListOptions: ListOptions{
				Page: i,
			},
			State: "all",
		})
		if err != nil {
			return nil, nil, err
		}
		if len(miles) == 0 {
			return nil, nil, fmt.Errorf("milestone '%s' do not exist", name)
		}
		for _, m := range miles {
			if strings.ToLower(strings.TrimSpace(m.Title)) == strings.ToLower(strings.TrimSpace(name)) {
				return m, resp, nil
			}
		}
	}
}
