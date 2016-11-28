package errors

import (
	"fmt"
)

func New(msg string) *Kind {
	return &Kind{Message: msg}
}

type Kind struct {
	Message string
}

func (k *Kind) New(values ...interface{}) *Error {
	return &Error{
		Kind:    k,
		Message: fmt.Sprintf(k.Message, values...),
	}
}

func (k *Kind) Wrap(err error) *Error {
	return &Error{
		Kind:    k,
		Child:   err,
		Message: k.Message + ": %s",
	}
}

func (k *Kind) Match(required ...*Kind) bool {

	for _, kind := range required {
		if k == kind {
			return true
		}
	}

	return false
}

type Error struct {
	Kind    *Kind
	Child   error
	Message string
}

func (err *Error) Error() string {
	if err.Child == nil {

		return err.Message
	}

	return fmt.Sprintf(err.Message, err.Child.Error())
}
