package erres

import "time"

var FormatTime = "2006-01-02 15:04:05"

type err string

func (e err) Error() string {
	return string(e)
}

func (e err) Extend() Error {
	return Error{err: e, time: time.Now().Format(FormatTime)}
}

const (
	InvalidArgument      err = "invalid argument"
	ConnectionError      err = "connection error"
	NotSupported         err = "not supported"
	ClosedChannel        err = "closed channel"
	TypeMismatch         err = "type mismatch"
	AccessDenied         err = "access denied"
	NotFound             err = "not found"
	AlreadyExist         err = "already exist"
	InternalServiceError err = "internal service error"
	SerializationError   err = "serialization error"
	DeserializationError err = "deserialization error"
	JustError            err = "error"
)
