package matrix

import (
	"fmt"
)

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
