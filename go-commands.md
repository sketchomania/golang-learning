# Go Notes

## Go Environment Variables

```bash
go env
GOOS="linux" go build
GOOS="darwin" go build
GOOS="windows" go build

# GOROOT tells us where the go runtime is installed
GOROOT

# GOPATH tells us where the local go projects live
# this is the go workspace that we create after the goruntime installation
GOPATH

## edit | add env variable in linux
# "$" -> variable
# User settings
export GOROOT=/usr/local/go
export GOPATH=$HOME/workspace
export PATH=$PATH:GOROOT/bin:$GOPATH/bin
```

## Go Sub folders

```bash
/go
1. /src
 - it contains all the library packages
2. /bin
 - /go
   - it contains executables or programs that come with go runtime
   - with go executables you can run commands like "go version", "go build" and so on
 - /gofmt
   - for making code look pretty while development


/goworkspace
# we need to create this for our local go projects
# we can name it whatever we want
# just make sure to create 3 sub folders in it
1. /scr
  - for source code of future go projects
2. /bin
  - for executable of future go projects
  - files inside this folder are auto-generated
3. /pkg
  - for package and dependencies of future go projects
  - files inside this folder are auto-generated


source
- https://www.youtube.com/watch?v=k8LClK96NZ4
```

## Go Mod Edit

go mod edit -go <version(1.16)>
go mod edit -module <module name>

go get -u github.com/gorilla/mux
go get go.mongodb.org/mongo-driver/mongo

// update which works when mux is not working properly
go install golang.org/x/tools/gopls@latest

## go mod

- Go modules is a native package management tool

```bash
# (an expensive operation so better keep eye on that while running this command)
go mod tidy
go mod verify
```

go list
go list all (get all package)
go list -m all (get all package used only)

go list -m --versions github.com/gorilla/mux
go list -m -versions github.com/gorilla/mux

go mod why github.com/gorilla/mux
go mod graph

go mod vendor (orver write modules)
(almost like node_modules)
(now it will not bring everything form the web)
(bring everything at one time pack everything inside)
(to run from vendor)
go run -mod=vendor main.go (first look everything into vendor folder)

go run main.go (bring everything from the web or chache to run not from the vendor)

go run -race example.go (check for the race condition in the program)

Ctrl + P then type "go tools" in vscode to install update tools for go (like gopls)

// run everything

- go run .

## check for race condition

go run --race .
// exit status 66

## PACKAGE 
### Types of packages in go
- executable packages
  - main function
  - main package
  - entry point of every program execution
- non-executable packages
  - non main | utility package (all not main packages are utility package)
  - not self executable 
  - utilized by main package

### Package Naming
- short, no camel/snake
- descriptive
- same name as directory name
  - don't declare multiple packages in multiple files under the same directory
- package should be **small and many**  

### Package scope
Note - circular references is not allowed among packagesF
1. block scope
  - inside the function block { }
2. unexported
  - outside of function body
  - camelcase ->  "variableName" 
3. exported
  - outside of function body
  - capitalcase -> "VariableName"

🌐 source
- https://www.youtube.com/watch?v=1MdX9Z9fWWw