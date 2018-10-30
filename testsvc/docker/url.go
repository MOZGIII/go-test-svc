package docker

import (
	"fmt"
	"net/url"
)

// URLFetcher takes *AllocatedService and provides the
// URL that can be used to connect to that allocated service.
type URLFetcher interface {
	FetchURL(*AllocatedService) (string, error)
}

// ExposedPortURLBuilder takes the exposed external
// port and gives the full URL.
type ExposedPortURLBuilder interface {
	BuildURL(externalPort string) (string, error)
}

// URLWithExposedPort utilizes the ExposedPortURLBuilder to build
// the URL with exposed port.
type URLWithExposedPort struct {
	Builder      ExposedPortURLBuilder
	InternalPort string
}

var _ URLFetcher = (*URLWithExposedPort)(nil)

// FetchURL returns URL to use to connect to the passed allocated service.
func (u *URLWithExposedPort) FetchURL(allocatedService *AllocatedService) (string, error) {
	if allocatedService == nil {
		panic("no allocated service passed")
	}
	if allocatedService.Resource == nil {
		panic("no dockertest resource in the allocated service")
	}
	externalPort := allocatedService.Resource.GetPort(u.InternalPort)
	return u.Builder.BuildURL(externalPort)
}

// ExternalPortOverwriteURLBuilder overwrites the port of the URL,
// leaving the rest intact
type ExternalPortOverwriteURLBuilder struct {
	TemplateURL *url.URL
}

var _ ExposedPortURLBuilder = (*ExternalPortOverwriteURLBuilder)(nil)

// BuildURL build postgres URL with the specified port.
func (b *ExternalPortOverwriteURLBuilder) BuildURL(externalPort string) (string, error) {
	url := *b.TemplateURL
	url.Host = fmt.Sprintf("%s:%s", url.Hostname(), externalPort)
	return url.String(), nil
}
