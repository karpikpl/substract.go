package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_SlowSolve(t *testing.T) {
	// Arrange
	reader := strings.NewReader("9223372036854775807 9223372036854775806")
	expected := int64(1)

	// Act
	SlowSolve(reader, func(k *int64) {

		// Assert
		assert.Equal(t, expected, *k)
	})
}

func Benchmark_SlowSolve(b *testing.B) {
	// Arrange
	reader := strings.NewReader("9223372036854775807 9223372036854775806")
	expected := int64(1)

	for i := 0; i < b.N; i++ {
		// Act
		SlowSolve(reader, func(k *int64) {

			// Assert
			assert.Equal(b, expected, *k)
		})
	}
}

func Benchmark_FastSolve(b *testing.B) {
	// Arrange
	reader := strings.NewReader("9223372036854775807 9223372036854775806")
	expected := int64(1)

	for i := 0; i < b.N; i++ {
		// Act
		FastSolve(reader, func(k *int64) {

			// Assert
			assert.Equal(b, expected, *k)
		})
	}
}
