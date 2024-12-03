package puzzles

import (
	"aoc2024/util"
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
)

func Onea() error {
	lc := make(chan string)

	go func() {
		if err := util.Read("./inputs/1.txt", lc); err != nil {
			fmt.Printf("file read err %v", err)
		}
	}()

	left, right := []int{}, []int{}

	simMap := make(map[int]int)
	for line := range lc {
		// fmt.Println(line)
		nums := strings.Split(line, "   ")
		// nums[0], nums[1] = strings.Trim(nums[0], " "), strings.Trim(nums[1], " ")
		// fmt.Println(nums[0], nums[1])
		l, _ := strconv.Atoi(nums[0])
		r, _ := strconv.Atoi(nums[1])

		left = append(left, l)
		right = append(right, r)

		if _, ok := simMap[r]; !ok {
			simMap[r] = 0
		}

		simMap[r] += 1
	}

	sort.Ints(left)
	sort.Ints(right)

	resulta := 0
	resultb := 0
	for k, v := range left {
		resulta += int(math.Abs(float64(v - right[k])))

		resultb += v * simMap[v]

	}

	fmt.Println(resulta, resultb)
	return nil
}
