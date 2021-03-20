
package num1di

import (
    "os"
)

type Num1DI struct {
    Count int
    Value []int
}

func DifferentCountExit(a Num1DI, b Num1DI) {
    if a.Count != b.Count {
        os.Exit(1)
    }
}

func (me *Num1DI) Create(count int) Num1DI {
    return Num1DI {
        Count: count,
        Value: make([]int, count),
    }
}

func (me *Num1DI) CloneIt(src Num1DI) Num1DI {
    val := make([]int, src.Count)
    copy(val, src.Value)
    res := Num1DI {
        Count: src.Count,
        Value: val,
    }
    return res
}

func (me *Num1DI) CloneThis() Num1DI {
    return me.CloneIt(*me)
}

func (me *Num1DI) V(idx int) int {
    return me.Value[idx]
}

func (me *Num1DI) R(idx int) *int {
    return &me.Value[idx]
}

func (me *Num1DI) AdditionValue(value int) Num1DI {
    res := me.CloneThis()
    for i := 0; i < me.Count; i += 1 {
        res.Value[i] += value
    }
    return res;
}

func (me *Num1DI) Addition1D(n1d Num1DI) Num1DI {
    DifferentCountExit(*me, n1d)
    res := me.CloneThis()
    for i := 0; i < me.Count; i += 1 {
        res.Value[i] += n1d.Value[i]
    }
    return res;
}

func (me *Num1DI) SubtractValue(value int) Num1DI {
    res := me.CloneThis()
    for i := 0; i < me.Count; i += 1 {
        res.Value[i] -= value
    }
    return res;
}

func (me *Num1DI) Subtract1D(n1d Num1DI) Num1DI {
    DifferentCountExit(*me, n1d)
    res := me.CloneThis()
    for i := 0; i < me.Count; i += 1 {
        res.Value[i] -= n1d.Value[i]
    }
    return res;
}

func (me *Num1DI) DivisionValue(value int) Num1DI {
    res := me.CloneThis()
    for i := 0; i < me.Count; i += 1 {
        res.Value[i] /= value
    }
    return res;
}

func (me *Num1DI) Division1D(n1d Num1DI) Num1DI {
    DifferentCountExit(*me, n1d)
    res := me.CloneThis()
    for i := 0; i < me.Count; i += 1 {
        res.Value[i] /= n1d.Value[i]
    }
    return res;
}

func (me *Num1DI) Full(count int, value int) Num1DI {
    res := me.Create(count)
    for i := 0; i < count; i += 1 {
        res.Value[i] = value
    }
    return res;
}

func (me *Num1DI) Zeros(count int) Num1DI {
    return me.Full(count, 0);
}

func (me *Num1DI) Arange(start int, end int) Num1DI {
    count := end - start
    res := me.Create(count)
    for i := 0; i < count; i += 1 {
        *res.R(i) = start + i
    }
    return res;
}

