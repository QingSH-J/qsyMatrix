package matrix

import (
	"fmt"
)

type Operator interface {
	Add()
}

// Add performs matrix addition and returns a new matrix.
// The matrices must have the same dimensions.
func Add[T Addable](m *Matrix[T], other *Matrix[T]) (*Matrix[T], error) {
	if m.Rows != other.Rows || m.Columns != other.Columns {
		return nil, fmt.Errorf("matrix dimensions do not match: (%d,%d) vs (%d,%d)",
			m.Rows, m.Columns, other.Rows, other.Columns)
	}

	result := New[T](m.Rows, m.Columns)
	for i := 0; i < m.Rows; i++ {
		for j := 0; j < m.Columns; j++ {
			result.Data[i][j] = m.Data[i][j] + other.Data[i][j]
		}
	}

	return result, nil
}

// Subtract performs matrix subtraction and returns a new matrix.
func Subtract[T Addable](m *Matrix[T], other *Matrix[T]) (*Matrix[T], error) {
	if m.Rows != other.Rows || m.Columns != other.Columns {
		return nil, fmt.Errorf("matrix dimensions do not match: (%d,%d) vs (%d,%d)",
			m.Rows, m.Columns, other.Rows, other.Columns)
	}
	result := New[T](m.Rows, m.Columns)
	for i := 0; i < m.Rows; i++ {
		for j := 0; j < m.Columns; j++ {
			result.Data[i][j] = m.Data[i][j] - other.Data[i][j]
		}
	}
	return result, nil
}

// Multiply performs matrix multiplication and returns a new matrix.
func Multiply[T Multipliable](m *Matrix[T], other *Matrix[T]) (*Matrix[T], error) {
	if m.Columns != other.Rows {
		return nil, fmt.Errorf("matrix dimensions do not match for multiplication: (%d,%d) vs (%d,%d)",
			m.Rows, m.Columns, other.Rows, other.Columns)
	}

	result := New[T](m.Rows, other.Columns)
	for i := 0; i < m.Rows; i++ {
		for j := 0; j < other.Columns; j++ {
			var sum T
			for k := 0; k < m.Columns; k++ {
				sum += m.Data[i][k] * other.Data[k][j]
			}
			result.Data[i][j] = sum
		}
	}

	return result, nil
}
