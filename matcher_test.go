package errors

import (
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIs(t *testing.T) {
	k := New("foo")
	err := k.New("bar")

	assert.Equal(t, Is(err, k), true)
	assert.Equal(t, Is(io.EOF, k), false)
}

func TestIsMultiple(t *testing.T) {
	k1 := New("foo")
	err1 := k1.New("bar")

	k2 := New("qux")
	err2 := k2.Wrap(err1)

	assert.Equal(t, Is(err2, k1, k2), true)
}

func TestAny(t *testing.T) {
	k1 := New("foo")
	err := k1.New("bar")

	k2 := New("qux")

	assert.Equal(t, Any(err, k1, k2), true)
	assert.Equal(t, Any(io.EOF, k1, k2), false)
}
