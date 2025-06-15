package main

import (
	"github.com/QingSH-J/qsyMatrix/matrix"
)

func main() {
	m := matrix.Zero[int](3, 3)

	m1 := matrix.Eye[int](3)

	newM, err := matrix.Add(m, m1)

	if err != nil {
		panic(err)
	}
	newM.PrintMatrix()

}
