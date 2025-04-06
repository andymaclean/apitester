package apitester

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func RESTCall(opts *Options, method string, path string, body io.Reader) {
	client := &http.Client{}
	req, _ := http.NewRequest(method, opts.BaseURL+path, body)

	if opts.Verbose {
		fmt.Printf("REQ:  %s %s\n", method, path)
	}

	resp, _ := client.Do(req)

	var resbody string

	if resp != nil {
		buf := new(strings.Builder)
		io.Copy(buf, resp.Body)
		resbody = buf.String()
	}

	statuscode := resp.Status

	if opts.Verbose {
		fmt.Printf("RESP:  (%s) %s\n", statuscode, resbody)
	}
}
