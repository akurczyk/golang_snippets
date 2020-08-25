package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	seededRand := rand.New(rand.NewSource(time.Now().UnixNano()))
	fmt.Println(seededRand.Intn(100)) // 0-100
	fmt.Println(seededRand.Float64() * 100) // 0-100
}
