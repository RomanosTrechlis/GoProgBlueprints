package main

import (
	"fmt"
	"flag"
	"os"
)

var fatalErr error

func fatal(e error) {
	fmt.Println(e)
	flag.PrintDefaults()
	fatalErr = e
}

func main() {
	// this must be the last function to run before program ends,
	// that is way it's the first function we defer (defer is LIFO).
	defer func() {
		if fatalErr != nil {
			os.Exit(1)
		}
	}()
}
