// Copyright 2014 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package gitea

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
)

// BasicAuthEncode generate base64 of basic auth head
func BasicAuthEncode(user, pass string) string {
	return base64.StdEncoding.EncodeToString([]byte(user + ":" + pass))
}

// AccessToken represents a API access token.
// swagger:response AccessToken
type AccessToken struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
	Sha1 string `json:"sha1"`
	Content string `json:"content,omitempty"`
}

// AccessTokenList represents a list of API access token.
// swagger:response AccessTokenList
type AccessTokenList []*AccessToken

// ListAccessTokens lista all the access tokens of user
func (c *Client) ListAccessTokens(user, pass string) ([]*AccessToken, error) {
	tokens := make([]*AccessToken, 0, 10)
	return tokens, c.getParsedResponse("GET", fmt.Sprintf("/users/%s/tokens", user),
		http.Header{"Authorization": []string{"Basic " + BasicAuthEncode(user, pass)}}, nil, &tokens)
}

// AdminListAccessTokens lista all the access tokens of user, authenticate with pre-generated server token.
// this allows server app to list and generate access token for a user already authenticated through other means.
func (c *Client) AdminListAccessTokens(user, server_token string) ([]*AccessToken, error) {
	tokens := make([]*AccessToken, 0, 10)
	return tokens, c.getParsedResponse("GET", fmt.Sprintf("/users/%s/tokens", user),
		http.Header{"X-Gitea-Server-Access-Token": []string{server_token}}, nil, &tokens)
}

// CreateAccessTokenOption options when create access token
// swagger:parameters userCreateToken
type CreateAccessTokenOption struct {
	Name string `json:"name" binding:"Required"`
	MatchOwner []string `json:"match_owner,omitempty"`
	MatchRepo []string `json:"match_repo,omitempty"`
	WildcardMatchBranch []string `json:"wildcard_match_branch,omitempty"`
	WildcardMatchRoute []string `json:"wildcard_match_route,omitempty"`
	MatchMethod []string `json:"match_method,omitempty"`
	ExpiresAt int64 `json:"expires_at,omitempty"`
	// allow integrated server app to authenticate by pre-generated token
	// and to deprecate basic auth by username and password.
	// this will also give server app the option to generate user access token
	// on the fly without storing token.
	GiteaServerAccessToken string `json:"-"`
	UserName 		string `json:"user_name,omitempty"`
}

// CreateAccessToken create one access token with options
func (c *Client) CreateAccessToken(user, pass string, opt CreateAccessTokenOption) (*AccessToken, error) {
	body, err := json.Marshal(&opt)
	if err != nil {
		return nil, err
	}
	t := new(AccessToken)
	return t, c.getParsedResponse("POST", fmt.Sprintf("/users/%s/tokens", user),
		http.Header{
			"content-type":  []string{"application/json"},
			"Authorization": []string{"Basic " + BasicAuthEncode(user, pass)},
			"X-Gitea-Server-Access-Token": []string{opt.GiteaServerAccessToken},
		},
		bytes.NewReader(body), t)
}

// DeleteAccessToken delete token with key id
func (c *Client) DeleteAccessToken(user string, keyID int64) error {
	_, err := c.getResponse("DELETE", fmt.Sprintf("/user/%s/tokens/%d", user, keyID), nil, nil)
	return err
}
