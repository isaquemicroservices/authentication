package utils_test

import (
	"testing"
	"time"

	"github.com/isaqueveras/authentication-microservice/utils"
	"github.com/stretchr/testify/assert"
)

func TestPointer(t *testing.T) {
	t.Run("TestGetPointerInt64", func(t *testing.T) {
		var value int64 = 1
		assert.Equal(t, value, *utils.GetPointerInt64(value))
	})

	t.Run("TestGetPointerString", func(t *testing.T) {
		var value string = "@isaqueveras"
		assert.Equal(t, value, *utils.GetPointerString(value))
	})

	t.Run("TestGetPointerFloat64", func(t *testing.T) {
		var value float64 = 24.00
		assert.Equal(t, value, *utils.GetPointerFloat64(value))
	})

	t.Run("TestGetPointerTime", func(t *testing.T) {
		var value time.Time = time.Now()
		assert.Equal(t, value, *utils.GetPointerTime(value))
	})

	t.Run("TestGetPointerBool", func(t *testing.T) {
		var value bool = true
		assert.Equal(t, value, *utils.GetPointerBool(value))
	})
}
