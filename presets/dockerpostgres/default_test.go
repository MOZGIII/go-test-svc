package dockerpostgres_test

import (
	"testing"

	"github.com/MOZGIII/go-test-svc/presets/dockerpostgres"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDefaultAlloc(t *testing.T) {
	allocator := dockerpostgres.NewDefaultAllocator()
	allocatedDB, err := allocator.Allocate()
	require.NoError(t, err)

	url, err := allocatedDB.URL()
	assert.NoError(t, err)
	assert.Regexp(t, `^postgres://postgres:postgres@localhost:\d+/postgres$`, url)

	err = allocatedDB.Close()
	require.NoError(t, err)
}
