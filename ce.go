package erres

import (
	"runtime"
	"time"
)

type Consterror string

type ce = Consterror
type CE = Consterror

const (
	USOE                            = UnacceptableStateOfExecution
	UnacceptableStateOfExecution ce = "unacceptable state of execution"
	InternalServiceError         ce = "internal service error"
	InvalidArgument              ce = "invalid argument"
	NilArgument                  ce = "nil argument"
	NilReturnValue               ce = "nil return value"
	InvalidReturnValue           ce = "invalid return value"
	ConnectionError              ce = "connection error"
	NotSupported                 ce = "not supported"
	ClosedChannel                ce = "closed channel"
	TypeMismatch                 ce = "type mismatch"
	AccessDenied                 ce = "access denied"
	NotFound                     ce = "not found"
	AlreadyExist                 ce = "already exist"
	FileError                    ce = "file error"
	SerializationError           ce = "serialization error"
	DeserializationError         ce = "deserialization error"
	JustError                    ce = "error"
)

func (i CE) Error() string {
	return string(i)
}

// Skip value tells how many func calls to skip before getting caller name
func (i CE) Extend(skip int) *Error {
	return &Error{ce: i, time: time.Now().UnixNano(), fname: funcName(skip)}
}

// Extend method receiver, link it to err and return extended receiver
func (i CE) ExtendAndLink(skip int, err *Error) *Error {
	return &Error{last: err, ce: i, time: time.Now().UnixNano(), fname: funcName(skip)}
}

func funcName(i int) string {
	pc := make([]uintptr, 15)
	n := runtime.Callers(3+i, pc)
	frames := runtime.CallersFrames(pc[:n])
	frame, _ := frames.Next()
	return frame.Function
}
