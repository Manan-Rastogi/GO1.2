package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSumTillN(t *testing.T) {
	var n = 100
	sum_n := SumTillN(n)

	sum_t := 0
	for i := 1; i <= 100; i++ {
		sum_t += i
	}

	require.NotZero(t, sum_n)
	require.Equal(t, sum_t, sum_n)
}
