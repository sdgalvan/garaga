package fptower

import (
	// "errors"
	// "github.com/consensys/gnark-crypto/ecc"
	// "github.com/keep-starknet-strange/garaga/fp"
	// "github.com/consensys/gnark-crypto/ecc/bn254/fr"
	// "math/big"
	// "sync"

	// "github.com/keep-starknet-strange/garaga/py/sage/common/arithmetic/curves"
)

type GTPoint struct {
	C0, C1 E12
}

func (z *E12) Pair(x, y *GTPoint) *E12 {
	return &x.C0;
	// return long_pairing(x, y); call python method
}