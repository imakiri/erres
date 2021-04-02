package erres

import (
	"bytes"
	"time"
)

type Error struct {
	last  *Error
	ce    Consterror
	time  int64
	fname string
	ename string
	desc  string
}

func (e Error) String() string {
	var b = new(bytes.Buffer)
	var es = struct {
		Err   string
		Time  string
		Fname string
		Ename string
		Desc  string
	}{
		Err:   string(e.ce),
		Time:  time.Unix(0, e.time).Format(fTime),
		Fname: e.fname,
		Ename: e.ename,
		Desc:  e.desc,
	}

	var err = fError.Execute(b, es)
	if err != nil {
		return err.Error()
	}

	return b.String()
}

func (e Error) Error() string {
	return string(e.ce)
}

func (e Error) IsRoot() bool {
	if e.last == nil {
		return true
	} else {
		return false
	}
}

func (e *Error) Last() {
	*e = *e.last
}

// Extend and link err to method receiver and return pointer to an extended err
func (e *Error) Link(err CE) *Error {
	return err.ExtendAndLink(0, e)
}

func (e *Error) SetName(name string) *Error {
	e.ename = name
	return e
}

func (e *Error) SetDescription(description string) *Error {
	e.desc = description
	return e
}
