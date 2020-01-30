package gitea

import (
	"fmt"
	"net/url"
)

// ListOptions options for using Gitea's API pagination
type ListOptions struct {
	Page    int
	PerPage int
}

func (o ListOptions) getURLQueryEncoded() string {
	o.setDefaults()

	query := make(url.Values)
	query.Add("page", fmt.Sprintf("%d", o.Page))
	query.Add("limit", fmt.Sprintf("%d", o.PerPage))

	return query.Encode()
}

func (o ListOptions) getURLQuery() url.Values {
	o.setDefaults()

	query := make(url.Values)
	query.Add("page", fmt.Sprintf("%d", o.Page))
	query.Add("limit", fmt.Sprintf("%d", o.PerPage))

	return query
}

func (o ListOptions) setDefaults() {
	if o.Page < 1 {
		o.Page = 1
	}

	if o.PerPage < 0 || o.PerPage > 50 {
		o.PerPage = 10
	}
}

func (o ListOptions) getPerPage() int {
	if o.PerPage < 0 || o.PerPage > 50 {
		o.PerPage = 10
	}

	return o.PerPage
}
