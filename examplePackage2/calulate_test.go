package caculate

import (
  "testing"
  "github.com/stretchr/testify/assert"
  )

func TestCalculate(t *testing.T) {
    assert.True(t, isEven(2), "test1")
    assert.False(t, isEven(3), "test2")
    assert.False(t, isEven(5), "test3")
}
