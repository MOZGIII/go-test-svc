package docker

import (
	"github.com/MOZGIII/go-test-svc/testsvc"
	"github.com/ory/dockertest"
)

// Allocator allocates test service using docker.
type Allocator struct {
	// An endpoint to use to send commands to docker.
	// If it's empty, the actual endpoint will be deduce from
	// env if possible.
	DockerEndpoint string

	// Parameters to run docker container with.
	RunOptions *dockertest.RunOptions

	// A way to get the service URL from the allocated service.
	URLFetcher URLFetcher
}

var _ testsvc.Allocator = (*Allocator)(nil)

// Allocate starts docker container.
func (a *Allocator) Allocate() (testsvc.AllocatedService, error) {
	// Prepare docker connection.
	pool, err := dockertest.NewPool("")
	if err != nil {
		return nil, err
	}

	// Pull an image, create a container based on it and run it.
	resource, err := pool.RunWithOptions(a.RunOptions)
	if err != nil {
		return nil, err
	}

	// Return the allocated service.
	return &AllocatedService{
		Pool:       pool,
		Resource:   resource,
		URLFetcher: a.URLFetcher,
	}, nil
}

// AllocatedService holds the internal dockertest resource that allows
// docker container manipulation and querying.
type AllocatedService struct {
	Pool     *dockertest.Pool
	Resource *dockertest.Resource

	URLFetcher URLFetcher
}

var _ testsvc.AllocatedService = (*AllocatedService)(nil)

// URL applies the URLFetcher and returns a URL or an error
// if it was unable to fetch the URL.
func (d *AllocatedService) URL() (string, error) {
	if d.URLFetcher == nil {
		panic("no url fetcher")
	}
	return d.URLFetcher.FetchURL(d)
}

// Close removes docker container.
func (d *AllocatedService) Close() error {
	// Remove container and assocaited volumes.
	return d.Pool.Purge(d.Resource)
}
