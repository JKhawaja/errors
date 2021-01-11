package errors

import (
	"encoding/json"
	"errors"
	"path/filepath"
	"runtime"
	"strings"
	"time"
)

// Error represents a structured error object
type Error struct {
	File      string                 `json:"file"`
	Function  string                 `json:"func"`
	Line      int                    `json:"line"`
	Package   string                 `json:"pkg"`
	Timestamp *time.Time             `json:"time,omitempty"`
	Value     *error                 `json:"value,omitempty"`
	Scope     map[string]interface{} `json:"scope,omitempty"`
}

// Error --
func (e *Error) Error() string {
	data, _ := json.Marshal(e)
	return string(data)
}

// Unwrap will return the error value
func (e *Error) Unwrap() error {
	var errStruct *Error
	errVal := *e.Value
	if errors.As(errVal, &errStruct) {
		return errStruct
	}
	return errVal
}

// New will create a structured error object and return it as an `error` value.
func New(value error, scope map[string]interface{}) error {
	var pkg, function string
	pkgAndfunction := strings.Split(filepath.Base(frame(1).Function), ".")
	if len(pkgAndfunction) < 2 {
		pkg = ""
		function = ""
	} else {
		pkg = pkgAndfunction[0]
		function = strings.Join(pkgAndfunction[1:], ".")
	}

	t := time.Now()

	return &Error{
		Timestamp: &t,
		Value:     &value,
		Package:   pkg,
		Function:  function,
		File:      filepath.Base(frame(1).File),
		Line:      frame(1).Line,
		Scope:     scope,
	}
}

func frame(skipFrames int) runtime.Frame {
	// always skip runtime.Callers() and frame()
	targetIdx := skipFrames + 2
	counters := make([]uintptr, targetIdx+2)
	n := runtime.Callers(0, counters)

	var frame runtime.Frame
	if n > 0 {
		frames := runtime.CallersFrames(counters[:n])
		var idx int
		for idx <= targetIdx {
			candidate, more := frames.Next()
			if !more {
				break
			}

			if idx == targetIdx {
				frame = candidate
			}
			idx++
		}
	}

	return frame
}
