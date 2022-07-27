package task6kyu

// Sum of Digits / Digital Root
// https://www.codewars.com/kata/541c8630095125aba6000c00

func Sum(n int)int {
	res := 0
	if n==0 {
		return 0
	}
	res = n%10 + Sum(n/10)
	if res >= 10 {
		return Sum(res)
	} 
	return res
}