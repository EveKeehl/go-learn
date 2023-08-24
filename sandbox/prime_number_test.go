package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPrimeNumber(t *testing.T) {
	t.Parallel()
	t.Run("простое число 1", func(t *testing.T) {
		num := 1
		expected := 2

		res := PrimeNumber(num)

		require.Equal(t, expected, res)
	})

	t.Run("простое число 2", func(t *testing.T) {
		num := 2
		expected := 3

		res := PrimeNumber(num)

		require.Equal(t, expected, res)
	})

	t.Run("простое число 3", func(t *testing.T) {
		num := 3
		expected := 5

		res := PrimeNumber(num)

		require.Equal(t, expected, res)
	})

	t.Run("простое число 4", func(t *testing.T) {
		num := 4
		expected := 7

		res := PrimeNumber(num)

		require.Equal(t, expected, res)
	})

	t.Run("простое число 5", func(t *testing.T) {
		num := 5
		expected := 11

		res := PrimeNumber(num)

		require.Equal(t, expected, res)
	})

	t.Run("простое число 6", func(t *testing.T) {
		num := 6
		expected := 13

		res := PrimeNumber(num)

		require.Equal(t, expected, res)
	})

	t.Run("простое число 7", func(t *testing.T) {
		num := 7
		expected := 17

		res := PrimeNumber(num)

		require.Equal(t, expected, res)
	})

	t.Run("простое число 8", func(t *testing.T) {
		num := 8
		expected := 19

		res := PrimeNumber(num)

		require.Equal(t, expected, res)
	})

	t.Run("простое число 9", func(t *testing.T) {
		num := 9
		expected := 23

		res := PrimeNumber(num)

		require.Equal(t, expected, res)
	})

	t.Run("простое число 10", func(t *testing.T) {
		num := 10
		expected := 29

		res := PrimeNumber(num)

		require.Equal(t, expected, res)
	})

}
