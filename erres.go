package erres

import (
	"bytes"
	"text/template"
)

const errorTemplate = "[{{ .Time}}] {{ .Err}} | root{{range .Routs}}.{{.}}{{end}} | description{{range .Desc}}:{{.}}{{end}}"

func init() {
	var err error

	FormatError, err = template.New("error").Parse(errorTemplate)
	if err != nil {
		panic(err.Error())
	}
}

var FormatError *template.Template

type Error struct {
	err   err
	time  string
	routs []string
	desc  []string
}

func (e Error) Error() string {
	var b = new(bytes.Buffer)
	var err = FormatError.Execute(b, e.ext())
	if err != nil {
		return err.Error()
	}

	return b.String()
}

func (e Error) AddRoute(route string) Error {
	e.routs = append(e.routs, route)
	return e
}

func (e Error) AddDescription(description string) Error {
	e.desc = append(e.desc, description)
	return e
}

func (e Error) ext() struct {
	Err   err
	Time  string
	Routs []string
	Desc  []string
} {
	return struct {
		Err   err
		Time  string
		Routs []string
		Desc  []string
	}{
		Err:   e.err,
		Time:  e.time,
		Routs: e.routs,
		Desc:  e.desc,
	}
}

func Compare(err error, base err) bool {
	var e, ok = err.(Error)
	if !ok {
		return false
	}

	if e.err == base {
		return true
	}

	return false
}
