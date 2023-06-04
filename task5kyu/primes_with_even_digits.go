package task5kyu

// Primes with Even Digits
// https://www.codewars.com/kata/582dcda401f9ccb4f0000025

import (
	"math"
	"strconv"
)

func F(n int) int {
	var highestPrime int

	count := 0
	prev := 0

	var nn []int

	// create a slice for elements from 1 to n
	for i := 1; i < n; i++ {
		// instead of the number 1,
		// or if the number is divisible by 2 without a remainder,
		// we set the number 0
		if i == 1 || i%2 == 0 {
			nn = append(nn, 0)
		}

		// other numbers are recorded in a slice
		if i%2 != 0 {
			nn = append(nn, i)
		}
	}

	x := math.Sqrt(float64(n))

	// instead of the rest of the non-prime numbers, we write 0
	for i := 3; i < int(x); i++ {
		for j := i * i; j < n; j += i {
			nn[j] = 0
		}
	}

	var nums []int

	// create a new slice and write only numbers that are not equal to 0 to it
	for i := 0; i < len(nn); i++ {
		if nn[i] != 0 {
			nums = append(nums, nn[i])
		}
	}

	// We find the closest prime number to a certain integer n,
	// which has the maximum possible number of even digits.
	for i := len(nums) - 1; i > 0; i-- {
		count = 0
		num := nums[i]
		for j := 0; j < len(strconv.Itoa(nums[i])); j++ {
			num /= 10
			if num%2 == 0 {
				count++
			}
		}

		if count > prev {
			prev = count
			highestPrime = nums[i]
		}
	}
	return highestPrime
}
