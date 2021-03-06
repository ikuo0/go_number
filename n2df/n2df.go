
package n2df

import (
    "math"
    "math/rand"
    
    "github.com/ikuo0/go_number/n1df"
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

func LineM(x N, row int) ([]float64) {
    return x[row]
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
        res = append(res, LineN(x, int(idx)))
    }
    return Transpose(res)
}

////////////////////////////////////////
// calculation
////////////////////////////////////////

func SubtractM(a N, b n1df.N) N {
    rowSize := len(a)
    colSize := len(a[0])
    res := New(len(a), len(a[0]))
    for m := 0; m < rowSize; m += 1 {
        for n := 0; n < colSize; n += 1 {
            res[m][n] = a[m][n] - b[n]
        }
    }
    return res
}

func SubtractX(a N, b N) N {
    rowSize := len(a)
    colSize := len(a[0])
    res := New(len(a), len(a[0]))
    for m := 0; m < rowSize; m += 1 {
        for n := 0; n < colSize; n += 1 {
            res[m][n] = a[m][n] - b[m][n]
        }
    }
    return res
}

func DivisionM(a N, b n1df.N) N {
    rowSize := len(a)
    colSize := len(a[0])
    res := New(len(a), len(a[0]))
    for m := 0; m < rowSize; m += 1 {
        for n := 0; n < colSize; n += 1 {
            res[m][n] = a[m][n] / b[n]
        }
    }
    return res
}

func Power(a N, b float64) N {
    rowSize := len(a)
    colSize := len(a[0])
    res := New(len(a), len(a[0]))
    for m := 0; m < rowSize; m += 1 {
        for n := 0; n < colSize; n += 1 {
            res[m][n] = math.Pow(a[m][n], b)
        }
    }
    return res
}

func Sqrt(a N) N {
    rowSize := len(a)
    colSize := len(a[0])
    res := New(len(a), len(a[0]))
    for m := 0; m < rowSize; m += 1 {
        for n := 0; n < colSize; n += 1 {
            res[m][n] = math.Sqrt(a[m][n])
        }
    }
    return res
}

func TotalM(x N) n1df.N {
    rowSize := len(x)
    colSize := len(x[0])
    res := n1df.Zeros(colSize)
    for m := 0; m < rowSize; m += 1 {
        res = n1df.AdditionI(res, LineM(x, m))
    }
    return res
}

func TotalN(x N) n1df.N {
    rowSize := len(x)
    res := n1df.Zeros(rowSize)
    for m := 0; m < rowSize; m += 1 {
        res[m] = n1df.Total(x[m])
    }
    return res
}

func TotalX(x N) float64 {
    rowSize := len(x)
    colSize := len(x[0])
    res := float64(0)
    for m := 0; m < rowSize; m += 1 {
        for n := 0; n < colSize; n += 1 {
            res += x[m][n]
        }
    }
    return res
}

func MeanM(x N) n1df.N {
    res := TotalM(x)
    res = n1df.Division(res, float64(len(x)))
    return res
}

func VarianceM(x N, ddof float64) n1df.N {
    mean := MeanM(x)
    diff := SubtractM(x, mean)
    pow := Power(diff, 2)
    total := TotalM(pow)
    div := n1df.Division(total, float64(len(diff)) - ddof)
    return div
}

func StddevM(x N, ddof float64) n1df.N {
    variance := VarianceM(x, 1)
    sqrt := n1df.Sqrt(variance)
    return sqrt
}



func New(rowSize int, colSize int) N {
    res := make([][]float64, rowSize)
    for i := range(res) {
        res[i] = make([]float64, colSize)
    }
    return res
}

func Arange(rowSize int, colSize int, start int) N {
    res := New(rowSize, colSize)
    counter := int(0)
    for m := 0; m < rowSize; m += 1 {
        for n := 0; n < colSize; n += 1 {
            res[m][n] = float64(counter + start)
            counter += 1
        }
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

