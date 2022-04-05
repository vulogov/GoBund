package stdlib

import (
  "github.com/gammazero/deque"
  "github.com/zach-klippenstein/goregen"
  tc "github.com/vulogov/ThreadComputation"
)

func BUNDrandomString(l *tc.TCExecListener, name string, q *deque.Deque) (interface{}, error) {
  res, err := regen.Generate("[A-Za-z]{16,64}")
  if err != nil {
    return nil, err
  }
  return res, nil
}


func init() {
  tc.SetCommand("random.String", BUNDrandomString)
}
