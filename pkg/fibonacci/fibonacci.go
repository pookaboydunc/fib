package fibonacci

import (
	"errors"
	"math/big"
	"strconv"
)

const (
	MAX_32_BIT_N = 46
	MAX_64_BIT_N = 92
)

var (
	ErrNonNegativeNumber = errors.New("n must be non-negative")
)

// SafeLimit returns the largest value that can be passed to NthNumSequence or NthBigNumSequence
func SafeLimit() int {
	if strconv.IntSize == 32 {
		return MAX_32_BIT_N
	}
	return MAX_64_BIT_N
}

// Sequence returns the first n fibonacci numbers as a slice.
// The slice is returned as an interface to allow for both []int and []*big.Int
func Sequence(n int) (any, error) {
	if n < 0 {
		return []int{}, ErrNonNegativeNumber
	}
	if n > SafeLimit() {
		return nthBigNumSequence(n), nil
	}
	return nthNumSequence(n), nil
}

// nthNumSequence returns the first n fibonacci numbers as a slice.
// Example: nthNumSequence(5) -> [0, 1, 1, 2, 3]
func nthNumSequence(n int) []int {
	if n == 0 {
		return []int{}
	}
	if n == 1 {
		return []int{0}
	}

	seq := make([]int, n)
	seq[0] = 0
	seq[1] = 1

	for i := 2; i < n; i++ {
		seq[i] = seq[i-1] + seq[i-2]
	}

	return seq
}

// nthBigNumSequence returns the first n fibonacci numbers as a slice.
// The numbers are represented as big.Int so should be used when you need your sequence is expected to contain large numbers.
// Example: nthBigNumSequence(5) -> [0, 1, 1, 2, 3]
func nthBigNumSequence(n int) []*big.Int {
	if n <= 0 {
		return []*big.Int{}
	}
	if n == 1 {
		return []*big.Int{big.NewInt(0)}
	}

	seq := make([]*big.Int, n)
	seq[0] = big.NewInt(0)
	seq[1] = big.NewInt(1)

	for i := 2; i < n; i++ {
		seq[i] = new(big.Int).Add(seq[i-1], seq[i-2])
	}

	return seq
}
