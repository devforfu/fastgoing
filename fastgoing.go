package fastgoing

import (
    "os"
    "regexp"
    "strconv"
)

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

// RegexpMap converts results of regexp match into a map.
type RegexpMap struct {
    Compiled *regexp.Regexp
}

func MustRegexMap(pattern string) *RegexpMap {
    return &RegexpMap{Compiled:regexp.MustCompile(pattern)}
}

// Search matches string against compiled regexp and converts match results
// into a map of matched group.
func (r *RegexpMap) Search(value string) map[string]string {
    matched := r.Compiled.FindStringSubmatch(value)
    params := make(map[string]string)
    for i, name := range r.Compiled.SubexpNames() {
        if i > 0 && i <= len(matched) {
            params[name] = matched[i]
        }
    }
    return params
}

// MustInt wraps ParseInt function and panics if err is not nil.
func MustInt(number string) int {
    n, err := strconv.ParseInt(number, 10, 32)
    if err != nil { panic(err) }
    return int(n)
}
