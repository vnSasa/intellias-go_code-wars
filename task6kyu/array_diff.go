package task6kyu

// Array.diff
// https://www.codewars.com/kata/523f5d21c841566fde000009

func ArrayDiff(a, b []int) []int {
	diff := make(map[int]struct{}, len(b))
	for _, x := range b {
		diff[x] = struct{}{}
	}
	var res []int
	for _, x := range a {
		if _, found := diff[x]; !found {
			res = append(res, x)
		}
	}
	return res
  }