# Glint

Glint glues the approachable syntax of [gopatch](https://github.com/uber-go/gopatch) to the power of [staticcheck](https://honnef.co/go/tools).

# Non Goals (Currently)

Formatting. The initial implementation of this project doesn't care
about formatting. You can run the resultant code through `go fmt` or
something afterwards.

Imports. Goimports does this for you.

# Related Projects
* gopatch - https://github.com/uber-go/gopatch
* eg - https://pkg.go.dev/golang.org/x/tools/cmd/eg
