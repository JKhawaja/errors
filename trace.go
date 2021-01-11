package errors

import (
	"encoding/json"
	"errors"
)

// Trace --
type Trace struct {
	trace []error
	val   error
}

// NewTrace --
func NewTrace(val error) error {
	t := &Trace{
		trace: make([]error, 0),
		val:   val,
	}

	return t.Unwrap()
}

// Error --
func (t *Trace) Error() string {
	data, _ := json.Marshal(t.trace)
	return string(data)
}

// Unwrap --
func (t *Trace) Unwrap() error {
	var errStruct *Error
	if errors.As(t.val, &errStruct) {
		var errVal *Error
		t.val = *errStruct.Value
		if errors.As(*errStruct.Value, &errVal) {
			errStruct.Timestamp = nil
			errStruct.Value = nil
		}
		t.trace = append(t.trace, errStruct)
		return t.Unwrap()
	}
	return t
}
