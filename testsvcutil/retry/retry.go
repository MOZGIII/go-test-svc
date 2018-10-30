package retry

import (
	"time"

	"github.com/cenkalti/backoff"
)

const (
	// ExponentialBackOffMaxInterval sets the max interval
	// for exponential back off.
	ExponentialBackOffMaxInterval = 5 * time.Second

	// ExponentialBackOffMaxElapsedTime sets the max elapsed
	// time for exponential back off.
	ExponentialBackOffMaxElapsedTime = 1 * time.Minute
)

// WithExponentialBackOff retries the op with exponential back off.
// If you need to fine-tune the params - just go ahead and use backoff
// package that this package depends on directly.
func WithExponentialBackOff(op backoff.Operation) error {
	bo := backoff.NewExponentialBackOff()
	bo.MaxInterval = ExponentialBackOffMaxInterval
	bo.MaxElapsedTime = ExponentialBackOffMaxElapsedTime
	return backoff.Retry(op, bo)
}
