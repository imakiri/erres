package erres

import (
	"runtime"
	"time"
)

type Error string

const (
	USOE                               = UnacceptableStateOfExecution
	UnacceptableStateOfExecution Error = "unacceptable state of execution"
	InternalServiceError         Error = "internal service error"
	InvalidArgument              Error = "invalid argument"
	InvalidMethod                Error = "invalid method"
	NilArgument                  Error = "nil argument"
	NilReturnValue               Error = "nil return value"
	InvalidReturnValue           Error = "invalid return value"
	InvalidReturnType            Error = "invalid return type"
	ConnectionError              Error = "connection error"
	NotSupported                 Error = "not supported"
	Overdue                      Error = "overdue"
	Expired                      Error = "expired"
	ClosedChannel                Error = "closed channel"
	TypeMismatch                 Error = "type mismatch"
	AccessDenied                 Error = "access denied"
	NotFound                     Error = "not found"
	AlreadyExist                 Error = "already exist"
	FileError                    Error = "file error"
	SerializationError           Error = "serialization error"
	DeserializationError         Error = "deserialization error"
	ExternalError                Error = "external error"
	JustError                    Error = "error"
	UnknownError                 Error = "unknown error"
)

func (e Error) Error() string {
	return string(e)
}

// Extends Error to extendedError.
// The skip value tells how many function calls to skip before getting the function name
func (e Error) Extend(skip int) *extendedError {
	return &extendedError{error: e, time: time.Now().UnixNano(), function: funcName(skip)}
}

// Extends Error to extendedError and links the given error as an ExternalError.
// The skip value tells how many function calls to skip before getting the function name
func (e Error) ExtendAndLink(skip int, err error) *extendedError {
	switch err := err.(type) {
	case *extendedError:
		return &extendedError{previous: err, error: e, time: time.Now().UnixNano(), function: funcName(skip)}
	case error:
		if err == nil {
			return &extendedError{error: e, time: time.Now().UnixNano(), function: funcName(skip)}
		}
		var pe = &extendedError{error: ExternalError, description: err.Error()}
		return &extendedError{previous: pe, error: e, time: time.Now().UnixNano(), function: funcName(skip)}
	default:
		return &extendedError{error: e, time: time.Now().UnixNano(), function: funcName(skip)}
	}
}

func funcName(i int) string {
	var pc = make([]uintptr, 15)
	var n = runtime.Callers(3+i, pc)
	var frames = runtime.CallersFrames(pc[:n])
	var frame, _ = frames.Next()
	return frame.Function
}
