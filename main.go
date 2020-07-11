package main

import (
	"fmt"
	"strconv"
	"strings"
)

func contains(slice []int, item int) bool {
	for _, a := range slice {
		if a == item {
			return true
		}
	}
	return false
}

func validate(unit [9]int) bool {
	var nine []int
	for i := range unit {
		value := unit[i]
		if contains(nine, value) {
			return false
		}
		if value != 0 {
			nine = append(nine, unit[i])
		}
	}
	return true
}

// Board is a sudoku board
type Board struct {
	Cells [81]int
}

func (b Board) getColumn(i int) [9]int {
	var column [9]int
	for j := 0; j < 9; j++ {
		item := i + 9*j
		column[j] = b.Cells[item]
	}
	return column
}

func (b Board) getRow(i int) [9]int {
	var column [9]int
	for j := 0; j < 9; j++ {
		item := j + 9*i
		column[j] = b.Cells[item]
	}
	return column
}

func (b Board) getBlock(i int) [9]int {
	var block [9]int
	start := (27 * (i / 3)) + (i%3)*3
	for j := 0; j < 9; j++ {
		item := start + 9*(j/3) + (j % 3)
		block[j] = b.Cells[item]
	}
	return block
}

func GetBlockByNr(i int) int {
	return (i/27)*3 + ((i - ((i / 27) * 9)) % 3)
}

func GetRowByNr(i int) int {
	return i / 9
}

func GetColumnByNr(i int) int {
	return i - (i/9)*9
}

func (b Board) validateCell(i int) bool {
	bnr := GetBlockByNr(i)
	cnr := GetColumnByNr(i)
	rnr := GetRowByNr(i)
	block := b.getBlock(bnr)
	column := b.getColumn(cnr)
	row := b.getRow(rnr)

	if !(validate(block) && validate(column) && validate(row)) {
		return false
	}
	return true
}

func (b Board) validate() bool {
	for j := 0; j < 9; j++ {
		block := b.getBlock(j)
		column := b.getColumn(j)
		row := b.getRow(j)

		if !(validate(block) && validate(column) && validate(row)) {
			return false
		}
	}
	return true
}

func (b *Board) solve() (Board, bool) {
	var cursorHistory []int
	var empties []int

	for i := range b.Cells {
		if b.Cells[i] == 0 {
			empties = append(empties, i)
		}
	}

	for i := 0; i < len(empties); i++ {
		cursor := empties[i]
		cursorValue := b.Cells[cursor]
		if cursorValue == 9 {
			b.Cells[cursor] = 0
			i = i - 2
			continue
		}
		b.Cells[cursor] = b.Cells[cursor] + 1

		if b.validateCell(cursor) {
			cursorHistory = append(cursorHistory, cursor)
		} else {
			i = i - 1
		}
	}
	return *b, true
}

func boardFromString(s string) Board {
	var cells [81]int
	split := strings.Split(s, "")
	for i := range cells {
		x, _ := strconv.Atoi(split[i])
		cells[i] = x
	}
	return Board{Cells: cells}
}

func main() {
	// definition := "530070000600195000098000060800060003400803001700020006060000280000419005000080079"
	definition := "800000000003600000070090200050007000000045700000100030001000068008500010090000400"
	board := boardFromString(definition)
	board.solve()
	fmt.Println(board.Cells)

}
