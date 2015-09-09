package util

import (
    "fmt"
    "strconv"
    "crypto/rand"
    "golang.org/x/crypto/bcrypt"
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
func Encrypt(str string)string {
  byteValue := []byte(str)

  // Hashing the password with the cost of 10
  hashedPassword, err := bcrypt.GenerateFromPassword(byteValue, 10)
  if err != nil {
    panic(err)
  }
  return string(hashedPassword)
}
func Decrypt(hashedPassword  []byte,password string)bool {
  byteValue := []byte(password)
    // Comparing the password with the hash
    err := bcrypt.CompareHashAndPassword(hashedPassword, byteValue)
    if err == nil {
      return true
    }
    return false
}

