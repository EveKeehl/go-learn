package main

import (
	"testing"
)

func TestItoa(t *testing.T) {
	type pair struct {
		i        int
		expected string
	}

	test := []pair{
		{0, "0"},
		{22, "22"},
		{32432523, "32432523"},
		{-3, "-3"},
	}
	for _, tst := range test {
		if tst.expected != itoa(tst.i) {
			t.Errorf("%d - %s\n", tst.i, "FAIL")
		}
	}
}
