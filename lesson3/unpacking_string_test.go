package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestExtractString(t *testing.T) {
	t.Parallel()
	t.Run("correct string", func(t *testing.T) {
		str := "a4bc2d5e"
		expected := "aaaabccddddde"

		res, err := UnpackingString(str)

		require.NoError(t, err)
		require.Equal(t, expected, res)
	})

	t.Run("correct string2", func(t *testing.T) {
		str := "abcd"
		expected := "abcd"

		res, err := UnpackingString(str)

		require.NoError(t, err)
		require.Equal(t, expected, res)
	})

	t.Run("escaped string", func(t *testing.T) {
		str := `qwe\4\5`
		expected := "qwe45"

		res, err := UnpackingString(str)

		require.NoError(t, err)
		require.Equal(t, expected, res)
	})

	t.Run("escaped string2", func(t *testing.T) {
		str := `qwe\45`
		expected := "qwe44444"

		res, err := UnpackingString(str)

		require.NoError(t, err)
		require.Equal(t, expected, res)
	})

	t.Run("escaped string3", func(t *testing.T) {
		str := `qwe\\5`
		expected := `qwe\\\\\`

		res, err := UnpackingString(str)

		require.NoError(t, err)
		require.Equal(t, expected, res)
	})

	t.Run("incorrect string", func(t *testing.T) {
		str := "45"
		expected := ""

		res, err := UnpackingString(str)

		require.ErrorIs(t, err, ErrIncorrectString)
		require.Equal(t, expected, res)
	})

}
