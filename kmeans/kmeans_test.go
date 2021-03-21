
package kmeans

import (
    "testing"
    
    "fmt"
    "os"
    "strconv"
    "github.com/ikuo0/go_number/tsv"
    "github.com/ikuo0/go_number/n1di"
    "github.com/ikuo0/go_number/n2df"
    "github.com/ikuo0/go_number/scaler"
)

func ToFloat(X [][]string) (error, n2df.N) {
    rowSize := len(X)
    colSize := len(X[0])
    res := n2df.New(rowSize, colSize)
    for i, row := range(X) {
        for j, col := range(row) {
            if n, err := strconv.ParseFloat(col, 64); err != nil {
                return err, n2df.N{}
            } else {
                res[i][j] = n
            }
        }
    }
    return nil, res
}

func ReadData() n2df.N {
    if err, tsvData := tsv.ReadTsv("./seeds_dataset.txt"); err != nil {
        fmt.Println(err)
        os.Exit(9)
    } else {
        if err, data := ToFloat(tsvData); err != nil {
            fmt.Println(err)
            os.Exit(9)
        } else {
            return data
        }
    }
    return nil
}

func Test1(t* testing.T) {
    data := ReadData()
    //fmt.Println(data)
    //yData := n2df.LineN(data, 7)
    xIndexes := n1di.Arange(0, 7)
    xData := n2df.IndexingN(data, xIndexes)
    
    stdScaler := scaler.Standardization()
    scaled := stdScaler.FitTransoform(xData)
    
    fmt.Println(scaled)
}
