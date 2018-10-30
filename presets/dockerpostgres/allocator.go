package dockerpostgres

import (
	"fmt"
	"net/url"

	"github.com/MOZGIII/go-test-svc/testsvc"
	"github.com/MOZGIII/go-test-svc/testsvc/docker"
	"github.com/ory/dockertest"
)

// Allocator allocates the postgres database using ephemeral docker
// container.
type Allocator struct {
	// DockerEndpoint to pass to docker
	// database allocator.
	DockerEndpoint string

	// The docker repository to use.
	DockerRepository string

	// The docker tag to use.
	DockerTag string

	// The internal port postgres listens on;
	// this port will be exposed from the container.
	PostgresInternalPort string

	// The value of POSTGRES_PASSWORD env var.
	// On the official image this sets the password for
	// `postgres` user.
	PostgresPassword string

	// The tempate of the URL to return as the database URL.
	// Port will be replaced by external port exposed from
	// the container.
	TemplateURL string
}

var _ testsvc.Allocator = (*Allocator)(nil)

// Allocate starts the postgres container.
func (a *Allocator) Allocate() (testsvc.AllocatedService, error) {
	parsedTemplateURL, err := url.Parse(a.TemplateURL)
	if err != nil {
		return nil, err
	}

	urlFetcher := docker.URLWithExposedPort{
		Builder: &docker.ExternalPortOverwriteURLBuilder{
			TemplateURL: parsedTemplateURL,
		},
		InternalPort: a.PostgresInternalPort,
	}

	allocator := docker.Allocator{
		DockerEndpoint: a.DockerEndpoint,
		URLFetcher:     &urlFetcher,
		RunOptions: &dockertest.RunOptions{
			Repository: a.DockerRepository,
			Tag:        a.DockerTag,

			Env: []string{
				fmt.Sprintf("POSTGRES_PASSWORD=%s", a.PostgresPassword),
			},

			ExposedPorts: []string{
				a.PostgresInternalPort,
			},
		},
	}
	return allocator.Allocate()
}
