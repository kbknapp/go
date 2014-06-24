package math

type Matrix struct {
	M     []int
	Size  int
	Rows  [][]int
	Cols  [][]int
	Diags [][]int
}

func NewMatrix(s int) Matrix {
	m := Matrix{
		M:     make([]int, s*s),
		Size:  s,
		Rows:  make([][]int, s),
		Cols:  make([][]int, s),
		Diags: make([][]int, 2),
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
	m.Diags[0] = make([]int, s)
	j := 0
	for i := 0; i < (s * s); i += (s + 1) {
		m.Diags[0][j] = i
		j++
	}
	m.Diags[1] = make([]int, s)
	j = 0
	for i := (s - 1); i < ((s * s) - 1); i += (s - 1) {
		m.Diags[1][j] = i
		j++
	}

	return m
}
