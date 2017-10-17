package server

import (
	"fmt"
	"net/http"
	"net/http/httputil"
)

type ReqDumper struct {
	port int
}

func NewReqDumper(port int) *ReqDumper {
	return &ReqDumper{port: port}
}

func (dumper *ReqDumper) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")

	output, err := httputil.DumpRequest(r, true)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	fmt.Fprintln(w, string(output))
}
