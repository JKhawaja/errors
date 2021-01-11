package errors

import (
	"errors"
	"testing"
)

func TestTrace(t *testing.T) {
	err := errors.New("this is an error")
	scope := map[string]interface{}{
		"var1": 1,
		"var2": "value",
		"var3": struct {
			Stuff string
		}{
			Stuff: "stuff",
		},
	}

	newErr := New(err, scope)
	newErrWrapped := New(newErr, nil)

	tr := NewTrace(newErrWrapped)
	trace, ok := tr.(*Trace)
	if !ok {
		t.Error("trace type assertion error")
	}

	if len(trace.trace) != 2 {
		t.Errorf("trace length was not 2: %+v", len(trace.trace))
	}

	err0, ok := trace.trace[0].(*Error)
	if !ok {
		t.Error("trace 0 error type assertion error")
	}

	if err0.Timestamp != nil {
		t.Errorf("timestamp redundant")
	}

	err1, ok := trace.trace[1].(*Error)
	if !ok {
		t.Error("trace 1 error type assertion error")
	}

	if err1.Timestamp == nil {
		t.Errorf("no timestamp in trace")
	}
}
