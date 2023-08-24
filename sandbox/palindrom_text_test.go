package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPalindromText(t *testing.T) {
	t.Parallel()
	t.Run("слово не палиндром", func(t *testing.T) {
		str := "слово"

		res := PalindromText(str)

		require.False(t, res)
	})

	t.Run("слово палиндром", func(t *testing.T) {
		str := "АБОБА"

		res := PalindromText(str)

		require.True(t, res)
	})

	t.Run("фраза не палиндром", func(t *testing.T) {
		str := "Аргентина манит негра"

		res := PalindromText(str)

		require.False(t, res)
	})

	t.Run("фраза палиндром", func(t *testing.T) {
		str := "аргентина манит негра"

		res := PalindromText(str)

		require.True(t, res)
	})

}
