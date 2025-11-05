package errorsx

import (
	"errors"
	"fmt"
)

type ErrorX struct {
	Code    int    `json:"code,omitempty"`
	Reason  string `json:"reason,omitempty"`
	Message string `json:"message,omitempty"`
}

func (err *ErrorX) Error() string {
	return fmt.Sprintf("error: code = %d reason = %s message = %s",
		err.Code, err.Reason, err.Message)
}

func New(code int, reason string, format string, args ...any) *ErrorX {
	return &ErrorX{
		Code:    code,
		Reason:  reason,
		Message: fmt.Sprintf(format, args...),
	}
}

// WithMessage is to set the Message field in struct
func (err *ErrorX) WithMessage(format string, args ...any) *ErrorX {
	err.Message = fmt.Sprintf(format, args...)
	return err
}

// FromError To convert a standard error to an ErrorX
func FromError(err error) *ErrorX {
	if err == nil {
		return nil
	}

	if errx := new(ErrorX); errors.As(err, &errx) {
		return errx
	}

	// By default, the program will return an unknown error which represents the server side goes wrong
	return New(ErrInternal.Code, ErrInternal.Reason, err.Error())
}
