package error

import (
	"net/http"

	"github.com/fenky-ng/edot-test-case/user/internal/constant"
)

func GetHttpCodeAndErrorMessage(err error) (code int, message string) {
	if err == nil {
		return
	}

	err = unwrapAndFoundError(err)

	code, exists := errorHttpCodeMap[err]
	if !exists {
		code = http.StatusInternalServerError
	}

	message, exists = errorMessage[err]
	if !exists {
		message = constant.GeneralErrorMessage
	}

	return
}

func unwrapAndFoundError(err error) error {
	switch x := err.(type) {

	case interface{ Unwrap() error }:
		err = x.Unwrap()
		if err != nil {
			// check recursively if is wrapped error
			return unwrapAndFoundError(err)
		}

	case interface{ Unwrap() []error }:
		// return first error found in errorHttpCodeMap
		wrappedErrors := x.Unwrap()
		for _, wrappedError := range wrappedErrors {
			we := unwrapAndFoundError(wrappedError)
			_, ok := errorHttpCodeMap[we]
			if ok {
				return we
			}
		}

	default:
		return err

	}

	return err
}
