package puzzles

import (
	"aoc2024/util"
	"fmt"
	"math"
	"strconv"
	"strings"
)

func Twoa() error {
	lc := make(chan string)

	go func() {
		if err := util.Read("./inputs/2.txt", lc); err != nil {
			fmt.Printf("file read err %v", err)
		}
	}()

	result := 0
	for line := range lc {
		arr := make([]int, 0)
		for _, num := range strings.Fields(line) {
			n, _ := strconv.Atoi(num)
			arr = append(arr, n)
		}

		descending := false

		if arr[0] > arr[1] {
			descending = true
		} else if arr[0] == arr[1] {
			continue
		}

		safe := true
		for i := 1; i < len(arr); i++ {
			// Abs diff >= 1 and <= 3
			diff := math.Abs(float64(arr[i] - arr[i-1]))
			if diff < 1 || diff > 3 {
				safe = false
				break
			}

			if !descending && arr[i] <= arr[i-1] {
				safe = false
				break
			}

			if descending && arr[i] >= arr[i-1] {
				safe = false
				break
			}
		}

		if safe {
			result++
		} else {
			fmt.Println("UNSAFE:", arr, descending)
		}

	}

	fmt.Println(result)
	return nil
}
