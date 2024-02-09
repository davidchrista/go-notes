package main

import (
  "testing"
  "github.com/stretchr/testify/assert"
)

func TestFooer2(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(Fooer(3), "Foo")
	assert.Equal(Fooer(4), "4")
}
