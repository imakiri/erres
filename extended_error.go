package erres

import (
	"bytes"
	"time"
)

type extendedError struct {
	previous    *extendedError
	error       Error
	time        int64
	name        string
	function    string
	description string
}

func (e extendedError) String() string {
	var b = new(bytes.Buffer)
	var es = struct {
		Error       string
		Time        string
		Function    string
		Name        string
		Description string
	}{
		Error:       string(e.error),
		Time:        time.Unix(0, e.time).Format(timeFormat),
		Function:    e.function,
		Name:        e.name,
		Description: e.description,
	}

	var err = errorTemplate.Execute(b, es)
	if err != nil {
		return err.Error()
	}

	return b.String()
}

func (e extendedError) Error() string {
	return string(e.error)
}

func (e extendedError) Copy() *extendedError {
	var ee = new(extendedError)
	*ee = e
	return ee
}

func (e extendedError) Equal(err Error) bool {
	if e.error == err {
		return true
	}
	return false
}

// Returns true if the previous is nil
func (e extendedError) Last() bool {
	if e.previous == nil {
		return true
	} else {
		return false
	}

}

// Goes to the previous element in the linked list
func (e *extendedError) Previous() {
	if e.previous != nil {
		*e = *e.previous
	}
}

// Links the given error as an ExternalError.
func (e *extendedError) Link(err error) *extendedError {
	switch err := err.(type) {
	case *extendedError:
		e.previous = err
		return e
	default:
		if err == nil {
			return e
		}
		var pe = &extendedError{error: ExternalError, description: err.Error()}
		e.previous = pe
		return e
	}
}

func (e *extendedError) SetName(name string) *extendedError {
	e.name = name
	return e
}

func (e *extendedError) SetDescription(description string) *extendedError {
	e.description = description
	return e
}
