package apitester

import (
	"fmt"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

func Run(opts *Options) {
	if opts.Verbose {
		fmt.Printf("APITester.\n%s", opts.Summary())
	}

	testfile, err := os.ReadFile(opts.TestFile)

	if err != nil {
		log.Fatalf("Failed to reat %v: %e", opts.TestFile, err)
	}

	var tests map[string]any

	err = yaml.Unmarshal(testfile, &tests)

	if err != nil {
		log.Fatalf("Failed to parse test data: %e", err)
	}

	fmt.Printf("%v", tests)

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
