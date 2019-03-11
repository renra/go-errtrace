package main

import (
  "fmt"
  "errors"
  "app/errtrace"
)

func main() {
  err := errtrace.New("Something went wrong")

  fmt.Println(fmt.Sprintf("%v", err.Error()))
  fmt.Println(fmt.Sprintf("%v", err.Frames))
  fmt.Println(fmt.Sprintf("%v", err.StringStack()))

  fmt.Println("")

  err = theCrasher()

  fmt.Println(fmt.Sprintf("%v", err.Error()))
  fmt.Println(fmt.Sprintf("%v", err.Frames))
  fmt.Println(fmt.Sprintf("%v", err.StringStack()))

  fmt.Println("")

  defer func() {
    if e := recover(); e != nil {
      fmt.Println(fmt.Sprintf("%v", err.Error()))
      fmt.Println(fmt.Sprintf("%v", err.Frames))
      fmt.Println(fmt.Sprintf("%v", err.StringStack()))
    }
  }()

  panic(err)
}

func theCrasher() *errtrace.Error {
  return theCrasherDependent()
}

func theCrasherDependent() *errtrace.Error {
  return errtrace.Wrap(errors.New("Crash now"))
}
