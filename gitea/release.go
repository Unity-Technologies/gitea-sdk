// Copyright 2016 The Gitea Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package gitea

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"
)

// Release represents a repository release
type Release struct {
	ID           int64         `json:"id"`
	TagName      string        `json:"tag_name"`
	Target       string        `json:"target_commitish"`
	Title        string        `json:"name"`
	Note         string        `json:"body"`
	URL          string        `json:"url"`
	HTMLURL      string        `json:"html_url"`
	TarURL       string        `json:"tarball_url"`
	ZipURL       string        `json:"zipball_url"`
	IsDraft      bool          `json:"draft"`
	IsPrerelease bool          `json:"prerelease"`
	CreatedAt    time.Time     `json:"created_at"`
	PublishedAt  time.Time     `json:"published_at"`
	Publisher    *User         `json:"author"`
	Attachments  []*Attachment `json:"assets"`
}

// ListReleasesOptions options for listing repository's releases
type ListReleasesOptions struct {
	ListOptions
}

// ListReleases list releases of a repository
// response support Next()
func (c *Client) ListReleases(user, repo string, opt ListReleasesOptions) ([]*Release, *Response, error) {
	if err := opt.saveSetDefaults(c); err != nil {
		return nil, nil, err
	}
	releases := make([]*Release, 0, opt.PageSize)
	resp, err := c.getParsedResponse("GET", fmt.Sprintf("/repos/%s/%s/releases?%s", user, repo, opt.getURLQuery().Encode()), nil, nil, &releases)
	if err = c.preparePaginatedResponse(resp, &opt.ListOptions, len(releases)); err != nil {
		return releases, resp, err
	}
	return releases, resp, err
}

// GetRelease get a release of a repository by id
func (c *Client) GetRelease(user, repo string, id int64) (*Release, *Response, error) {
	r := new(Release)
	resp, err := c.getParsedResponse("GET",
		fmt.Sprintf("/repos/%s/%s/releases/%d", user, repo, id),
		jsonHeader, nil, &r)
	return r, resp, err
}

// GetReleaseByTag get a release of a repository by tag
func (c *Client) GetReleaseByTag(user, repo string, tag string) (*Release, *Response, error) {
	if c.checkServerVersionGreaterThanOrEqual(version1_13_0) != nil {
		return c.fallbackGetReleaseByTag(user, repo, tag)
	}
	r := new(Release)
	resp, err := c.getParsedResponse("GET",
		fmt.Sprintf("/repos/%s/%s/releases/tags/%s", user, repo, tag),
		nil, nil, &r)
	return r, resp, err
}

// CreateReleaseOption options when creating a release
type CreateReleaseOption struct {
	TagName      string `json:"tag_name"`
	Target       string `json:"target_commitish"`
	Title        string `json:"name"`
	Note         string `json:"body"`
	IsDraft      bool   `json:"draft"`
	IsPrerelease bool   `json:"prerelease"`
}

// Validate the CreateReleaseOption struct
func (opt CreateReleaseOption) Validate() error {
	if len(strings.TrimSpace(opt.Title)) == 0 {
		return fmt.Errorf("title is empty")
	}
	return nil
}

// CreateRelease create a release
func (c *Client) CreateRelease(user, repo string, opt CreateReleaseOption) (*Release, *Response, error) {
	if err := opt.Validate(); err != nil {
		return nil, nil, err
	}
	body, err := json.Marshal(opt)
	if err != nil {
		return nil, nil, err
	}
	r := new(Release)
	resp, err := c.getParsedResponse("POST",
		fmt.Sprintf("/repos/%s/%s/releases", user, repo),
		jsonHeader, bytes.NewReader(body), r)
	return r, resp, err
}

// EditReleaseOption options when editing a release
type EditReleaseOption struct {
	TagName      string `json:"tag_name"`
	Target       string `json:"target_commitish"`
	Title        string `json:"name"`
	Note         string `json:"body"`
	IsDraft      *bool  `json:"draft"`
	IsPrerelease *bool  `json:"prerelease"`
}

// EditRelease edit a release
func (c *Client) EditRelease(user, repo string, id int64, form EditReleaseOption) (*Release, *Response, error) {
	body, err := json.Marshal(form)
	if err != nil {
		return nil, nil, err
	}
	r := new(Release)
	resp, err := c.getParsedResponse("PATCH",
		fmt.Sprintf("/repos/%s/%s/releases/%d", user, repo, id),
		jsonHeader, bytes.NewReader(body), r)
	return r, resp, err
}

// DeleteRelease delete a release from a repository, keeping its tag
func (c *Client) DeleteRelease(user, repo string, id int64) (*Response, error) {
	_, resp, err := c.getResponse("DELETE",
		fmt.Sprintf("/repos/%s/%s/releases/%d", user, repo, id),
		nil, nil)
	return resp, err
}

// DeleteReleaseTag deletes a tag from a repository, if no release refers to it.
func (c *Client) DeleteReleaseTag(user, repo string, tag string) (*Response, error) {
	if err := c.checkServerVersionGreaterThanOrEqual(version1_14_0); err != nil {
		return nil, err
	}
	_, resp, err := c.getResponse("DELETE",
		fmt.Sprintf("/repos/%s/%s/releases/tags/%s", user, repo, tag),
		nil, nil)
	return resp, err
}

// fallbackGetReleaseByTag is fallback for old gitea installations ( < 1.13.0 )
func (c *Client) fallbackGetReleaseByTag(user, repo string, tag string) (*Release, *Response, error) {
	for i := 1; ; i++ {
		rl, resp, err := c.ListReleases(user, repo, ListReleasesOptions{ListOptions{Page: i}})
		if err != nil {
			return nil, resp, err
		}
		if len(rl) == 0 {
			return nil,
				&Response{Response: &http.Response{StatusCode: 404}},
				fmt.Errorf("release with tag '%s' not found", tag)
		}
		for _, r := range rl {
			if r.TagName == tag {
				return r, resp, nil
			}
		}
	}
}
