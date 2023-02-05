package piscine

import "github.com/01-edu/z01"

func QuadA(x, y int) {
	if x > 0 && y > 0 {
		z01.PrintRune('o')
		for i := 1; i < x; i++ {
			if i < x-1 {
				z01.PrintRune('-')
			} else {
				z01.PrintRune('o')
			}
		}
		z01.PrintRune('\n')

		for j := 1; j < y-1; j++ {
			z01.PrintRune('|')
			for k := 1; k < x; k++ {
				if k < x-1 {
					z01.PrintRune(' ')
				} else {
					z01.PrintRune('|')
				}
			}
			z01.PrintRune('\n')
		}

		if y > 1 {
			z01.PrintRune('o')
			for i := 1; i < x; i++ {
				if i < x-1 {
					z01.PrintRune('-')
				} else {
					z01.PrintRune('o')
				}
			}
			z01.PrintRune('\n')
		}
	}
}
