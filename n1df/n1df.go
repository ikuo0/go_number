
package n1df

import (
    "math"
    "math/rand"
)

type N []float64

func New(count int) N {
    res := make([]float64, count)
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

func Random(count int, min float64, max float64) N {
    res := New(count)
    randRange := max - min
    for i := 0; i < count; i += 1 {
        res[i] = rand.Float64() * randRange + min
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

func Addition(a N, b float64) N {
    count := len(a)
    res := New(count)
    for i := 0; i < count; i += 1 {
        res[i] = a[i] + b
    }
    return res
}

func Subtract(a N, b N) N {
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

func Sqrt(a N) N {
    count := len(a)
    res := New(count)
    for i := 0; i < count; i += 1 {
        res[i] = math.Sqrt(a[i])
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
