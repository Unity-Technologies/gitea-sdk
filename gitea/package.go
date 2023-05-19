package gitea

import (
	"fmt"
	"time"
)

// Package represents a package
type Package struct {
	// the package's id
	ID int64 `json:"id"`
	// the package's owner
	Owner User `json:"owner"`
	// the repo this package belongs to (if any)
	Repository *string `json:"repository"`
	// the package's creator
	Creator User `json:"creator"`
	// the type of package:
	Type string `json:"type"`
	// the name of the package
	Name string `json:"name"`
	// the version of the package
	Version string `json:"version"`
	// the date the package was uploaded
	CreatedAt time.Time `json:"created_At"`
}

// PackageFile represents a file from a package
type PackageFile struct {
	// the file's ID
	ID int64 `json:"id"`
	// the size of the file in bytes
	Size int64 `json:"size"`
	// the name of the file
	Name string `json:"name"`
	// the md5 hash
	MD5 string `json:"md5"`
	// the md5 hash
	SHA1 string `json:"sha1"`
	// the md5 hash
	SHA256 string `json:"sha256"`
	// the md5 hash
	SHA512 string `json:"sha512"`
}

// ListOwnerPackagesOptions options for listing packages
type ListOwnerPackagesOptions struct {
	ListOptions
}

// List all the packages owned by a given owner (user, organisation)
func (c *Client) ListOwnerPackages(owner string, opt ListOwnerPackagesOptions) ([]*Package, *Response, error) {
	if err := escapeValidatePathSegments(&owner); err != nil {
		return nil, nil, err
	}
	opt.setDefaults()
	packages := make([]*Package, 0, opt.PageSize)
	resp, err := c.getParsedResponse("GET", fmt.Sprintf("/packages/%s?%s", owner, opt.getURLQuery().Encode()), nil, nil, &packages)
	return packages, resp, err
}

// List the versions of a given package
func (c *Client) GetPackageByVersion(owner, packageType, name, version string) (*Package, *Response, error) {
	if err := escapeValidatePathSegments(&owner, &packageType, &name, &version); err != nil {
		return nil, nil, err
	}
	foundPackage := new(Package)
	resp, err := c.getParsedResponse("GET", fmt.Sprintf("/packages/%s/%s/%s/%s", owner, packageType, name, version), nil, nil, foundPackage)
	return foundPackage, resp, err
}

// List the files of a given package
func (c *Client) GetPackageFilesByVersion(owner, packageType, name, version string) ([]*PackageFile, *Response, error) {
	if err := escapeValidatePathSegments(&owner, &packageType, &name, &version); err != nil {
		return nil, nil, err
	}
	packageFiles := make([]*PackageFile, 0)
	resp, err := c.getParsedResponse("GET", fmt.Sprintf("/packages/%s/%s/%s/%s/files", owner, packageType, name, version), nil, nil, &packageFiles)
	return packageFiles, resp, err
}
