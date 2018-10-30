package testsvc

// Allocator allocates the test service.
type Allocator interface {
	// Allocate test service.
	Allocate() (AllocatedService, error)
}

// AllocatedService describes the common actions
// that can be performed on the allocated service.
type AllocatedService interface {
	// URL returns the service URL.
	URL() (string, error)

	// Close has to free the allocated resources.
	Close() error
}
