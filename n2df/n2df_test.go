
package n2df

import (
    "testing"
    
    "fmt"
    
    "github.com/ikuo0/go_number/n1df"
)


func Test1(t* testing.T) {
    x1 := New(3, 4)
    fmt.Println("x1", x1)
    x2 := Random(8, 5, -10, 10)
    fmt.Println("x2", x2)
    ln := LineN(x2, 1)
    fmt.Println("LineN", ln)
    x3 := IndexingM(x2, []int{2, 3})
    fmt.Println("x3", x3)
    x4 := IndexingN(x2, []int{2, 3})
    fmt.Println("x4", x4)
    x5 := TotalM(x2)
    fmt.Println("x5", x5)
    x6 := MeanM(x2)
    fmt.Println("x6", x6)
    x7 := VarianceM(x2, 1)
    fmt.Println("x7", x7)
    x8 := StddevM(x2, 1)
    fmt.Println("x8", x8)
    
    x9 := Arange(4, 3, 1)
    fmt.Println(x9)
    x10 := n1df.Full(10, 3)
    fmt.Println(x10)
    x11 := SubtractM(x9, x10)
    fmt.Println(x11)
    x12 := DivisionM(x9, x10)
    fmt.Println(x12)
}
