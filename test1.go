
package main

import (
    "fmt"
)

func Create(Count int) []interface{} {
    return make([]interface{}, Count)
}

func Zeros(src []interface{}) {
    for i := 0; i < len(src); i += 1 {
        src[i] = 0
    }
}

func Replace1(src []interface{}) []interface{} {
    length := len(src)
    dst := make([]interface{}, length)
    copy(dst, src)
    tmp := src[0]
    dst[0] = src[length - 1]
    dst[length - 1] = tmp
    return dst
}

func TestFunc1(src interface{}) [][]float64 {
    if x, ok := src.([][]float64); ok {
        dst := make([][]float64, len(x))
        copy(dst, x)
        tmp := x[0]
        dst[0] = x[len(x) - 1]
        dst[len(x) - 1] = tmp
        return dst
    } else {
        return [][]float64{}
    }
}

func main() {
    fmt.Printf("Hello\n")
    x1 := Create(10)
    Zeros(x1)
    
    x2 := [][]float64{
        {0, 0},
        {0, 1, 9},
        {1, 0},
        {1, 1},
    }
    fmt.Println(x2)
    
    x3 := TestFunc1(x2)
    fmt.Println(x3)
    
    x4 := []int{1, 2, 3, 4}
    fmt.Println(x4)
    x5 := x4
    x5[1] = 9
    fmt.Println(x4)
    fmt.Println(x5)
}
