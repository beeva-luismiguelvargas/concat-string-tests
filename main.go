package main

import (
	"bytes"
	"fmt"
	"math/rand"
	"time"
)

const maxIterations = 200000

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func GenerateRandom() string {
	return RandStringRunes(1)
}

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func ConcatenatePlus() string {
	myVar := ""
	for i := 0; i < maxIterations; i++ {
		myVar = myVar + GenerateRandom()
	}
	return myVar
}

func ConcatenateBuffer() string {
	var buffer bytes.Buffer
	for i := 0; i < maxIterations; i++ {
		buffer.WriteString(GenerateRandom())
	}
	return buffer.String()
}

func ConcatenateSprint() string {
	myVar := ""
	for i := 0; i < maxIterations; i++ {
		myVar = fmt.Sprint(myVar, GenerateRandom())
	}
	return myVar
}

func ConcatenateSprintf() string {
	myVar := ""
	for i := 0; i < maxIterations; i++ {
		myVar = fmt.Sprintf("%v%v",myVar, GenerateRandom())
	}
	return myVar
}


func main() {
	start := time.Now()
	ConcatenatePlus()
	elapsed := time.Since(start)
	fmt.Printf("ConcatenatePlus - Elapsed time is %s\n", elapsed)

	start = time.Now()
	ConcatenateBuffer()
	elapsed = time.Since(start)
	fmt.Printf("ConcatenateBuffer - Elapsed time is %s\n", elapsed)

	start = time.Now()
	ConcatenateSprint()
	elapsed = time.Since(start)
	fmt.Printf("ConcatenateSprint - Elapsed time is %s\n", elapsed)

	start = time.Now()
	ConcatenateSprintf()
	elapsed = time.Since(start)
	fmt.Printf("ConcatenateSprint - Elapsed time is %s\n", elapsed)

}
