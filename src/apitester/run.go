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
			status: &IntEq{200},
			body:   &StringEq{fmt.Sprintf("Hello from %v\n", serveraddr)},
		},
	})
	session.RESTCall(&APITest{
		method: "GET",
		path:   "/foobar",
		checker: APIStringResCheck{
			status: &IntEq{404},
		},
	})
	session.RESTCall(&APITest{
		method: "GET",
		path:   "/exit",
		checker: APIStringResCheck{
			status: &IntEq{200},
			body:   &StringEq{fmt.Sprintf("Goodbye from %v\n", serveraddr)},
		},
	})
}
