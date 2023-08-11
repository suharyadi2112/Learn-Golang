package main

import (
    "fmt"
    "github.com/xuri/excelize/v2"
)

func main() {
    f, err := excelize.OpenFile("template_excel_hc 10k.xlsx")
    if err != nil {
        fmt.Println(err)
        return
    }

    // Get all the rows in the Sheet1.
    rows, err := f.GetRows("HC Data")
    if err != nil {
        fmt.Println(err)
        return
    }

    for _, row := range rows {
        go func(row []string) {
            // Do something with each row.
            fmt.Println(row)
        }(row)
    }

    var input string
    fmt.Scanln(&input)
}