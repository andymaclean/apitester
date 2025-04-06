package apitester

import (
	"flag"
	"fmt"
)

type Options struct {
	BaseURL  string
	TestFile string
	Verbose  bool
}

func DefaultOptions() (o Options) {
	return Options{
		BaseURL:  "https://localhost:443",
		TestFile: "test.yaml",
		Verbose:  false,
	}
}

func (o *Options) ReadCommandArguments() {
	BaseURL := flag.String("baseurl", "", "Base URL to query")
	TestFile := flag.String("testfile", "", "Test file")
	Verbose := flag.Bool("verbose", false, "Verbose output (default)")
	Silent := flag.Bool("silent", false, "Silent output")

	flag.Parse()

	if (*BaseURL) != "" {
		o.BaseURL = *BaseURL
	}

	if (*TestFile) != "" {
		o.TestFile = *TestFile
	}

	if *Verbose {
		o.Verbose = true
	}

	if *Silent {
		o.Verbose = false
	}

}

func (o *Options) Summary() string {
	return fmt.Sprintf("Base URL: %s\nTest file: %s\nVerbose %t\n", o.BaseURL, o.TestFile, o.Verbose)
}
