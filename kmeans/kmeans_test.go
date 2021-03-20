
package kmeans

import (
    "testing"
    
    "fmt"
    "strconv"
    "github.com/ikuo0/go_number/tsv"
    "github.com/ikuo0/go_number/n2df"
)

func ToFloat(X [][]string) (error, n2df.N) {
    rowSize := len(X)
    colSize := len(X[0])
    res := n2df.New(rowSize, colSize)
    for i, row := range(X) {
        for j, col := range(row) {
            if n, err := strconv.ParseFloat(X[i][j], 64); err != nil {
                return err, n2df.N{}
            } else {
                res[i][j] = n
            }
        }
    }
    return nil, res
}

func Test1(t* testing.T) {
    if err, tsvData := tsv.ReadTsv("./seeds_dataset.txt"); err != nil {
        fmt.Println(err)
    } else {
        if err, data := ToFloat(tsvData); err != nil {
            fmt.Println(err)
        } else {
            fmt.Println(data)
        }
    }
}
