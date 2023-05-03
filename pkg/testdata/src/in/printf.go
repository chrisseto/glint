package in

import "fmt"

func main() {
	_ = fmt.Sprintf("bar") // want `bar`
	_ = fmt.Sprintf("bar %s", 1) // want `foo`
}
