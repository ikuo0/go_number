

package tsv

import (
    "encoding/csv"
    //"errors"
    //"fmt"
    "os"
)

func Read(fileName string, comma rune) (error, [][]string) {
    file, err := os.Open(fileName)
    if err != nil {
        return err, nil
    }
    defer file.Close()

    reader := csv.NewReader(file)
    reader.Comma = comma
    var line []string
    var lines [][]string

    for {
        line, err = reader.Read()
        if err != nil {
            break
        }
        lines = append(lines, line)
    }
    
    // csv.NewReaderは異なる桁数が現れた時点で読み込みを辞めているのでこういったチェックは不要
    //var columns int = 0
    //for i, l := range(lines) {
    //    if columns == 0 {
    //        columns = len(l)
    //    } else if(columns != len(l)) {
    //        return errors.New(fmt.Sprintf("different column size %d: %d != %d", i, columns, len(l))), nil
    //    }
    //}
    
    return nil, lines
}

func ReadTsv(fileName string) (error, [][]string) {
    return Read(fileName, '\t')
}

