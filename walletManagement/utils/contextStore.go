package utils

var currentContextTracer string

// TODO: seems no race condition =>
// but still quite not clean if this the best way to pass on e2e values between request and response

func AddValue(value string) {
	currentContextTracer = value
}

func ReadValue() string {
	return currentContextTracer
}
