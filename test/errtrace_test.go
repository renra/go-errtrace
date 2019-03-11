package test

import (
  "testing"
  "errors"
  "strings"
  "app/errtrace"
)

func TestNew(t *testing.T) {
  err := errtrace.New("Baf")

  expected := "Baf"
  actual := err.Error()

  if actual != expected {
    t.Errorf("Expected err.Error() to equal %s, but is %s", expected, actual)
  }

  expectedStackSize := 3
  actualStackSize := len(err.Frames)

  if actualStackSize != expectedStackSize {
    t.Errorf("Expected len(err.Frames) to equal %v, but is %v", expectedStackSize, actualStackSize)
  }

  expectedStringStack := serializeFrames(err.Frames)
  actualStringStack := err.StringStack()

  if actualStringStack != expectedStringStack {
    t.Errorf(
      "Expected err.StringStack() to equal %v, but is %v",
      expectedStringStack,
      actualStringStack,
    )
  }
}

func TestWrap(t *testing.T) {
  err := errtrace.Wrap(errors.New("Baf"))

  expected := "Baf"
  actual := err.Error()

  if actual != expected {
    t.Errorf("Expected err.Error() to equal %s, but is %s", expected, actual)
  }

  expectedStackSize := 3
  actualStackSize := len(err.Frames)

  if actualStackSize != expectedStackSize {
    t.Errorf("Expected len(err.Frames) to equal %v, but is %v", expectedStackSize, actualStackSize)
  }

  expectedStringStack := serializeFrames(err.Frames)
  actualStringStack := err.StringStack()

  if actualStringStack != expectedStringStack {
    t.Errorf(
      "Expected err.StringStack() to equal %v, but is %v",
      expectedStringStack,
      actualStringStack,
    )
  }
}

func serializeFrames(frames []errtrace.Frame) string {
  fs := make([]string, 0, len(frames))

  for _, f := range frames {
    fs = append(fs, f.String())
  }

  return strings.Join(fs, ", ")
}

