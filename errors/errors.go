package errors

import "fmt"

type ErrorResponse struct {
	StatusCode int
	Code       string
	Reason     string
}

func (e ErrorResponse) Error() string {
	return fmt.Sprintf("{\n code: %v \n reason: %v\n}", e.Code, e.Reason)

}
