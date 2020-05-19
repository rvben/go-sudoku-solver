package main

import (
	"fmt"
	"strings"
	"testing"
)

type Unit [9]int

func unitToString(u Unit) string {
	return fmt.Sprintf("Unit{%s}", strings.Trim(strings.Join(strings.Fields(fmt.Sprint(u)), ","), "[]"))
}

// definition := "530070000600195000098000060800060003400803001700020006060000280000419005000080079"
func TestBoard(t *testing.T) {
	definition := "534678912672195348198342567859761423426853791713924856961537284287419635345286179"
	got := boardFromString(definition)

	if !got.validate() {
		t.Errorf("valid sudoku is regarded unvalid")
	}
}

func TestGetRow(t *testing.T) {
	// Create board with position numbers
	b := Board{}
	for i := range b.Cells {
		b.Cells[i] = i
	}

	tests := []struct {
		row  int
		want Unit
	}{
		{0, Unit{0, 1, 2, 3, 4, 5, 6, 7, 8}},
		{1, Unit{9, 10, 11, 12, 13, 14, 15, 16, 17}},
		{2, Unit{18, 19, 20, 21, 22, 23, 24, 25, 26}},
		{3, Unit{27, 28, 29, 30, 31, 32, 33, 34, 35}},
		{4, Unit{36, 37, 38, 39, 40, 41, 42, 43, 44}},
		{5, Unit{45, 46, 47, 48, 49, 50, 51, 52, 53}},
		{6, Unit{54, 55, 56, 57, 58, 59, 60, 61, 62}},
		{7, Unit{63, 64, 65, 66, 67, 68, 69, 70, 71}},
		{8, Unit{72, 73, 74, 75, 76, 77, 78, 79, 80}},
	}

	for _, tt := range tests {
		got := b.getRow(tt.row)
		if got != tt.want {
			t.Errorf("got %s want %s", unitToString(got), unitToString(tt.want))
		}
	}
}

func TestGetColumn(t *testing.T) {
	// Create board with position numbers
	b := Board{}
	for i := range b.Cells {
		b.Cells[i] = i
	}

	tests := []struct {
		column int
		want   Unit
	}{
		{0, Unit{0, 9, 18, 27, 36, 45, 54, 63, 72}},
		{1, Unit{1, 10, 19, 28, 37, 46, 55, 64, 73}},
		{2, Unit{2, 11, 20, 29, 38, 47, 56, 65, 74}},
		{3, Unit{3, 12, 21, 30, 39, 48, 57, 66, 75}},
		{4, Unit{4, 13, 22, 31, 40, 49, 58, 67, 76}},
		{5, Unit{5, 14, 23, 32, 41, 50, 59, 68, 77}},
		{6, Unit{6, 15, 24, 33, 42, 51, 60, 69, 78}},
		{7, Unit{7, 16, 25, 34, 43, 52, 61, 70, 79}},
		{8, Unit{8, 17, 26, 35, 44, 53, 62, 71, 80}},
	}

	for _, tt := range tests {
		got := b.getColumn(tt.column)
		if got != tt.want {
			t.Errorf("got %s want %s", unitToString(got), unitToString(tt.want))
		}
	}
}

func TestGetBlock(t *testing.T) {
	// Create board with position numbers
	b := Board{}
	for i := range b.Cells {
		b.Cells[i] = i
	}

	tests := []struct {
		block int
		want  Unit
	}{
		{0, Unit{0, 1, 2, 9, 10, 11, 18, 19, 20}},
		{1, Unit{3, 4, 5, 12, 13, 14, 21, 22, 23}},
		{2, Unit{6, 7, 8, 15, 16, 17, 24, 25, 26}},
		{3, Unit{27, 28, 29, 36, 37, 38, 45, 46, 47}},
		{4, Unit{30, 31, 32, 39, 40, 41, 48, 49, 50}},
		{5, Unit{33, 34, 35, 42, 43, 44, 51, 52, 53}},
		{6, Unit{54, 55, 56, 63, 64, 65, 72, 73, 74}},
		{7, Unit{57, 58, 59, 66, 67, 68, 75, 76, 77}},
		{8, Unit{60, 61, 62, 69, 70, 71, 78, 79, 80}},
	}

	for _, tt := range tests {
		got := b.getBlock(tt.block)
		if got != tt.want {
			t.Errorf("got %s want %s", unitToString(got), unitToString(tt.want))
		}
	}
}

func TestValidate(t *testing.T) {
	tests := []struct {
		unit Unit
		want bool
	}{
		{Unit{1, 2, 3, 4, 5, 6, 7, 8, 9}, true},
		{Unit{1, 1, 2, 3, 4, 5, 6, 7, 8}, false},
		{Unit{1, 2, 3, 4, 5, 6, 0, 0, 0}, true},
		{Unit{1, 2, 3, 4, 5, 6, 0, 0, 2}, false},
	}

	for _, tt := range tests {
		got := validate(tt.unit)
		if got != tt.want {
			t.Errorf("Validation failed.")
		}
	}
}

func TestBoardSolveSolved(t *testing.T) {
	definition := "534678912672195348198342567859761423426853791713924856961537284287419635345286179"
	board := boardFromString(definition)

	if _, ok := board.solve(); !ok {
		t.Errorf("Solving failed.")
	}
}

func TestBoardSolveSimple(t *testing.T) {
	definition := "530070000600195000098000060800060003400803001700020006060000280000419005000080079"
	board := boardFromString(definition)

	if _, ok := board.solve(); !ok {
		t.Errorf("Solving failed.")
	}
}

func TestBoardSolveHardest(t *testing.T) {
	definition := "800000000003600000070090200050007000000045700000100030001000068008500010090000400"
	board := boardFromString(definition)

	if _, ok := board.solve(); !ok {
		t.Errorf("Solving failed.")
	}
}
