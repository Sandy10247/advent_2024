package utils

import (
	"fmt"
	"time"
)

func TimeFn(name string) func() {
	start := time.Now()
	return func() {
		fmt.Printf("%v took %v seconds\n", name, time.Since(start).Seconds())
	}
}
