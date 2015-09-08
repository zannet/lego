package util

import (
    "fmt"
    "strconv"
    "crypto/rand"
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
func GenarateId(strSize int, randType string) string {

   var dictionary string

   if randType == "alphanum" {
       dictionary = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
   }

   if randType == "alpha" {
       dictionary = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
   }

   if randType == "number" {
       dictionary = "0123456789"
   }

   var bytes = make([]byte, strSize)
   rand.Read(bytes)
   for k, v := range bytes {
       bytes[k] = dictionary[v%byte(len(dictionary))]
   }
   return string(bytes)
}