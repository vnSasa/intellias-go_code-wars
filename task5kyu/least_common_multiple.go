package task5kyu

import "math/big"

func LCM(nums ...int64) *big.Int {
	if len(nums) == 0 {
		return big.NewInt(1)
	}
	lcm := big.NewInt(nums[0])
	for _, n := range nums[1:] {
		if n == 0 {
			return big.NewInt(0)
		}
		gcd := new(big.Int).GCD(nil, nil, lcm, big.NewInt(n))
		lcm.Mul(lcm, big.NewInt(n)).Div(lcm, gcd)
	}
	return lcm
}
