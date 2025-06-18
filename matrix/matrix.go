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

// Addable defines types that support the + operator - operator
type Addable interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 |
		~float32 | ~float64 |
		~complex64 | ~complex128 |
		~string
	int | int8 | int16 | int32 | int64 |
		uint | uint8 | uint16 | uint32 | uint64 |
		float32 | float64 |
		complex64 | complex128
}

// Multipliable defines types that support the * operator
type Multipliable interface {
	int | int8 | int16 | int32 | int64 |
		uint | uint8 | uint16 | uint32 | uint64 |
		float32 | float64 |
		complex64 | complex128
}

// Dividable defines types that support the / operator
type Numeric interface {
	Number | Addable
}

type Matrix[T Numeric] struct {
	Rows    int
	Columns int
	Data    [][]T
}

// New creates a new matrix with the specified number of rows and columns.
func New[T Numeric](rows, columns int) *Matrix[T] {
	data := make([][]T, rows)
	for i := range data {
		data[i] = make([]T, columns)
	}
	return &Matrix[T]{Rows: rows, Columns: columns, Data: data}
}

func NewFromOther[T Numeric](other *Matrix[T]) *Matrix[T] {
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
func Zero[T Numeric](rows, columns int) *Matrix[T] {
	return New[T](rows, columns)
}

func Eye[T Number](n int) *Matrix[T] {
	m := New[T](n, n)
	for i := 0; i < n; i++ {
		// 使用反射或类型断言来设置单位值
		var one any
		switch any(m.Data[i][i]).(type) {
		case int:
			one = 1
		case int8:
			one = int8(1)
		case int16:
			one = int16(1)
		case int32:
			one = int32(1)
		case int64:
			one = int64(1)
		case uint:
			one = uint(1)
		case uint8:
			one = uint8(1)
		case uint16:
			one = uint16(1)
		case uint32:
			one = uint32(1)
		case uint64:
			one = uint64(1)
		case float32:
			one = float32(1.0)
		case float64:
			one = float64(1.0)
		case complex64:
			one = complex64(1 + 0i)
		case complex128:
			one = complex128(1 + 0i)
		case string:
			one = "1"
		case []byte:
			one = []byte("1")
		}
		m.Data[i][i] = one.(T)
	}
	return m
}

// PrintMatrix prints the matrix in a readable format.
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

// Update updates the matrix with the values from user
func (m *Matrix[T]) Update(row, column int, value T) error {
	if row < 0 || row >= m.Rows || column < 0 || column >= m.Columns {
		return fmt.Errorf("index out of bounds")
	}
	m.Data[row][column] = value
	return nil
}

// Transpose returns the transpose of the matrix.
func (m *Matrix[T]) Transpose() *Matrix[T] {
	transposed := New[T](m.Columns, m.Rows)
	for i := 0; i < m.Rows; i++ {
		for j := 0; j < m.Columns; j++ {
			transposed.Data[j][i] = m.Data[i][j]
		}
	}
	return transposed
}

// Rank returns the rank of the matrix.
func (m *Matrix[T]) Rank() int {
	var zero T
	switch any(zero).(type) {
	case int, int8, int16, int32, int64,
		uint, uint8, uint16, uint32, uint64,
		float32, float64,
		complex64, complex128:
		return m.RankNumeric() // For numeric types, we can implement a rank calculation
	default:
		return 0
	}
}

// ConvertFunction converts the matrix to a float64 matrix.
func (m *Matrix[T]) ConvertFunction() *Matrix[float64] {
	converted := New[float64](m.Rows, m.Columns)
	for i := 0; i < m.Rows; i++ {
		for j := 0; j < m.Columns; j++ {
			val := any(m.Data[i][j])
			switch v := val.(type) {
			case int:
				converted.Data[i][j] = float64(v)
			case int8:
				converted.Data[i][j] = float64(v)
			case int16:
				converted.Data[i][j] = float64(v)
			case int32:
				converted.Data[i][j] = float64(v)
			case int64:
				converted.Data[i][j] = float64(v)
			case uint:
				converted.Data[i][j] = float64(v)
			case uint8:
				converted.Data[i][j] = float64(v)
			case uint16:
				converted.Data[i][j] = float64(v)
			case uint32:
				converted.Data[i][j] = float64(v)
			case uint64:
				converted.Data[i][j] = float64(v)
			case float32:
				converted.Data[i][j] = float64(v)
			case float64:
				converted.Data[i][j] = v
			case complex64:
				converted.Data[i][j] = float64(real(v))
			case complex128:
				converted.Data[i][j] = real(v)
			default:
				// Handle other types as needed
				converted.Data[i][j] = 0.0
			}
		}
	}
	return converted
}

// Swap swaps two rows in the matrix.
func (m *Matrix[T]) Swap(row1, row2 int) error {
	if row1 < 0 || row1 >= m.Rows || row2 < 0 || row2 >= m.Rows {
		return fmt.Errorf("row index out of bounds")
	}
	m.Data[row1], m.Data[row2] = m.Data[row2], m.Data[row1]
	return nil
}

// MultiplyRow multiplies a row by a scalar value.
func MultiplyRow[T Multipliable](m *Matrix[T], row int, scalar T) error {
	if row < 0 || row >= m.Rows {
		return fmt.Errorf("row index out of bounds")
	}
	for j := 0; j < m.Columns; j++ {
		m.Data[row][j] = m.Data[row][j] * scalar
	}
	return nil
}

func AddRowMultipleGeneric[T Multipliable](m *Matrix[T], row1, row2 int, scalar T) error {
	if row1 < 0 || row1 >= m.Rows || row2 < 0 || row2 >= m.Rows {
		return fmt.Errorf("row index out of bounds")
	}

	for j := 0; j < m.Columns; j++ {
		m.Data[row2][j] = m.Data[row2][j] + (m.Data[row1][j] * scalar)
	}
	return nil
}

func (m *Matrix[T]) RankNumeric() int {
	rankMatrix := NewFromOther(m).ConvertFunction()
	rows := rankMatrix.Rows
	cols := rankMatrix.Columns
	rank := 0

	const eps = 1e-9 

	for col := 0; col < cols && rank < rows; col++ {
		maxRow := rank
		for row := rank + 1; row < rows; row++ {
			if abs(rankMatrix.Data[row][col]) > abs(rankMatrix.Data[maxRow][col]) {
				maxRow = row
			}
		}
		if abs(rankMatrix.Data[maxRow][col]) < eps {
			continue
		}

		if maxRow != rank {
			rankMatrix.Swap(maxRow, rank)
		}


		pivot := rankMatrix.Data[rank][col]
		for row := rank + 1; row < rows; row++ {
			if abs(rankMatrix.Data[row][col]) > eps {

				factor := rankMatrix.Data[row][col] / pivot

				for j := col; j < cols; j++ {
					rankMatrix.Data[row][j] -= factor * rankMatrix.Data[rank][j]
				}
			}
		}
		rank++
	}

	return rank
}


func abs(x float64) float64 {
	if x < 0 {
		return -x
	}
	return x
}
