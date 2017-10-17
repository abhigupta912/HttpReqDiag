package server

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

type ReqTester struct {
	port int
}

func NewReqTester(port int) *ReqTester {
	return &ReqTester{port: port}
}

func (t *ReqTester) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprintf(w, t.parseReq(r))
}

func (t *ReqTester) parseReq(r *http.Request) string {
	output := bytes.NewBufferString("Parsing request\n")

	output.WriteString("Proto: ")
	output.WriteString(r.Proto)
	output.WriteString("\n")

	output.WriteString("Origin: ")
	output.WriteString(r.RemoteAddr)
	output.WriteString("\n")

	output.WriteString("URL: ")
	output.WriteString(r.URL.String())
	output.WriteString("\n")

	output.WriteString("Method: ")
	output.WriteString(r.Method)
	output.WriteString("\n")

	output.WriteString("Host: ")
	output.WriteString(r.Host)
	output.WriteString("\n")

	output.WriteString("Headers: ")
	numHeaders := len(r.Header)
	if numHeaders > 0 {
		for header, value := range r.Header {
			output.WriteString("\n\t")
			output.WriteString(header)
			output.WriteString(" : ")
			valueBytes, valueParseErr := json.Marshal(value)
			if valueParseErr != nil {
				output.WriteString("Unable to parse header value")
			} else {
				output.Write(valueBytes)
			}
		}
	} else {
		output.WriteString("None")
	}
	output.WriteString("\n")

	output.WriteString("Cookies: ")
	cookies := r.Cookies()
	numCookies := len(cookies)
	if numCookies > 0 {
		for _, cookie := range cookies {
			output.WriteString("\n\t")
			output.WriteString(cookie.String())
		}
	} else {
		output.WriteString("None")
	}
	output.WriteString("\n")

	output.WriteString("Transfer Encoding: ")
	txEncodingBytes, txEncodingParseErr := json.Marshal(r.TransferEncoding)
	if txEncodingParseErr != nil {
		output.WriteString("Unable to parse Transfer Encoding")
	} else {
		output.Write(txEncodingBytes)
	}
	output.WriteString("\n")

	output.WriteString("Content Length: ")
	output.WriteString(strconv.FormatInt(r.ContentLength, 10))
	output.WriteString("\n")

	output.WriteString("Content: ")
	body, readErr := ioutil.ReadAll(r.Body)
	if readErr != nil {
		output.WriteString("Unable to read Request Body")
	} else {
		output.Write(body)
	}
	output.WriteString("\n")

	return output.String()
}
