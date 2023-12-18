package main

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGetTime(t *testing.T) {
	time_g, err := GetTime()
	t_now := time.Now()
	assert.Equal(t, t_now.Round(time.Minute), time_g.Round(time.Minute))
	assert.Equal(t, err, nil)
}
