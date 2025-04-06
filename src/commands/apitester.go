package main

import (
	apitester "github.com/andymaclean/apitester/src/apitester"
)

func main() {
	opts := apitester.DefaultOptions()
	opts.ReadCommandArguments()

	apitester.Run(&opts)
}
