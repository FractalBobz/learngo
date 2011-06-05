package main

import (
	"http"
	"json"
	"os"
	"strings"
)

const runUrl = "http://golang.org/compile?output=json"

type runResult struct {
	Output string "output"
	Errors string "compile_errors"
}

func run(code string) (*runResult, os.Error) {
	r, err := http.DefaultClient.Post(runUrl, "text/plain", strings.NewReader(code))
	if err != nil {
		return nil, err
	}
	defer r.Body.Close()
	dec := json.NewDecoder(r.Body)
	res := new(runResult)
	if err = dec.Decode(res); err != nil {
		return nil, err
	}
	return res, nil
}
