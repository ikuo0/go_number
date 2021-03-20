
package tsv

import (
    "testing"
    "fmt"
)

func Test1(t* testing.T) {
    if err, lines := ReadTsv("./seeds_dataset.txt"); err != nil {
        fmt.Println(err)
    } else {
        fmt.Println("len(lines)", len(lines))
        /*for _, l := range(lines) {
            fmt.Println(l)
        }*/
    }
}
