# Errtrace

Errors with traces. Inspired heavily by [tracerr](https://github.com/ztrue/tracerr), only simplified and added string formatting of the whole trace.

## Usage

### Creating new errors

```
package main

import (
  "fmt"
  "github.com/renra/go-errtrace/errtrace"
)

func main() {
  err := errtrace.New("Something went wrong")

  fmt.Println(fmt.Sprintf("%v", err.Error()))
  fmt.Println(fmt.Sprintf("%v", err.Frames))
  fmt.Println(fmt.Sprintf("%v", err.StringStack()))
}

```

### Wrapping existing errors

```
package main

import (
  "fmt"
  "github.com/renra/go-errtrace/errtrace"
)

func main() {
  err := theCrasher()

  fmt.Println(fmt.Sprintf("%v", err.Error()))
  fmt.Println(fmt.Sprintf("%v", err.Frames))
  fmt.Println(fmt.Sprintf("%v", err.StringStack()))
}

func theCrasher() *errtrace.Error {
  return theCrasherDependent()
}

func theCrasherDependent() *errtrace.Error {
  return errtrace.Wrap(errors.New("Crash now"))
}
```
