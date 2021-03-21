
package n1df

import (
    "math"
    "math/rand"
    "sort"
)

type N []float64

type IndexValue struct {
    Index int
    Value float64
}

func New(count int) N {
    res := make(N, count)
    return res
}

func Full(count int, value float64) N {
    res := New(count)
    for i := 0; i < count; i += 1 {
        res[i] = value
    }
    return res
}

func Zeros(count int) N {
    return Full(count, 0)
}

func Arange(start float64, end float64, step float64) N {
    count := int((end - start) / step)
    res := New(count)
    for i := 0; i < count; i += 1 {
        res[i] = float64(i) * step + start
    }
    return res
}

func Random(count int, min float64, max float64) N {
    res := New(count)
    randRange := max - min
    for i := 0; i < count; i += 1 {
        res[i] = float64(rand.Float64() * float64(randRange)) + min
    }
    return res
}

////////////////////////////////////////
// operation
////////////////////////////////////////
func ArgSort(x N) []int {
    count := len(x)
    idxval := make([]IndexValue, count)
    for i := 0; i < count; i += 1 {
        idxval[i] = IndexValue {i, x[i]}
    }
    sort.Slice(idxval, func(i, j int) bool {
        return idxval[i].Value < idxval[j].Value
    })
    indexes := make([]int, count)
    for i := 0; i < count; i += 1 {
        indexes[i] = idxval[i].Index
    }
    return indexes
}

func Sort(x N) N {
    indexes := ArgSort(x)
    count := len(x)
    res := New(count)
    for i := 0; i < count; i += 1 {
        res[i] = x[indexes[i]]
    }
    return res
}

func Shuffle(x N) N {
    count := len(x)
    rnd := Random(count, 0, math.MaxInt32)
    indexes := ArgSort(rnd)
    res := New(count)
    for i := 0; i < count; i += 1 {
        res[i] = x[indexes[i]]
    }
    return res
}

func ArgMin(x N) int {
    indexes := ArgSort(x)
    return indexes[0]
}


////////////////////////////////////////
// calculation
////////////////////////////////////////
func AdditionI(a N, b N) N {
    count := len(a)
    res := New(count)
    for i := 0; i < count; i += 1 {
        res[i] = a[i] + b[i]
    }
    return res
}

func Addition(a N, b float64) N {
    count := len(a)
    res := New(count)
    for i := 0; i < count; i += 1 {
        res[i] = a[i] + b
    }
    return res
}

func SubtractI(a N, b N) N {
    count := len(a)
    res := New(count)
    for i := 0; i < count; i += 1 {
        res[i] = a[i] - b[i]
    }
    return res
}

func Division(a N, b float64) N {
    count := len(a)
    res := New(count)
    for i := 0; i < count; i += 1 {
        res[i] = a[i] / b
    }
    return res
}

func Power(a N, b float64) N {
    count := len(a)
    res := New(count)
    for i := 0; i < count; i += 1 {
        res[i] = float64(math.Pow(float64(a[i]), b))
    }
    return res
}

func Sqrt(a N) N {
    count := len(a)
    res := New(count)
    for i := 0; i < count; i += 1 {
        res[i] = float64(math.Sqrt(float64(a[i])))
    }
    return res
}

func Total(x N) float64 {
    res := float64(0)
    count := len(x)
    for i := 0; i < count; i += 1 {
        res += x[i]
    }
    return res
}
