package main

import (
	"bytes"
	"fmt"
	"http"
	"log"
	"template"
)

func main() {
	http.Handle("/static/", http.FileServer("static/", "/static/"))
	http.HandleFunc("/step", stepHandler)
	http.HandleFunc("/test", testHandler)
	http.ListenAndServe(":8080", nil)
}

var step1 = Step{
	"Double my input",
	"Write a function that doubles its input.",
	"func fn(i int) int {\n}",
	TestData{Setup: "n := 1", Args: "n", Expect: "2"},
}

func stepHandler(w http.ResponseWriter, r *http.Request) {
	err := stepTemplate.Execute(w, step1)
	if err != nil {
		log.Print(err)
	}
}

var stepTemplate = template.MustParseFile("tmpl/step.html", nil)

func testHandler(w http.ResponseWriter, r *http.Request) {
	// read user code
	userCode := new(bytes.Buffer)
	defer r.Body.Close()
	if _, err := userCode.ReadFrom(r.Body); err != nil {
		http.Error(w, err.String(), http.StatusInternalServerError)
		return
	}
	// put user code in harness
	t := step1.Test
	t.UserCode = userCode.String()
	code := new(bytes.Buffer)
	if err := TestHarness.Execute(code, t); err != nil {
		http.Error(w, err.String(), http.StatusInternalServerError)
		return
	}
	// run code
	out, err := run(code.String())
	if err != nil {
		http.Error(w, err.String(), http.StatusInternalServerError)
		return
	}
	if out.Output == t.Expect {
		fmt.Fprint(w, "Well done!")
	} else {
		fmt.Fprintf(w, "Not quite.\n\n%s", out.Errors)
	}
}
