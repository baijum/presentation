package simple

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSomething(t *testing.T) {

	assert.Equal(t, "hello", "hello", "some message")

	assert.NotEqual(t, 12, 67, "message")

	var a *string

	assert.Nil(t, a)
}
