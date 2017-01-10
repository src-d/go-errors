package errors

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCaller(t *testing.T) {
	s := NewStackTrace(0)
	assert.NotEmpty(t, s.StackTrace)

	assert.Equal(t, fmt.Sprintf("%s", s), "[stack_test.go testing.go asm_amd64.s]")
}

func TestCallerSkip(t *testing.T) {
	s := NewStackTrace(1)
	assert.NotEmpty(t, s.StackTrace)

	assert.Equal(t, fmt.Sprintf("%s", s), "[testing.go asm_amd64.s]")
}
