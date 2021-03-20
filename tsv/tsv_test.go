
package tsv

import (
    "testing"
    "fmt"
)

func Test1(t* testing.T) {
    var lines = ReadTsv("./seeds_dataset.txt")
    for _, l := range(lines) {
        fmt.Println(l)
    }
}
