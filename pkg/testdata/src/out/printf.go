package a

import "fmt"

func main() {
	_ = fmt.Sprint("bar") // want `bar`
	_ = fmt.Sprintf("bar %s", 1) // want `foo`
}
