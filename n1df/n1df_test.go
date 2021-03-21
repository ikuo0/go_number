
package n1df

import (
    "fmt"
    "testing"
)

func Test1(t *testing.T) {
    x1 := Zeros(10)
    fmt.Println(x1)
    x2 := Full(10, 9)
    fmt.Println(x2)
    x3 := Random(10, 1, 10)
    fmt.Println(x3)
    x4 := AdditionI(x2, x3)
    fmt.Println(x4)
    x5 := Addition(x4, 100.0)
    fmt.Println(x5)
    x6 := Total(x5)
    fmt.Println(x6)
    x7 := Sqrt(x3)
    fmt.Println(x7)
    argmin := ArgMin(x3)
    fmt.Println("argmin", argmin)
    x8 := Sort(x3)
    fmt.Println(x8)
    x9 := Arange(10, 21, 1.1)
    fmt.Println(x9)
    x10 := Shuffle(x9)
    fmt.Println(x10)
}

