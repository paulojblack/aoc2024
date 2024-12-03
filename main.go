package main

import (
	"aoc2024/puzzles"
	"log"
)

func main() {

	err := puzzles.Onea()

	if err != nil {
		log.Fatalf("we got a probbem %v", err)
	}
}
