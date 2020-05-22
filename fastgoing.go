package fastgoing

import (
    "encoding/json"
    "os"
    "regexp"
    "strconv"
    "strings"
    "time"
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

func MustRegexpMap(pattern string) *RegexpMap {
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

// DateUTC create an instance of time.Time that doesn't contain time part but
// only date's components. The location is assigned to time.UTC.
func DateUTC(year, month, day int) time.Time {
    return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
}

// Exists checks if local file exists.
//
// There are situations when this function cannot return unambiguous result. If
// this case, the error should be inspected.
//
// Reference:
//     https://stackoverflow.com/questions/12518876/how-to-check-if-a-file-exists-in-go
func Exists(path string) (bool, error) {
    if _, err := os.Stat(path); err == nil {
        return true, nil
    } else if os.IsNotExist(err) {
        return false, nil
    } else {
        return false, err
    }
}

// Verbose returns configuration as a list of string, one line per property or
// as an array with a single element.
func Verbose(object interface{}, splitNewLine bool) (lines []string, err error) {
    indented, err := json.MarshalIndent(object, "", "\t")
    if err != nil { return nil, err }
    if splitNewLine {
        lines = strings.Split(string(indented), "\n")
    } else {
        lines = []string{string(indented)}
    }
    return lines, nil
}

func MustVerboseWithSplit(object interface{}) []string {
    lines, err := Verbose(object, true)
    Check(err)
    return lines
}