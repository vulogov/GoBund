package stdlib

import (
  "fmt"
  tc "github.com/vulogov/ThreadComputation"
  "github.com/gammazero/deque"
  "github.com/pieterclaerhout/go-log"
)

func StdlibFormat(q *deque.Deque) string {
  var args []interface{}

  if q.Len() > 0 {
    f := q.PopFront()
    switch f.(type) {
    case string:
      for q.Len() > 0 {
        v := q.PopFront()
        fun := tc.GetConverterCallback(v)
        if fun == nil {
          continue
        }
        args = append(args, fun(v, tc.String))
      }
    }
    return fmt.Sprintf(f.(string), args...)
  }
  return ""
}

func InitStdlib() {
  log.Debug("[ BUND ] standard library initialized")
}
