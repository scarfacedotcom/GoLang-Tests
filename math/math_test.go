package math_test

import (
	"scarfacetest/math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	expected := 5
	actual := math.Add(2, 3)
	assert.Equal(t, expected, actual, "Expected and actual values should be equal")
}
