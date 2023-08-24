package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPalindromNumber(t *testing.T) {
	t.Parallel()

	t.Run("число палиндром", func(t *testing.T) {
		num := 0

		res := PalindromNumber(num)

		require.True(t, res)
	})

	t.Run("число палиндром", func(t *testing.T) {
		num := 1

		res := PalindromNumber(num)

		require.True(t, res)
	})

	t.Run("число палиндром", func(t *testing.T) {
		num := 111

		res := PalindromNumber(num)

		require.True(t, res)
	})

	t.Run("число палиндром", func(t *testing.T) {
		num := 1111

		res := PalindromNumber(num)

		require.True(t, res)
	})

	t.Run("число не палиндром", func(t *testing.T) {
		num := 10

		res := PalindromNumber(num)

		require.False(t, res)
	})
}
