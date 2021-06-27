package main

import (
	"html/template"
	"os"
)

func main() {
	cal := &Calculation{
		Operator:  "-",
		MemberOne: 5,
		MemberTwo: 10,
		Result:    15,
	}

	tmpl, _ := template.New("output").Parse(templating)

	err := tmpl.Execute(os.Stdout, cal)
	if err != nil {
		panic(err)
	}
}

type Calculation struct {
	Operator  string
	MemberOne int
	MemberTwo int
	Result    int
}

const templating = "Calculation: {{.MemberOne}}{{.Operator}}{{.MemberTwo}}={{.Result}} \nPayment Cost: 1msat \n"
