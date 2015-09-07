package util

import (
    "fmt"
    "strconv"
)

func StringToInt(str string)int {
    // string to int
    i, err := strconv.Atoi(str)
    if err != nil {
        // handle error
        fmt.Println(err)
    }
    fmt.Println(str, i)
    return i
}
func IntToString(i int)string {
    // string to int
    str := strconv.Itoa(i)
    fmt.Println(str)
    return str
}
func Increment(i int)string {
    // string to int
    incremented := i + 1
    str := strconv.Itoa(incremented)
    fmt.Println(str)
    return str
}
