package piscine

import "github.com/01-edu/z01"

func QuadC(x, y int) {
	if x > 0 && y > 0 {
		z01.PrintRune('A')
		for i := 1; i < x; i++ {
			if i < x-1 {
				z01.PrintRune('B')
			} else {
				z01.PrintRune('A')
			}
		}
		z01.PrintRune('\n')
		for j := 1; j < y; j++ {
			if j < y-1 {
				z01.PrintRune('B')
				for k := 1; k < x; k++ {
					if k < x-1 {
						z01.PrintRune(' ')
					} else {
						z01.PrintRune('B')
					}
				}
				z01.PrintRune('\n')
			} else {
				z01.PrintRune('C')
				for i := 1; i < x; i++ {
					if i < x-1 {
						z01.PrintRune('B')
					} else {
						z01.PrintRune('C')
					}
				}
			}
		}
		if y > 1 {
			z01.PrintRune('\n')
		}
	}
}
