package gitea

import "fmt"

type searchUsersResponse struct {
	Users []*User `json:"data"`
}

// SearchUsers finds users by query
func (c *Client) SearchUsers(query string, limit int) ([]*User, error) {
	users := make([]*User, 0, 10)
	err := c.getParsedResponse("GET", fmt.Sprintf("/users/search?q=%s&limit=%d", query, limit), nil, nil, &users)
	return users, err
}
