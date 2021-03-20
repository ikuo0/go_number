
package n2df

import (
)

type N [][]float64

func New(rowSize int, colSize int) N {
    res := make([][]float64, rowSize)
    for i := range(res) {
        res[i] = make([]float64, colSize)
    }
    return res
}
