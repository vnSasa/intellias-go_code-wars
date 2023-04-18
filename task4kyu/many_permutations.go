package task4kyu

import (
    "sort"
)

func Permutations(s string) []string {
    if len(s) <= 1 {
        return []string{s}
    }

    result := make(map[string]bool)

    for i, c := range s {
        sub := s[:i] + s[i+1:]
        for _, perm := range Permutations(sub) {
            str := string(c) + perm
            result[str] = true
        }
    }

    keys := make([]string, 0, len(result))
    for k := range result {
        keys = append(keys, k)
    }
    sort.Strings(keys)

    return keys
}