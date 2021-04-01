package erres

import (
	"runtime"
	"time"
)

type Consterror string

type ce = Consterror
type CE = Consterror

const (
	InternalServiceError ce = "internal service error"
	InvalidArgument      ce = "invalid argument"
	NilArgument          ce = "nil argument"
	FileError            ce = "file error"
	ConnectionError      ce = "connection error"
	NotSupported         ce = "not supported"
	ClosedChannel        ce = "closed channel"
	TypeMismatch         ce = "type mismatch"
	AccessDenied         ce = "access denied"
	NotFound             ce = "not found"
	AlreadyExist         ce = "already exist"
	SerializationError   ce = "serialization error"
	DeserializationError ce = "deserialization error"
	JustError            ce = "error"
)

func (i CE) Error() string {
	return string(i)
}

func (i CE) Extend() *Error {
	return &Error{ce: i, time: time.Now().UnixNano(), fname: funcName()}
}

// Extend and link ce to method receiver and return extended receiver
func (i CE) ExtendAndLink(err *Error) *Error {
	return &Error{last: err, ce: i, time: time.Now().UnixNano(), fname: funcName()}
}

func funcName() string {
	pc := make([]uintptr, 15)
	n := runtime.Callers(3, pc)
	frames := runtime.CallersFrames(pc[:n])
	frame, _ := frames.Next()
	return frame.Function
}
