package main

import (
	"math/rand"
	"time"
)

var omikuji = []string{
	"1",
	"1",
	"2",
	"3",
	"4",
	"5",
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	result := omikuji[rand.Intn(len(omikuji))]
	println(result)
}
