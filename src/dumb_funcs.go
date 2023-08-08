package godb

import "fmt"

func init() {
    fmt.Println("Runs when this package is imported elsewhere")
}

func DoubleNum(num int) int {
	return num * 2
}