package common

type AppError struct {
	message string
	code    int
}

var (
	INTERNAL_SERVER_ERROR = AppError{message: "Internal server error", code: 500}
)
