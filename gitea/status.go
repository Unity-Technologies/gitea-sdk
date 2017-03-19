package gitea

import (
	"bytes"
	"encoding/json"
	"fmt"
	"time"
)

type StatusState string

const (
	StatusPending StatusState = "pending"
	StatusSuccess StatusState = "success"
	StatusError   StatusState = "error"
	StatusFailure StatusState = "failure"
	StatusWarning StatusState = "warning"
)

type Status struct {
	ID          int64       `json:"id"`
	State       StatusState `json:"status"`
	TargetURL   string      `json:"target_url"`
	Description string      `json:"description"`
	URL         string      `json:"url"`
	Context     string      `json:"context"`
	Creator     *User       `json:"creator"`
	Created     time.Time   `json:"created_at"`
	Updated     time.Time   `json:"updated_at"`
}

type CombinedStatus struct {
	State      StatusState `json:"state"`
	SHA        string      `json:"sha"`
	TotalCount int         `json:"total_count"`
	Statuses   []*Status   `json:"statuses"`
	Repository *Repository `json:"repository"`
	CommitURL  string      `json:"commit_url"`
	URL        string      `json:"url"`
}

type CreateStatusOption struct {
	State       StatusState `json:"state"`
	TargetURL   string      `json:"target_url"`
	Description string      `json:"description"`
	Context     string      `json:"context"`
}

type ListStatusesOption struct {
	Page int
}

// POST /repos/:owner/:repo/statuses/:sha
func (c *Client) CreateStatus(owner, repo, sha string, opts CreateStatusOption) (*Status, error) {
	body, err := json.Marshal(&opts)
	if err != nil {
		return nil, err
	}
	status := &Status{}
	return status, c.getParsedResponse("POST", fmt.Sprintf("/repos/%s/%s/statuses/%s", owner, repo, sha),
		jsonHeader, bytes.NewReader(body), status)
}

// GET /repos/:owner/:repo/commits/:ref/statuses
func (c *Client) ListStatuses(owner, repo, sha string, opts ListStatusesOption) ([]*Status, error) {
	statuses := make([]*Status, 0, 10)
	return statuses, c.getParsedResponse("GET", fmt.Sprintf("/repos/%s/%s/commits/%s/statuses?page=%d", owner, repo, sha, opts.Page), nil, nil, &statuses)
}

// GET /repos/:owner/:repo/commits/:ref/status
func (c *Client) GetCombinedStatus(owner, repo, sha string) (*CombinedStatus, error) {
	status := &CombinedStatus{}
	return status, c.getParsedResponse("GET", fmt.Sprintf("/repos/%s/%s/commits/%s/status", owner, repo, sha), nil, nil, status)
}
