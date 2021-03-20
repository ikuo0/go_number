

package tsv

import (
    "encoding/csv"
    //"fmt"
    "os"
)

func Read(fileName string, comma rune) [][]string {
    file, err := os.Open(fileName)
    if err != nil {
        panic(err)
    }
    defer file.Close()

    reader := csv.NewReader(file)
    reader.Comma = comma;
    var line []string
    var lines [][]string;

    for {
        line, err = reader.Read()
        if err != nil {
            break
        }
        //fmt.Println(line)
        lines = append(lines, line)
    }
    return lines
}

func ReadTsv(fileName string) [][]string {
    return Read(fileName, '\t')
}