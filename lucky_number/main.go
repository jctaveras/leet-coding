package main

import (
	"fmt"
	"sort"
)

func main() {
	matrix := [][]int{
		{36376,85652,21002,4510},
		{68246,64237,42962,9974},
		{32768,97721,47338,5841},
		{55103,18179,79062,46542},
	}

	fmt.Println(luckyNumber(matrix))
}

func luckyNumber(matrix [][]int) []int {
	lucky, columns, rows := 0, make([][]int, len(matrix[0])), [][]int{}

	for _, row := range matrix {
		currentRow := make([]int, len(matrix[0]))
		
		copy(currentRow, row)
		sort.Ints(currentRow)
		rows = append(rows, currentRow)

		for index, value := range row {
			columns[index] = append(columns[index], value)
		}
	}

	for _, column := range columns {
		sort.SliceStable(column, func(i, j int) bool { return column[i] > column[j] })	
	}

	for _, row := range rows {
		for _, column := range columns {
			if column[0] == row [0] {
				lucky = column[0]
			}
		}
	}
	
	if lucky == 0 { return []int{} }

	return []int{lucky}
}
