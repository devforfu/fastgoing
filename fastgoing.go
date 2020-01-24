package fastgoing

import "os"

type OnError func(interface{})

var errorHandler OnError = func(i interface{}) {
    panic(i)
}

// Check checks if the error is nil and panics if it is not.
func Check(err error) {
    if err != nil { errorHandler(err) }
}

// WorkDir is a version of os.Getwd that doesn't return error but panics instead
// if the working directory cannot be returned.
func WorkDir() string {
    cwd, err := os.Getwd()
    Check(err)
    return cwd
}

// DefaultEnv returns env variable if it present; otherwise, a fallback value
// is returned.
func DefaultEnv(name, fallback string) string {
    if value := os.Getenv(name); value == "" {
        return fallback
    } else {
        return value
    }
}