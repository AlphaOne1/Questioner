// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"errors"
	"net/url"
	golangswaggerpaths "path"
	"strings"

	"github.com/go-openapi/swag"
)

// GetAPITaskIDSolveURL generates an URL for the get API task ID solve operation
type GetAPITaskIDSolveURL struct {
	ID uint64

	SolutionID uint64

	_basePath string
	// avoid unkeyed usage
	_ struct{}
}

// WithBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *GetAPITaskIDSolveURL) WithBasePath(bp string) *GetAPITaskIDSolveURL {
	o.SetBasePath(bp)
	return o
}

// SetBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *GetAPITaskIDSolveURL) SetBasePath(bp string) {
	o._basePath = bp
}

// Build a url path and query string
func (o *GetAPITaskIDSolveURL) Build() (*url.URL, error) {
	var _result url.URL

	var _path = "/api/task/{id}/solve"

	id := swag.FormatUint64(o.ID)
	if id != "" {
		_path = strings.Replace(_path, "{id}", id, -1)
	} else {
		return nil, errors.New("id is required on GetAPITaskIDSolveURL")
	}

	_basePath := o._basePath
	_result.Path = golangswaggerpaths.Join(_basePath, _path)

	qs := make(url.Values)

	solutionIDQ := swag.FormatUint64(o.SolutionID)
	if solutionIDQ != "" {
		qs.Set("solutionID", solutionIDQ)
	}

	_result.RawQuery = qs.Encode()

	return &_result, nil
}

// Must is a helper function to panic when the url builder returns an error
func (o *GetAPITaskIDSolveURL) Must(u *url.URL, err error) *url.URL {
	if err != nil {
		panic(err)
	}
	if u == nil {
		panic("url can't be nil")
	}
	return u
}

// String returns the string representation of the path with query string
func (o *GetAPITaskIDSolveURL) String() string {
	return o.Must(o.Build()).String()
}

// BuildFull builds a full url with scheme, host, path and query string
func (o *GetAPITaskIDSolveURL) BuildFull(scheme, host string) (*url.URL, error) {
	if scheme == "" {
		return nil, errors.New("scheme is required for a full url on GetAPITaskIDSolveURL")
	}
	if host == "" {
		return nil, errors.New("host is required for a full url on GetAPITaskIDSolveURL")
	}

	base, err := o.Build()
	if err != nil {
		return nil, err
	}

	base.Scheme = scheme
	base.Host = host
	return base, nil
}

// StringFull returns the string representation of a complete url
func (o *GetAPITaskIDSolveURL) StringFull(scheme, host string) string {
	return o.Must(o.BuildFull(scheme, host)).String()
}