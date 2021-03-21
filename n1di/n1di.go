
package n1di

import (
    //"fmt"
    "math"
    "math/rand"
    "sort"
)

type N []int

type IndexValue struct {
    Index int
    Value int
}

func New(count int) N {
    res := make(N, count)
    return res
}

func Full(count int, value int) N {
    res := New(count)
    for i := 0; i < count; i += 1 {
        res[i] = value
    }
    return res
}

func Zeros(count int) N {
    return Full(count, 0)
}

func Arange(start int, end int, step int) N {
    count := int((end - start) / step)
    res := New(count)
    for i := 0; i < count; i += 1 {
        res[i] = int(i) * step + start
    }
    return res
}

func Random(count int, min int, max int) N {
    res := New(count)
    randRange := max - min
    for i := 0; i < count; i += 1 {
        res[i] = int(rand.Float64() * float64(randRange)) + min
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

func WhereEq(x N, n int) N {
    res := make([]int, 0)
    count := len(x)
    for i := 0; i < count; i += 1 {
        if x[i] == n {
            res = append(res, i)
        }
    }
    return res
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

func Addition(a N, b int) N {
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

func Division(a N, b int) N {
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
        res[i] = int(math.Pow(float64(a[i]), b))
    }
    return res
}

func Sqrt(a N) N {
    count := len(a)
    res := New(count)
    for i := 0; i < count; i += 1 {
        res[i] = int(math.Sqrt(float64(a[i])))
    }
    return res
}

func Total(x N) int {
    res := int(0)
    count := len(x)
    for i := 0; i < count; i += 1 {
        res += x[i]
    }
    return res
}


