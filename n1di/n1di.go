
package n1di

import (
    "math"
    "math/rand"
    //"os"
    "sort"
)

type N []int

type IndexValue struct {
    Index int
    Value int
}

func New(count int) N {
    res := make([]int, count)
    return res
}

func Arange(start int, end int) N {
    count := end - start
    res := New(count)
    for i := 0; i < count; i += 1 {
        res[i] = start + i
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
