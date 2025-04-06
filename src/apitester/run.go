package apitester

import "fmt"

func Run(opts *Options) {
	if opts.Verbose {
		fmt.Printf("APITester.\n%s", opts.Summary())
	}

	session := NewSession(opts)

	serveraddr := opts.BaseURL[7:]

	session.RESTCall(&APITest{
		method: "GET",
		path:   "/hello",
		checker: APIStringResCheck{
			statuscode: 200,
			body:       fmt.Sprintf("Hello from %v\n", serveraddr),
		},
	})
	session.RESTCall(&APITest{
		method: "GET",
		path:   "/foobar",
		checker: APIStringResCheck{
			statuscode: 404,
		},
	})
	session.RESTCall(&APITest{
		method: "GET",
		path:   "/exit",
		checker: APIStringResCheck{
			statuscode: 200,
			body:       fmt.Sprintf("Goodbye from %v\n", serveraddr),
		},
	})
}
