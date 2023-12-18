package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOr(t *testing.T) {
	testI := make(chan interface{})
	close(testI)
	assert.Empty(t, or(testI))
}
