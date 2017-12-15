package euler

import "math/big"

// BigNumberLength 桁数を計算
func BigNumberLength(n *big.Int) int {
	return len(new(big.Int).Abs(n).String())
}
