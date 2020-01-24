package fastgoing

type OnError func(interface{})

var errorHandler OnError = func(i interface{}) {
    panic(i)
}

// Check checks if the error is nil and panics if it is not.
func Check(err error) {
    if err != nil { errorHandler(err) }
}

