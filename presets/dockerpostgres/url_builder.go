package dockerpostgres

import (
	"fmt"
	"net/url"

	"github.com/MOZGIII/go-test-svc/testsvc/docker"
)

// ExposedPortURLBuilder build postgres URL.
// Suited for use with docker.ExposedPortURLBuilder.
type ExposedPortURLBuilder struct {
	URL *url.URL
}

var _ docker.ExposedPortURLBuilder = (*ExposedPortURLBuilder)(nil)

// BuildURL build postgres URL with the specified port.
func (b *ExposedPortURLBuilder) BuildURL(externalPort string) (string, error) {
	url := *b.URL
	url.Host = fmt.Sprintf("%s:%s", url.Hostname(), externalPort)
	return url.String(), nil
}
