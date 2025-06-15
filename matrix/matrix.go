package matrix

import (
	"fmt"
)

type Number interface {
	int | int8 | int16 | int32 | int64 |
		uint | uint8 | uint16 | uint32 | uint64 |
		float32 | float64 |
		complex64 | complex128 |
		~string | ~[]byte
}

// Addable defines types that support the + operator
type Addable interface {
	int | int8 | int16 | int32 | int64 |
		uint | uint8 | uint16 | uint32 | uint64 |
		float32 | float64 |
		complex64 | complex128 |
		~string
}

type Matrix[T Number] struct {
	Rows    int
	Columns int
	Data    [][]T
}

// New creates a new matrix with the specified number of rows and columns.
func New[T Number](rows, columns int) *Matrix[T] {
	data := make([][]T, rows)
	for i := range data {
		data[i] = make([]T, columns)
	}
	return &Matrix[T]{Rows: rows, Columns: columns, Data: data}
}

func NewFromOther[T Number](other *Matrix[T]) *Matrix[T] {
	data := make([][]T, other.Rows)
	for i := range data {
		data[i] = make([]T, other.Columns)
		copy(data[i], other.Data[i])
	}
	return &Matrix[T]{Rows: other.Rows, Columns: other.Columns, Data: data}
}

// GetRow returns the row at the specified index.
func (m *Matrix[T]) GetRow() int {
	if m.Rows == 0 {
		return 0
	}
	return m.Rows
}

// GetColumn returns the number of columns in the matrix.
func (m *Matrix[T]) GetColumn() int {
	if m.Columns == 0 {
		return 0
	}
	return m.Columns
}

// Zero creates a new matrix filled with zero values of type T.
func Zero[T Number](rows, columns int) *Matrix[T] {
	return New[T](rows, columns)
}

func Eye[T Number](n int) *Matrix[int] {
	m := New[int](n, n)
	for i := 0; i < n; i++ {
		m.Data[i][i] = int(1) // Assuming T can be converted to 1
	}
	return m
}

func (m *Matrix[T]) PrintMatrix() {
	for i := 0; i < m.Rows; i++ {
		fmt.Print("[")
		for j := 0; j < m.Columns; j++ {
			if j > 0 {
				fmt.Print(", ")
			}
			fmt.Printf("%v", m.Data[i][j])
		}
		fmt.Println("]")
	}
}
