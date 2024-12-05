package main

import (
	"aoc2024/puzzles"
	"log"
)

func main() {

	err := puzzles.Twoa()

	if err != nil {
		log.Fatalf("we got a probbem %v", err)
	}
}
