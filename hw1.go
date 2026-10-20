package main

import (
	"fmt"
	"go/types"
)

type Board struct {
	rows []Row
	row  Row
}

type Row struct {
	squares []rune
}

func (b *Board) fill(size int) {
	var white rune = ' '
	var black rune = '#'
	var isWhite bool = false

	for i := 0; i < size; i++ {
		b.rows = make([]Row, size)
		for j := 0; j < size; j++ {
			if isWhite {
				b.rows[i] = append(b.row.squares, white)
				isWhite = false
			} else {
				b.row.squares = append(b.row.squares, black)
				isWhite = true
			}
		}
		b.rows = append(b.rows, types.Slice{})
		for _, t := range b.squares {
			fmt.Print(string(t))
		}
	}
}

func DrawBoard() {

}
