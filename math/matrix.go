package math

type Matrix struct {
	M    []int
	Size int
	Rows [][]int
	Cols [][]int
}

func NewMatrix(s int) Matrix {
	m := Matrix{
		M:    make([]int, s*s),
		Size: s,
		Rows: make([][]int, s),
		Cols: make([][]int, s),
	}

	r := 0
	c := 0
	for i := 0; i < s; i++ {
		m.Rows[i] = make([]int, s)
		m.Cols[i] = make([]int, s)
		for j := 0; j < s; j++ {
			m.Rows[i][j] = r
			r++
			m.Cols[i][j] = (j * s) + c
		}
		c++
	}

	return m
}
