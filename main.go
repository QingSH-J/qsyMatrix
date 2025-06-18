package main

import (
	"github.com/QingSH-J/qsyMatrix/matrix"
	"fmt"
)

func main() {
	m := matrix.Zero[float32](3, 3)

	for i := 0; i < 3; i ++{
		for j := 0; j < 3; j++ {
			m.Update(i, j, 1.0)
		}
	}

	rank := m.Rank()
	m.PrintMatrix()
	fmt.Println("Rank of the matrix:", rank)

}
