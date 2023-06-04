package task6kyu

// The Pyramid of Cubes
// https://www.codewars.com/kata/61707b71059070003793bc0f

func FindHeight(cubes int) int {
	if cubes == 0 {
		return 0
	}
	layer := 2
	used := 1
	n := 1
	for i := 0; ; i++ {
		used += n + layer
		if used == cubes {
			return layer
		}
		if used > cubes {
			return layer - 1
		}
		n += layer
		layer++
	}
}
