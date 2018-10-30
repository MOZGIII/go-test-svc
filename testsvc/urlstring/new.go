package urlstring

// New returns a new allocator with the specified URL.
func New(url string) *Allocator {
	return &Allocator{
		AllocatedServiceURL: url,
	}
}
