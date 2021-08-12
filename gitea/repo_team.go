package gitea

import (
	"fmt"
	"net/http"
)

// GetTeamsOfRepo return teams from a repository
func (c *Client) GetTeamsOfRepo(user, repo string) ([]*Team, *Response, error) {
	if err := escapeValidatePathSegments(&user, &repo); err != nil {
		return nil, nil, err
	}
	teams := make([]*Team, 0, 5)
	resp, err := c.getParsedResponse("GET", fmt.Sprintf("/repos/%s/%s/teams", user, repo), nil, nil, &teams)
	return teams, resp, err
}

// AddTeamToRepo add a team to a repository
func (c *Client) AddTeamToRepo(user, repo, team string) (*Response, error) {
	if err := escapeValidatePathSegments(&user, &repo, &team); err != nil {
		return nil, err
	}
	_, resp, err := c.getResponse("PUT", fmt.Sprintf("/repos/%s/%s/teams/%s", user, repo, team), nil, nil)
	return resp, err
}

// RemoveTeamFromRepo delete a team from a repository
func (c *Client) RemoveTeamFromRepo(user, repo, team string) (*Response, error) {
	if err := escapeValidatePathSegments(&user, &repo, &team); err != nil {
		return nil, err
	}
	_, resp, err := c.getResponse("DELETE", fmt.Sprintf("/repos/%s/%s/teams/%s", user, repo, team), nil, nil)
	return resp, err
}

// IsTeamAssignedToRepo return team if assigned to repo else nil
func (c *Client) IsTeamAssignedToRepo(user, repo, team string) (*Team, *Response, error) {
	if err := escapeValidatePathSegments(&user, &repo, &team); err != nil {
		return nil, nil, err
	}
	t := new(Team)
	resp, err := c.getParsedResponse("GET", fmt.Sprintf("/repos/%s/%s/teams/%s", user, repo, team), nil, nil, &t)
	if resp != nil && resp.StatusCode == http.StatusNotFound {
		// if not found it's not an error, it indicates it's not assigned
		return nil, resp, nil
	}
	return t, resp, err
}
