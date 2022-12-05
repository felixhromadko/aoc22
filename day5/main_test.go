package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Staple(t *testing.T) {
	s := &Staple{}

	s.Push(1)
	s.Push(2)
	s.Push(3)

	assert.Equal(t, uint8(3), s.Pop())
	assert.Equal(t, uint8(2), s.Pop())
	assert.Equal(t, uint8(1), s.Pop())
}
