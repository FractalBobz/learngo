package main

import "template"

type Step struct {
	Title      string
	Body       string
	SampleCode string
	Test       TestData
}

type TestData struct {
	Setup    string
	Args     string
	Expect   string
	UserCode string
}

var TestHarness = template.New(nil)

func init() {
	TestHarness.SetDelims("{{", "}}")
	if err := TestHarness.Parse(harness); err != nil {
		panic(err)
	}
}

const harness = `package main

import "fmt"

func main() {
	{{Setup}}
	fmt.Print(fn({{Args}}))
}

{{UserCode}}
`
