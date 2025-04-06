package apitester

import "fmt"

func Run(opts *Options) {
	if opts.Verbose {
		fmt.Printf("APITester.\n%s", opts.Summary())
	}
}
