package erres

import (
	"database/sql"
	"fmt"
	"testing"
)

func TestError(t *testing.T) {
	var err = new(extendedError)
	var e error

	e = sql.ErrTxDone
	if e != nil {
		err = InternalServiceError.ExtendAndLink(0, e).SetName("Foo")
	}

	e = sql.ErrNoRows
	if e != nil {
		err = InvalidArgument.ExtendAndLink(0, err).SetName("Bar")
	}

	e = nil
	if e != nil {
		err = InvalidArgument.ExtendAndLink(0, err).SetName("Buzz")
	}

	var ee = err.Copy()
	for ; !err.Last(); err.Previous() {
		fmt.Println(err.Error())
	}

	if !ee.Equal(InvalidArgument) {
		t.Error("not InvalidArgument")
	}
	ee.Previous()
	if !ee.Equal(InternalServiceError) {
		t.Error("not InternalServiceError")
	}
	ee.Previous()
	if !ee.Equal(ExternalError) {
		t.Error("not ExternalError")
	}
}
