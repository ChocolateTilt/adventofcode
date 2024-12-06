package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	leftList, rightList, err := createLists("day01_input.txt")
	if err != nil {
		panic(err)
	}

	fmt.Printf("Left len: %v\nRight len: %v\n", len(leftList), len(rightList))

	total := 0
	for i := 0; i < len(leftList); i++ {
		dist := 0
		lv := leftList[i]
		rv := rightList[i]

		if lv > rv {
			dist = lv - rv
		} else {
			dist = rv - lv
		}

		total = dist + total

	}
	fmt.Println("Total: ", total)
}

func createLists(f string) (leftList, rightList []int, err error) {
	file, err := os.Open(f)
	if err != nil {
		return nil, nil, err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		l, r, exists := strings.Cut(line, "   ")
		if !exists {
			fmt.Println("did not find sep, skipping")
			continue
		}

		li, err := strconv.Atoi(l)
		if err != nil {
			return nil, nil, err
		}

		ri, err := strconv.Atoi(r)
		if err != nil {
			return nil, nil, err
		}

		leftList = append(leftList, li)
		rightList = append(rightList, ri)
	}

	if err := scanner.Err(); err != nil {
		return nil, nil, err
	}

	slices.Sort(leftList)
	slices.Sort(rightList)

	return
}
