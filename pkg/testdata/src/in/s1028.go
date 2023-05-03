package in

import (
	"errors"
	"fmt"
)

func bad() {
	errors.New(fmt.Sprintf("%d", 1))
}
