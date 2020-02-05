// Copyright 2015 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package gitea

import "fmt"

// ListFollowersOptions options for listing followers
type ListFollowersOptions struct {
	ListOptions
}

// ListMyFollowers list all the followers of current user
func (c *Client) ListMyFollowers(opt ListFollowersOptions) ([]*User, error) {
	opt.setDefaults()
	users := make([]*User, 0, opt.PageSize)
	return users, c.getParsedResponse("GET", fmt.Sprintf("/user/followers?%s", opt.getURLQuery().Encode()), nil, nil, &users)
}

// ListFollowers list all the followers of one user
func (c *Client) ListFollowers(user string, opt ListFollowersOptions) ([]*User, error) {
	opt.setDefaults()
	users := make([]*User, 0, opt.PageSize)
	return users, c.getParsedResponse("GET", fmt.Sprintf("/users/%s/followers?%s", user, opt.getURLQuery().Encode()), nil, nil, &users)
}

// ListFollowingOptions options for listing a user's users being followed
type ListFollowingOptions struct {
	ListOptions
}

// ListMyFollowing list all the users current user followed
func (c *Client) ListMyFollowing(opt ListFollowingOptions) ([]*User, error) {
	opt.setDefaults()
	users := make([]*User, 0, opt.PageSize)
	return users, c.getParsedResponse("GET", fmt.Sprintf("/user/following?%s", opt.getURLQuery().Encode()), nil, nil, &users)
}

// ListFollowing list all the users the user followed
func (c *Client) ListFollowing(user string, opt ListFollowingOptions) ([]*User, error) {
	opt.setDefaults()
	users := make([]*User, 0, opt.PageSize)
	return users, c.getParsedResponse("GET", fmt.Sprintf("/users/%s/following?%s", user, opt.getURLQuery().Encode()), nil, nil, &users)
}

// IsFollowing if current user followed the target
func (c *Client) IsFollowing(target string) bool {
	_, err := c.getResponse("GET", fmt.Sprintf("/user/following/%s", target), nil, nil)
	return err == nil
}

// IsUserFollowing if the user followed the target
func (c *Client) IsUserFollowing(user, target string) bool {
	_, err := c.getResponse("GET", fmt.Sprintf("/users/%s/following/%s", user, target), nil, nil)
	return err == nil
}

// Follow set current user follow the target
func (c *Client) Follow(target string) error {
	_, err := c.getResponse("PUT", fmt.Sprintf("/user/following/%s", target), nil, nil)
	return err
}

// Unfollow set current user unfollow the target
func (c *Client) Unfollow(target string) error {
	_, err := c.getResponse("DELETE", fmt.Sprintf("/user/following/%s", target), nil, nil)
	return err
}
