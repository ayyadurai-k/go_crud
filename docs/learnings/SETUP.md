# Go Modules & Tooling Notes

## Go Modules

### `go mod init`

Creates a new Go module.

```bash
go mod init crud
```

Generates:

```go
module crud

go 1.26.3
```

---

## `go.mod`

Defines:

* Module name
* Go version
* Project dependencies

Example:

```go
module crud

go 1.26.3

require github.com/joho/godotenv v1.5.1
```

Equivalent to:

* Python → `requirements.txt`
* Node.js → `package.json`

---

## `go.sum`

Stores cryptographic hashes of dependencies.

Purpose:

* Verify dependency integrity
* Prevent tampered packages
* Ensure reproducible builds

Think:

```text
go.mod -> What dependencies are needed
go.sum -> Verify downloaded dependencies are authentic
```

---

## `go get`

Adds or updates project dependencies.

```bash
go get github.com/joho/godotenv
```

Updates:

* `go.mod`
* `go.sum`

Used for libraries imported by application code.

---

## `go install`

Installs executable tools.

```bash
go install github.com/githubnemo/CompileDaemon@latest
```

Installs binary into:

```text
$GOPATH/bin
```

Example tools:

* CompileDaemon
* gopls
* air

Used for development tooling, not application dependencies.

---

## `go mod tidy`

Synchronizes dependencies.

```bash
go mod tidy
```

Actions:

* Removes unused dependencies
* Adds missing dependencies
* Cleans `go.mod`
* Updates `go.sum`

Best practice:

```bash
go mod tidy
```

after adding or removing imports.

---

## GOROOT vs GOPATH

### GOROOT

Location of Go installation.

```text
/usr/local/go
```

Contains:

* Go compiler
* Standard library
* Go tools

---

### GOPATH

Workspace for user-installed tools and module cache.

```text
/Users/ayyadurai_k/go
```

Contains:

```text
bin/
pkg/
```

---

## PATH Issue Learned

`go install` places binaries in:

```text
~/go/bin
```

Shell can only execute commands if that directory is in `PATH`.

Add:

```bash
export PATH="$PATH:$HOME/go/bin"
```

to:

```bash
~/.zshrc
```

Reload:

```bash
source ~/.zshrc
```

Verify:

```bash
which CompileDaemon
```

---

## CompileDaemon

Purpose:

* Watch source files
* Rebuild automatically
* Restart application

Example:

```bash
CompileDaemon \
  -build="go build -o crud ." \
  -command="./crud"
```

Workflow:

```text
File Change
    ↓
Rebuild
    ↓
Restart Application
```

Equivalent to:

```text
Node.js -> nodemon
Python  -> watchdog / reload
Go      -> CompileDaemon
```

---

## Key Takeaways

* `go.mod` = dependency definitions
* `go.sum` = dependency verification
* `go get` = project dependency
* `go install` = CLI tool installation
* `go mod tidy` = cleanup and sync
* `GOROOT` = Go installation
* `GOPATH` = user workspace
* `PATH` determines whether installed tools can be executed from terminal
