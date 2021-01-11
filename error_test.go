package errors

import (
	"errors"
	"testing"
)

func TestError(t *testing.T) {
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
	var errStruct *Error
	if errors.As(newErr, &errStruct) {
		errVal := errStruct.Unwrap()
		if errVal.Error() != err.Error() {
			t.Errorf("unwrapped error value %+v does not equal original error value %+v", errVal.Error(), err.Error())
		}

		if errStruct.Line != 20 {
			t.Errorf("incorrect error line: %+v", errStruct.Line)
		}

		if errStruct.File != "error_test.go" {
			t.Errorf("incorrect file name: %+v", errStruct.File)
		}

		if errStruct.Function != "TestError" {
			t.Errorf("incorrect function name: %+v", errStruct.Function)
		}
	} else {
		t.Error("invalid error type assertion")
	}
}
