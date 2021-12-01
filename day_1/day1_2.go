package main

import (
	"bufio"
	"container/list"
	"fmt"
	"log"
	"os"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	file, err := os.Open("/home/dan/aoc-2021/day_1/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	counter := 0
	window := list.New()
	var old_sum int

	for scanner.Scan() {
		curr_depth, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}

		if window.Len() < 3 {
			old_sum += curr_depth
			window.PushBack(curr_depth)
			continue
		}

		new_sum := old_sum - window.Remove(window.Front()).(int) + curr_depth
		window.PushBack(curr_depth)

		if new_sum > old_sum {
			counter += 1
		}
		old_sum = new_sum
	}
	fmt.Println(counter)
}
