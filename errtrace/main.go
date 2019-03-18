package errtrace

import (
  "fmt"
  "strings"
  "runtime"
)

const DefaultCap = 20

type Frame struct {
	Func string
	Line int
	Path string
}

func (f Frame) String() string {
  return fmt.Sprintf("%s:%d %s", f.Path, f.Line, f.Func)
}

type Error struct {
  Err error
  Frames []Frame
}

func New(message string) *Error {
	return trace(fmt.Errorf(message), 2)
}

func Wrap(err interface{}) *Error {
	if err == nil {
		return nil
	}

  var resultingError error

  switch e := err.(type) {
  case *Error:
    return e
  case error:
    resultingError = e
  default:
    resultingError = fmt.Errorf("%v", e)
  }

	return trace(resultingError, 2)
}

func (e *Error) Error() string {
	if e == nil {
		return ""
	}

	return e.Err.Error()
}

func (e *Error) StringStack() string {
	if e == nil {
		return ""
	}

  frames := make([]string, 0, len(e.Frames))

  for _, frame := range e.Frames {
    frames = append(frames, frame.String())
  }

	return strings.Join(frames, ", ")
}

func trace(err error, skip int) *Error {
	frames := make([]Frame, 0, DefaultCap)

	for {
		pc, path, line, ok := runtime.Caller(skip)
		if !ok {
			break
		}

		fn := runtime.FuncForPC(pc)

		frame := Frame{
			Func: fn.Name(),
			Line: line,
			Path: path,
		}

		frames = append(frames, frame)

		skip++
	}

	return &Error{
		Err:    err,
		Frames: frames,
	}
}
