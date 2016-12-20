package errors

import (
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	k := New("foo")
	assert.Equal(t, k.Message, "foo")
}

func TestKindNew(t *testing.T) {
	k := New("foo")
	err := k.New()
	assert.Equal(t, err.Message, "foo")
	assert.Equal(t, err.Kind, k)
}

func TestKindNewWithFormat(t *testing.T) {
	k := New("foo %s")
	err := k.New("bar")
	assert.Equal(t, err.Message, "foo bar")
	assert.Equal(t, err.Kind, k)
}

func TestKindWrap(t *testing.T) {
	k := New("foo")
	err := k.Wrap(io.EOF)
	assert.Equal(t, err.Message, "foo: %s")
	assert.Equal(t, err.Kind, k)
	assert.Equal(t, err.Cause, io.EOF)
}

func TestKindIs(t *testing.T) {
	k := New("foo")
	err := k.New("bar")
	assert.Equal(t, k.Is(err), true)
	assert.Equal(t, k.Is(io.EOF), false)
	assert.Equal(t, k.Is(nil), false)
}

func TestKindIsChildren(t *testing.T) {
	k := New("foo")
	err := k.Wrap(io.EOF)
	assert.Equal(t, k.Is(err), true)
}

func TestError(t *testing.T) {
	err := New("foo %s").New("bar")
	assert.Equal(t, err.Error(), "foo bar")
}

func TestErrorCause(t *testing.T) {
	err := New("foo").Wrap(io.EOF)
	assert.Equal(t, err.Error(), "foo: EOF")
}
