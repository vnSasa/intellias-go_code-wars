package task4kyu

import "math"

func SumOfSquares(n uint64) uint64 {
    if isSquare(n) {
        return 1
    }
    for i := uint64(1); i*i <= n; i++ {
        if isSquare(n-i*i) {
            return 2
        }
    }
    for i := uint64(1); i*i <= n; i++ {
        for j := uint64(i); i*i+j*j <= n; j++ {
            if isSquare(n-i*i-j*j) {
                return 3
            }
        }
    }
    return 4
}

func isSquare(n uint64) bool {
    root := uint64(math.Sqrt(float64(n)))
    return root*root == n
}