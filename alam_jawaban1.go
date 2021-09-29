package main

import (
	"bytes"
	"fmt"
	"math/rand"
)

const letterBytes = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

func main() {
	var n int
	fmt.Scan(&n)
	result, result2 := Rand(n)

	fmt.Println("result : ", result2)
	for i := 0; i < 3; i++ {
		count := bytes.Count(result, []byte(string(result[i])))
		fmt.Println("Data ", string(result[i]), " : ", count)
	}
}

func Rand(n int) ([]byte, []string) {
	r := make([]byte, n)
	slice := []string{}
	for i := range r {
		r[i] = letterBytes[rand.Intn(len(letterBytes))]
		slice = append(slice, string(r[i]))
	}
	return r, slice
}
