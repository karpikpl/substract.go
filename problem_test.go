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

func Test_FastSolve(t *testing.T) {
	// Arrange
	reader := strings.NewReader("9223372036854775807 9223372036854775806\n5 12\n10 12\n1871293781758123 72784\n1 9223372036854775807")
	expected := []int64{1, 7, 2, 1871293781685339, 9223372036854775806}
	result := make([]int64, 0)

	// Act
	SlowSolve(reader, func(k *int64) {
		result = append(result, *k)
	})
	// Assert
	assert.Equal(t, expected, result)
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
