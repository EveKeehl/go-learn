package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFibonacciNumber(t *testing.T) {
	t.Parallel()

	t.Run("число Фибоначчи 10", func(t *testing.T) {
		num := 10
		expected := 55

		res := FibonacciNumber(num)

		require.Equal(t, expected, res)
	})

}
