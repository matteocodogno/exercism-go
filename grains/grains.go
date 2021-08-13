package grains

import (
	"errors"
	"math"
	"math/big"
)

func Square(cellNumber int) (uint64, error) {
	if cellNumber < 1 || cellNumber > 64 {
		return 0, errors.New("")
	}

	return uint64(math.Exp2(float64(cellNumber - 1))), nil
}

func Total() uint64 {
	tot, _ := big.NewFloat(math.Exp2(64) - 1).Uint64()

	return tot
}