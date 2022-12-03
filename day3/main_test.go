package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_RuneToPriority(t *testing.T) {
	assert.Equal(t, 1, RuneToPriority('a'))
	assert.Equal(t, 26, RuneToPriority('z'))
	assert.Equal(t, 27, RuneToPriority('A'))
	assert.Equal(t, 52, RuneToPriority('Z'))
}
