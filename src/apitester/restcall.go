package apitester

import (
	"fmt"
	"log"
	"net/http"
)

type RESTSession struct {
	opts   *Options
	client *http.Client
}

func NewSession(opts *Options) *RESTSession {
	return &RESTSession{
		client: &http.Client{},
		opts:   opts,
	}
}

func (session *RESTSession) RESTCall(test *APITest) {

	req, _ := http.NewRequest(test.method, session.opts.BaseURL+test.path, nil)

	if session.opts.Verbose {
		fmt.Printf("REQ:  %s %s\n", test.method, test.path)
	}

	resp, _ := session.client.Do(req)

	if resp != nil {

		if session.opts.Verbose {
			statuscode := resp.Status

			fmt.Printf("RESP:  (%s)\n", statuscode)
		}

		if test.checker != nil {
			OK, err := test.checker.CheckRes(resp)
			if !OK {
				log.Fatalf("Failed to run test: %v", err)
			}
		}
	}

}
