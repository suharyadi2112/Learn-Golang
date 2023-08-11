package main

import (
    "fmt"

    // "github.com/360EntSecGroup-Skylar/excelize/v2"
    "github.com/xuri/excelize/v2"
)

func main() {
    // Open an existing Excel file.
    f, err := excelize.OpenFile("test.xlsx")
    if err != nil {
        panic(err)
    }

    // Set the value of cell A1 to "Hello, world!".
    sheetName := "HC Data"
    cell := "P3"
    value := "tes nama bank"
    err = f.SetCellValue(sheetName, cell, value)
    if err != nil {
        panic(err)
    }

    // Save the Excel file.
    err = f.Save()
    if err != nil {
        fmt.Println(err.Error())
    }
}