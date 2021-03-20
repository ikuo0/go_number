
package n2df

import (
    "math/rand"
    
    "github.com/ikuo0/go_number/n1di"
)

type N [][]float64

func LineN(x N, col int) ([]float64) {
    res := make([]float64, len(x))
    for i := 0; i < len(x); i += 1 {
        res[i] = x[i][col]
    }
    return res
}

func Transpose(x N) N {
    rowSize := len(x)
    colSize := len(x[0])
    res := New(colSize, rowSize)
    for m := 0; m < rowSize; m += 1 {
        for n := 0; n < colSize; n += 1 {
            res[n][m] = x[m][n]
        }
    }
    return res
}

func IndexingM(x N, indexes n1di.N) N {
    res := make([][]float64, 0)
    for _, idx := range(indexes) {
        res = append(res, x[idx])
    }
    return res
}

func IndexingN(x N, indexes n1di.N) N {
    res := make([][]float64, 0)
    for _, idx := range(indexes) {
        res = append(res, LineN(x, idx))
    }
    return res
}



func New(rowSize int, colSize int) N {
    res := make([][]float64, rowSize)
    for i := range(res) {
        res[i] = make([]float64, colSize)
    }
    return res
}

func Random(rowSize int, colSize int, min float64, max float64) N {
    res := New(rowSize, colSize)
    randRange := max - min
    for m := 0; m < rowSize; m += 1 {
        for n := 0; n < colSize; n += 1 {
            res[m][n] = rand.Float64() * randRange + min
        }
    }
    return res
}

