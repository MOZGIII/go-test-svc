package dockerpostgres

// NewDefaultAllocator provides the postgres allocator
// with a sane-for-usual-case defaults.
// You can alter the params to fit your needs before
// calling Allocate.
func NewDefaultAllocator() *Allocator {
	return &Allocator{
		// Docker endpoint is empty - meaning it will be detected
		// from env.
		DockerEndpoint: "",

		// Use official postgres image by default.
		DockerRepository: "postgres",

		// No better default than "latest".
		DockerTag: "latest",

		// Postgres listens on TCP port 5432 by default.
		PostgresInternalPort: "5432/tcp",

		// Default password for postgres DB.
		PostgresPassword: "postgres",

		// Default connection params (excluding the port).
		TemplateURL: "postgres://postgres:postgres@localhost:0/postgres",
	}
}
