package api

import "fmt"

type ApiError struct {
	message    string
	statusCode int
}

//type AuthError struct {
//        *ApiError
//}

func newApiError(statusCode int, message string) error {
	return &ApiError{message, statusCode}
}

func (err *ApiError) Error() string {
	var message string
	if err.message == "" {
		message = "<No message>"
	} else {
		message = err.message
	}

	return fmt.Sprintf("%v Error: %s", err.statusCode, message)
}
