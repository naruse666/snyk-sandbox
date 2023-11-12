package main

import (
	"fmt"
	"os"
)

func main() {
	gopath := os.Getenv("GOPATH")

	fmt.Printf("GOPATH is %s", gopath)
}
