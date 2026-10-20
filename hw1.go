package main

import (
	"fmt"
)

func DrawBoard(size int) {
	white := ' '
	black := '#'
	isWhite := false

	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			if isWhite {
				fmt.Printf("%c", white)
				isWhite = false
			} else {
				fmt.Printf("%c", black)
				isWhite = true
			}

		}
		fmt.Println()
	}
}
