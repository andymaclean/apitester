package apitester

import "fmt"

func Run(opts *Options) {
	if opts.Verbose {
		fmt.Printf("APITester.\n%s", opts.Summary())
	}

	RESTCall(opts, "GET", "/hello", nil)
	RESTCall(opts, "GET", "/foobar", nil)
	RESTCall(opts, "GET", "/exit", nil)
}
