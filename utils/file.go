package utils

import (
	"fmt"
	"os"
)

func PWD() string {
	dir, err := os.Getwd()
	if err != nil {
		panic(fmt.Errorf("Idiot, couldn't get Current Working Directory :-%v", err))
	}
	return dir
}
