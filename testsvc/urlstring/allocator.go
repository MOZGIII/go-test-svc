package urlstring

import (
	"github.com/MOZGIII/go-test-svc/testsvc"
)

// Allocator allocates nothing and just returns the specified URL.
type Allocator struct {
	AllocatedServiceURL string
}

var _ testsvc.Allocator = (*Allocator)(nil)

// Allocate allocates nothing and just returns self.
func (a *Allocator) Allocate() (testsvc.AllocatedService, error) {
	return a, nil
}

var _ testsvc.AllocatedService = (*Allocator)(nil)

// URL returns the specified URL.
func (a *Allocator) URL() (string, error) {
	return a.AllocatedServiceURL, nil
}

// Close does nothing.
func (a *Allocator) Close() error {
	// noop
	return nil
}
