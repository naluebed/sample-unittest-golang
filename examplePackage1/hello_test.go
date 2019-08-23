package hello

import (
  "testing"
  "github.com/stretchr/testify/assert"
  )

func TestHello(t *testing.T) {
    want := "Hello, world."
    assert.Equal(t, Hello(), want, "test hello world")
}
