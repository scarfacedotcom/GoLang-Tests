package demo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUpdateArray(t *testing.T) {
	testArray := [2]string{"Value1", "Value2"}
	UpdateArray1(testArray)
	assert.Equal(t, NewValue, testArray[0])
}
