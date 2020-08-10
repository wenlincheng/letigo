package baseerror

import "errors"

type BaseError struct {
	Err  error `json:"err"`
	Code int   `json:"code"`
}

func NewBaseError(code int, errorMsg string) BaseError {
	return BaseError{Err: errors.New(errorMsg), Code: code}
}

func (b BaseError) Error() string {
	return b.Err.Error()
}
