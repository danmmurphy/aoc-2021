// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"io"
// 	"log"
// 	"os"
// 	"strconv"
// )

// func check(e error) {
// 	if e != nil {
// 		panic(e)
// 	}
// }

// func main() {
// 	file, err := os.Open("/home/dan/aoc-2021/day_1/input.txt")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer file.Close()

// 	scanner := bufio.NewScanner(file)
// 	counter := 0
// 	var prev_depth int

// 	for scanner.Scan() {
// 		prev_depth, err = strconv.Atoi(scanner.Text())
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		break
// 	}
// 	if scanner.Err() != nil {
// 		log.Println(scanner.Err())
// 	}
// 	_, err = file.Seek(0, io.SeekStart)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	scanner = bufio.NewScanner(file)
// 	for scanner.Scan() {
// 		curr_depth, err := strconv.Atoi(scanner.Text())
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		if curr_depth > prev_depth {
// 			counter += 1
// 		}
// 		prev_depth = curr_depth
// 	}
// 	fmt.Println(counter)
// }
