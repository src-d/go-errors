package errors

import (
	"fmt"
)

// Kind represents a kind of an error, from a Kind you can generate as many
// Error instances of you want of this Kind
type Kind struct {
	Message string
}

// New returns a Kind with the given msg
func New(msg string) *Kind {
	return &Kind{Message: msg}
}

// New returns a new Error, values can be pass to it if the Kind was created
// using an printf format
func (k *Kind) New(values ...interface{}) *Error {
	return &Error{
		Kind:    k,
		Message: fmt.Sprintf(k.Message, values...),
	}
}

// Wrap create a new Error of this Kind with cause error
func (k *Kind) Wrap(cause error) *Error {
	return &Error{
		Kind:    k,
		Cause:   cause,
		Message: k.Message + ": %s",
	}
}

// Is checks if the given error or any of his children are of this Kind
func (k *Kind) Is(err error) bool {
	if err == nil {
		return false
	}

	e, ok := err.(*Error)
	if !ok {
		return false
	}

	if k == e.Kind {
		return true
	}

	if e.Cause == nil {
		return false
	}

	return k.Is(e.Cause)
}

// Error represents an error of some Kind, implements the error interface
type Error struct {
	// Kind is a pointer to the Kind of this error
	Kind *Kind
	// Cause of the error
	Cause error
	// Message describing the error
	Message string
}

func (err *Error) Error() string {
	if err.Cause == nil {
		return err.Message
	}

	return fmt.Sprintf(err.Message, err.Cause.Error())
}
