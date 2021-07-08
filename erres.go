package erres

import (
	"text/template"
)

const rawErrorTemplate = "[{{ .Time}}] {{ .Error}} | {{ .Function}}.{{ .Name}} | {{ .Description}}"

func init() {
	var err error

	errorTemplate, err = template.New("error").Parse(rawErrorTemplate)
	if err != nil {
		panic(err.Error())
	}
}

var errorTemplate *template.Template
var timeFormat = "2006-01-02 15:04:05"

func SetErrorFormat(t *template.Template) error {
	if t == nil {
		return NilArgument
	} else {
		errorTemplate = t
		return nil
	}
}

func SetTimeFormat(f string) error {
	if f == "" {
		return InvalidArgument
	} else {
		timeFormat = f
		return nil
	}
}
