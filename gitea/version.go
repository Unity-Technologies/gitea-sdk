// Copyright 2020 The Gitea Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package gitea

import (
	"fmt"

	version "github.com/hashicorp/go-version"
)

var serverVersion *version.Version

// ServerVersion returns the version of the server
func (c *Client) ServerVersion() (string, error) {
	var v = struct {
		Version string `json:"version"`
	}{}
	return v.Version, c.getParsedResponse("GET", "/version", nil, nil, &v)
}

// CheckServerVersionConstraint validates that the login's server satisfies a
// given version constraint such as ">= 1.11.0+dev"
func (c *Client) CheckServerVersionConstraint(constraint string) error {
	if serverVersion == nil {
		raw, err := c.ServerVersion()
		if err != nil {
			return err
		}
		if serverVersion, err = version.NewVersion(raw); err != nil {
			return err
		}
	}
	check, err := version.NewConstraint(constraint)
	if err != nil {
		return err
	}
	if !check.Check(serverVersion) {
		return fmt.Errorf("gitea server at %s does not satisfy version constraint %s", c.url, constraint)
	}
	return nil
}
