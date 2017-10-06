package simple

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSomething(t *testing.T) {

	assert := assert.New(t)
	assert.Equal("hello", "hello", "some message")

	assert.NotEqual(12, 67, "message")

	var a *string

	assert.Nil(a)
}
