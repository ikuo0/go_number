
package n2df

import (
    "testing"
    
    "fmt"
)


func Test1(t* testing.T) {
    x1 := New(3, 4)
    fmt.Println("x1", x1)
    x2 := Random(8, 5, -100, 100)
    fmt.Println("x2", x2)
    ln := LineN(x2, 1)
    fmt.Println("LineN", ln)
    x3 := IndexingM(x2, []int{2, 3})
    fmt.Println("x3", x3)
    x4 := IndexingN(x2, []int{2, 3})
    fmt.Println("x4", x4)
}
