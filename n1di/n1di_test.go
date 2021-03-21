
package n1di

import (
    "fmt"
    "testing"
)

func Test1(t *testing.T) {
    x1 := Random(10, 1, 100)
    fmt.Println(x1)
    x2 := Sort(x1)
    fmt.Println(x2)
    x3 := Arange(3, 12, 1)
    fmt.Println(x3)
    x4 := Shuffle(x3)
    fmt.Println(x4)
    indexes := WhereEq(x3, 7)
    fmt.Println("indexes", indexes)
}

