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
	body   StringValueCheck
	status IntValueCheck
}

func (params APIStringResCheck) CheckRes(res *http.Response) (success bool, err error) {
	if params.status != nil {
		ok, err := params.status.CheckVal(res.StatusCode)
		if !ok {
			return false, fmt.Errorf("statuscode check failed: %e", err)
		}
	}

	if params.body != nil {
		buf := new(strings.Builder)
		io.Copy(buf, res.Body)
		resbody := buf.String()

		ok, err := params.body.CheckVal(&resbody)
		if !ok {
			return false, fmt.Errorf("message body check failed: %e", err)
		}
	}
	return true, nil
}
