package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPalindromText(t *testing.T) {
	t.Parallel()
	t.Run("не палиндром", func(t *testing.T) {
		str := "слово"

		res := PalindromText(str)

		require.False(t, res)
	})

	t.Run("палиндром", func(t *testing.T) {
		str := "АБОБА"

		res := PalindromText(str)

		require.True(t, res)
	})

	t.Run("не палиндром", func(t *testing.T) {
		str := "Аргентина манит негра"

		res := PalindromText(str)

		require.False(t, res)
	})

	t.Run("палиндром", func(t *testing.T) {
		str := "аргентина манит негра"

		res := PalindromText(str)

		require.True(t, res)
	})

}
