package main

import (
	"fmt"
	"testing"
)

func TestSum(t *testing.T) {
	tables := []struct {
		x int
		y int
		n int
	}{
		{10, 1, 11},
		{1, 2, 3},
		{2, 2, 4},
		{5, 2, 7},
	}

	for _, table := range tables {
		total := Sum(table.x, table.y)
		fmt.Printf("Sum of (%d+%d) was correct, got: %d, want: %d.", table.x, table.y, total, table.n)
		if total != table.n {
			t.Errorf("Sum of (%d+%d) was incorrect, got: %d, want: %d.", table.x, table.y, total, table.n)
		}
	}
}
