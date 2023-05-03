# Name Pending

Kinda like an AST/analysis swiss army knife.

Should make writing linters and correctors easy.
Should make complex one off refactors, easy.

Refactoring tool/toolkit that leverages the analysis package for all the underpinnings.

Bits and pieces:

# AST Pattern Matching Kit
The underlying implementation of an efficient AST
pattern matcher. 

Aims to be expressive and extensible.

go DSLs should output this interface.

Patching languages should "compile" into this.

Ideally, this should support both `eg` and `gopatch` style
"patterns".


# AST Rewriter

Similar to the above, this is just a nicer library to
rewrite golang ASTs.

Ideally, this outputs diff files.

# Ideas

Should start with some concrete cases.

Static check violations are probably a great place to start.

We can have the `eg` version that fixes it and the `gopatch` version to fix it.

Having the output as a diff would be pretty sick.

# Non Goals (Currently)

Formatting. The initial implementation of this project doesn't care
about formatting. You can run the resultant code through `go fmt` or
something afterwards.

Imports. Goimports does this for you.

# Links
* gopatch - https://github.com/uber-go/gopatch
* eg - https://pkg.go.dev/golang.org/x/tools/cmd/eg
