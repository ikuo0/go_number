
package n2df

import (
    "testing"
    
    "fmt"
)


func Test1(t* testing.T) {
    x := New(3, 4)
    fmt.Println(x)
    x = Random(8, 2, -100, 100)
    fmt.Println(x)
    ln := LineN(x, 1)
    fmt.Println(ln)
    x = Indexing(x, []int{2, 3})
    fmt.Println(x)
}
