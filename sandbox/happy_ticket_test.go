package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestHappyTicket(t *testing.T) {
	t.Parallel()

	t.Run("билет счастливый 1", func(t *testing.T) {
		num := 1111

		res := HappyTicket(num)

		require.True(t, res)
	})

	t.Run("билет счастливый 2", func(t *testing.T) {
		num := 111003

		res := HappyTicket(num)

		require.True(t, res)
	})

	t.Run("билет счастливый 3", func(t *testing.T) {
		num := 1119003

		res := HappyTicket(num)

		require.True(t, res)
	})

	t.Run("билет не счастливый 1", func(t *testing.T) {
		num := 123456

		res := HappyTicket(num)

		require.False(t, res)
	})

}
