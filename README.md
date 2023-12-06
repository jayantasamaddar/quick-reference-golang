# Table of Contents

1. [The Go Language](golang)
2. [Design Patterns in Go](design-patterns)
3. [Data Structures in Go](data-structures)
4. **Examples**
   - [CLI Reminder Tool](examples/cli-reminder)
   - [Basic HTTP Server](examples/basic-http-server)
   - [Mux CRUD API](examples/mux-CRUD-api/)
   - [Echo CRUD API](examples/echo-CRUD-api)
   - [MongoDB CRUD API](examples/mongo-CRUD-api)
   - [PostgreSQL CRUD API](examples/postgres-CRUD-api)
5. [gRPC](grpc)

---

# List of Go Commands

| Command                                                          | Use                                                               |
| ---------------------------------------------------------------- | ----------------------------------------------------------------- |
| `go help [command]`                                              | Get detailed help on a Go command                                 |
| `go mod init [module-path]`                                      | Initializes and writes a new go.mod file in the current directory |
| `go mod tidy`                                                    | Adds missing modules and removes unused modules                   |
| `go mod download`                                                | Downloads named modules. (Useful when building Docker image)      |
| `go mod edit -replace github.com/user/dependency=../localModule` | Replace a dependency from being fetched with a local module       |
| `go clean -modcache`                                             | Removes the entire module download cache                          |
| `go get [packages...]`                                           | Add or upgrade a dependency/dependencies to the current module    |
| `go get package@v1.2.3`                                          | Upgrade or downgrade a dependency to a specific version           |
| `go get package@none`                                            | Remove a dependency from the current module                       |
| `go run [filePath.go]`                                           | Run without building                                              |
| `go build`                                                       | Compiles the package but doesn't install it                       |
| `go build -o [output]`                                           | Compiles the package and names the outfile but doesn't install it |
| `go install [packages...]`                                       | Compiles and installs the packages named by the import paths      |
| `go work init [modules...]`                                      | Initialize a multi-module workspace                               |
| `go work use [module]`                                           | Add a module to the multi-module workspace                        |
| `go env`                                                         | View Go environment variables                                     |
| `go doc [package]`                                               | View documentation in the Terminal for the Standard Library       |
