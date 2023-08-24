package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFibonacciNumber(t *testing.T) {
	t.Parallel()
	t.Run("число Фибоначчи 1", func(t *testing.T) {
		num := 1
		expected := 0

		res := FibonacciNumber(num)

		require.Equal(t, expected, res)
	})

	t.Run("число Фибоначчи 2", func(t *testing.T) {
		num := 2
		expected := 1

		res := FibonacciNumber(num)

		require.Equal(t, expected, res)
	})

	t.Run("число Фибоначчи 3", func(t *testing.T) {
		num := 3
		expected := 1

		res := FibonacciNumber(num)

		require.Equal(t, expected, res)
	})

	t.Run("число Фибоначчи 4", func(t *testing.T) {
		num := 4
		expected := 2

		res := FibonacciNumber(num)

		require.Equal(t, expected, res)
	})

	t.Run("число Фибоначчи 5", func(t *testing.T) {
		num := 5
		expected := 3

		res := FibonacciNumber(num)

		require.Equal(t, expected, res)
	})

	t.Run("Фибоначчи 6", func(t *testing.T) {
		num := 6
		expected := 8

		res := FibonacciNumber(num)

		require.Equal(t, expected, res)
	})

	t.Run("число Фибоначчи 7", func(t *testing.T) {
		num := 7
		expected := 13

		res := FibonacciNumber(num)

		require.Equal(t, expected, res)
	})

	t.Run("число Фибоначчи 8", func(t *testing.T) {
		num := 8
		expected := 21

		res := FibonacciNumber(num)

		require.Equal(t, expected, res)
	})

	t.Run("число Фибоначчи 9", func(t *testing.T) {
		num := 9
		expected := 34

		res := FibonacciNumber(num)

		require.Equal(t, expected, res)
	})

	t.Run("число Фибоначчи 10", func(t *testing.T) {
		num := 10
		expected := 55

		res := FibonacciNumber(num)

		require.Equal(t, expected, res)
	})

}
