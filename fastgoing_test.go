package fastgoing

import (
    "fmt"
    "testing"
)

func TestCheck(t *testing.T) {
    var defaultHandler = errorHandler
    defer func() {errorHandler = defaultHandler}()
    cases := []error{fmt.Errorf("error"), nil}
    for _, testCase := range cases {
        errorHandler = func(x interface{}) {
            err := x.(error)
            if err == nil {
                t.Error("failed test case: error is nil but callback was invoked")
            }
        }
        Check(testCase)
    }
}