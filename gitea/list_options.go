package gitea

import "fmt"

// ListOptions options for using Gitea's API pagination
type ListOptions struct {
	Page    int
	PerPage int
}

func (o ListOptions) getURLQuery() string {
	o.setDefaults()

	return fmt.Sprintf("page=%d&limit=%d", o.Page, o.PerPage)
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
