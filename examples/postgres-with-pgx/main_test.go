package mydb

import (
	"log"
	"os"
	"testing"
	"time"

	"github.com/MOZGIII/go-test-svc/presets/dockerpostgres"
	"github.com/MOZGIII/go-test-svc/testsvc"
	"github.com/MOZGIII/go-test-svc/testsvc/urlstring"
	"github.com/MOZGIII/go-test-svc/testsvcutil/retry"
	"github.com/jackc/pgx"
)

var (
	testConnPool *pgx.ConnPool
)

func TestMain(m *testing.M) {
	testDatabaseURL := os.Getenv("TEST_DATABASE_URL")
	noDocker := os.Getenv("TEST_DONT_USE_DOCKER") == "true"
	useDocker := !noDocker && testDatabaseURL == ""

	var testDBAllocator testsvc.Allocator

	if useDocker {
		testDBAllocator = dockerpostgres.NewDefaultAllocator()
	} else {
		testDBAllocator = &urlstring.Allocator{
			AllocatedServiceURL: testDatabaseURL,
		}
	}

	allocatedTestDB, err := testDBAllocator.Allocate()
	if err != nil {
		log.Fatalf("Could not allocate the test database: %s", err)
	}

	// Exponential backoff-retry, because postgres in the container might
	// not be ready to accept connections yet.
	if err := retry.WithExponentialBackOff(func() error {
		// Get the URL.
		url, err := allocatedTestDB.URL()
		if err != nil {
			return err
		}

		// Prepare connection options.
		connConfig, err := pgx.ParseURI(url)
		if err != nil {
			return err
		}

		// Open database connection.
		testConnPool, err = pgx.NewConnPool(pgx.ConnPoolConfig{
			ConnConfig:     connConfig,
			AcquireTimeout: 60 * time.Second,
		})
		if err != nil {
			return err
		}

		// Ping the database.
		_, err = testConnPool.Exec(";")
		return err
	}); err != nil {
		log.Fatalf("Could not connect to test database: %s", err)
	}

	defer func() {
		if err := allocatedTestDB.Close(); err != nil {
			log.Fatalf("Could not close the test database: %s", err)
		}
	}()

	code := m.Run()

	os.Exit(code)
}
