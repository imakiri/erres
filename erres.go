package erres

import (
	"text/template"
)

const errorTemplate = "[{{ .Time}}] {{ .Err}} | {{ .Fname}}.{{ .Ename}} | {{ .Desc}}"

func init() {
	var err error

	fError, err = template.New("error").Parse(errorTemplate)
	if err != nil {
		panic(err.Error())
	}
}

var fError *template.Template
var fTime = "2006-01-02 15:04:05"

func SetErrorFormat(t *template.Template) error {
	if t == nil {
		return NilArgument
	} else {
		fError = t
		return nil
	}
}

func SetTimeFormat(f string) error {
	if f == "" {
		return InvalidArgument
	} else {
		fTime = f
		return nil
	}
}
