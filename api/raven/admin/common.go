package admin

import (
	"io/ioutil"
	"net/http"

	"github.com/palantir/stacktrace"
)

func getQuery(r *http.Request) ([]byte, error) {
	contentType := r.Header.Get("Content-Type")
	if contentType == "application/json" {
		query, err := ioutil.ReadAll(r.Body)
		return query, stacktrace.Propagate(err, "cannot read body content")
	}
	err := r.ParseForm()
	if err != nil {
		return nil, stacktrace.Propagate(err, "cannot parse form")
	}
	return []byte(r.Form.Get("query")), nil
}
