package task4kyu

func ExpSum(n uint64) uint64 {
    table := make([]uint64, n+1)
    table[0] = 1
    
    for i := uint64(1); i <= n; i++ {
        for j := i; j <= n; j++ {
            table[j] += table[j-i]
        }
    }

    return table[n]
}