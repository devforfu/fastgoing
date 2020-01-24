package fastgoing

import (
    "fmt"
    "os"
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

func TestDefaultEnv_VariableWasSet(t *testing.T) {
    defer func() { _ = os.Unsetenv("test") }()
    _ = os.Setenv("test", "test")
    value := DefaultEnv("test", "fallback")
    if value != "test" {
        t.Error("wrong env variable value")
    }
}

func TestDefaultEnv_VariableIsMissing(t *testing.T) {
    value := DefaultEnv("test", "fallback")
    if value != "fallback" {
        t.Error("wrong env variable value")
    }
}