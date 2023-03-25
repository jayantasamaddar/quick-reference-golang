# Table of Contents

1. [The Go Language](golang)
2. [Basic HTTP Server](basic-http-server)
3. [Basic CRUD API](basic-CRUD-api)
4. [CLI Reminder Tool](cli-reminder)

---

# List of Go Commands

| Command                     | Use                                                               |
| --------------------------- | ----------------------------------------------------------------- |
| `go help [command]`         | Get detailed help on a Go command                                 |
| `go mod init [module-path]` | Initializes and writes a new go.mod file in the current directory |
| `go mod tidy`               | Adds missing modules and removes unused modules                   |
| `go mod download`           | Downloads named modules. (Useful when building Docker image)      |
| `go clean -modcache`        | Removes the entire module download cache                          |
| `go get [packages...]`      | Add or upgrade a dependency/dependencies to the current module    |
| `go get package@v1.2.3`     | Upgrade or downgrade a dependency to a specific version           |
| `go get package@none`       | Remove a dependency from the current module                       |
| `go build`                  | Compiles the package but doesn't install it                       |
| `go build -o [output]`      | Compiles the package and names the outfile but doesn't install it |
| `go install [packages...]`  | Compiles and installs the packages named by the import paths      |
| `go work init [modules...]` | Initialize a multi-module workspace                               |
| `go work use [module]`      | Add a module to the multi-module workspace                        |

**Other commands**:

- `go mod edit -replace github.com/user/dependency=../localModule`: Replace a dependency from being fetched with a local module
