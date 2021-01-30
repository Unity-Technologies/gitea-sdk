package gitea

import (
	"encoding/json"
	"net/url"
	"reflect"
)

type paginateOption interface {
	getURLQuery() url.Values
	getPage() int
	getPageSize() int
	saveSetDefaults(c *Client) error
}

func (o ListOptions) getPage() int {
	return o.Page
}
func (o ListOptions) getPageSize() int {
	return o.PageSize
}

// getParsedPaginatedResponse is used for list apis
// response support Next() and pagination defaults are ensured
func (c *Client) getParsedPaginatedResponse(method string, path *url.URL, opts paginateOption, obj interface{}) (*Response, error) {
	if err := opts.saveSetDefaults(c); err != nil {
		return nil, err
	}
	path.RawQuery = opts.getURLQuery().Encode()
	data, resp, err := c.getResponse(method, path.String(), jsonHeader, nil)
	if err != nil {
		return resp, err
	}
	if err = json.Unmarshal(data, obj); err != nil {
		return resp, err
	}

	return resp, c.preparePaginatedResponse(resp, &ListOptions{Page: opts.getPage(), PageSize: opts.getPageSize()}, getListLen(obj))
}

func getListLen(list interface{}) int {
	listValue := reflect.ValueOf(list)
	listKind := reflect.TypeOf(list).Kind()
	switch listKind {
	case reflect.Map, reflect.Slice, reflect.Array:
		return listValue.Len()
	default:
		return 0
	}
}
