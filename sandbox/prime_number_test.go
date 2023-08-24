package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPrimeNumber(t *testing.T) {
	t.Parallel()

	t.Run("простое число 10", func(t *testing.T) {
		num := 10
		expected := 29

		res := PrimeNumber(num)

		require.Equal(t, expected, res)
	})

}
