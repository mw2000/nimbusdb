// package main

// import (
// "bytes"
// "encoding/gob"
// "fmt"
// "os"
// )

// func main() {
// f, _ := os.Create("a.dfile")
// fmt.Println(f)

// // int to buffer
// data := uint64(1234567890)
// var buf bytes.Buffer
// enc := gob.NewEncoder(&buf)

// enc.Encode(data)
// fmt.Println(buf.Bytes())

// // buffer to int
// var x uint64
// debuf := bytes.NewBuffer(buf.Bytes())
// dec := gob.NewDecoder(debuf)
// dec.Decode(&x)

// fmt.Println(x)
// }