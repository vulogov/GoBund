package stdlib

import (
  tc "github.com/vulogov/ThreadComputation"
)

const (
  RandomNormal    = 500
)

func BUND_Type(x interface{}) int {
  if x == nil {
    return tc.Nil
  }
  switch x.(type) {
  case *BUND_Random_Distribution_Normal:
    return RandomNormal
  }
  return tc.Unknown
}

func BUND_type_to_str(t interface{}) string {
  switch t.(type) {
  case int:
    switch t {
    case RandomNormal:
      return "RandomNormalDistribution"
    }
  }
  return "Unknown"
}

func init() {
  tc.SetExternalTypeHandlers(BUND_Type, BUND_type_to_str)
}
