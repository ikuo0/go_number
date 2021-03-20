
package n1di

import (
    "fmt"
    "testing"
)

func Test1(t *testing.T) {
    n1di := Num1DI{}
    n1 := n1di.Arange(2, 12)
    fmt.Println(n1)
    
    n2 := n1di.Arange(2, 12)
    n3 := n2.Addition1D(n2)
    fmt.Println(n3)
    
    n4 := n3.AdditionValue(1)
    fmt.Println(n4)
    
    n5 := n4.SubtractValue(2)
    fmt.Println(n5)
    
    n6 := n5.Subtract1D(n1)
    fmt.Println(n6)
    
    n7 := n5.DivisionValue(2)
    fmt.Println(n7)
    
    n8 := n1di.Full(n1.Count, 2)
    fmt.Println(n8)
    n9 := n7.Division1D(n8)
    fmt.Println(n9)
    
    
}

func Test2(t* testing.T) {
    x := Arange(3, 12)
    fmt.Println(x)
}