go env
GOOS="linux" go build
GOOS="darwin" go build
GOOS="windows" go build

go mod edit -go <version(1.16)>
go mod edit -module <module name>

go get -u github.com/gorilla/mux
go get go.mongodb.org/mongo-driver/mongo

// update which works when mux is not working properly
go install golang.org/x/tools/gopls@latest

## go mod

go mod tidy (an expensive operation so better keep eye on that while running this command)
go mod verify

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
