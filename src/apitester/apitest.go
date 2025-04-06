package apitester

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

type APIResultChecker interface {
	CheckRes(res *http.Response) (bool, error)
}

type APITest struct {
	method  string
	path    string
	checker APIResultChecker
}

type APIStringResCheck struct {
	body       string
	statuscode int
}

func (params APIStringResCheck) CheckRes(res *http.Response) (success bool, err error) {
	if params.statuscode != 0 {
		if params.statuscode != res.StatusCode {
			return false, fmt.Errorf("statuscode is %v not %v", res.StatusCode, params.statuscode)
		}
	}

	if params.body != "" {
		buf := new(strings.Builder)
		io.Copy(buf, res.Body)
		resbody := buf.String()

		if resbody != params.body {
			return false, fmt.Errorf("message body is '%v' not '%v'", resbody, params.body)
		}
	}
	return true, nil
}
