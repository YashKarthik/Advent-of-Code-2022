package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
    file, err := os.Open("./input1.txt")
    if err != nil {
        panic(err)
    }
    defer file.Close()

    caloriesOfElves := []int{0}

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {  // advances to next token; default: ReadLines
        calorieString := scanner.Text()

        calorieNum, err := strconv.Atoi(calorieString)
        if err != nil {
            caloriesOfElves = append(caloriesOfElves, 0)
            continue
        }
        caloriesOfElves[len(caloriesOfElves) - 1] = caloriesOfElves[len(caloriesOfElves) - 1] + calorieNum
    }

    // Each elem in caloriesOfElves is the total calories of a particular elf.
    sort.Ints(caloriesOfElves)

    // Part 1
    fmt.Println("Max calories among elves:", caloriesOfElves[len(caloriesOfElves)-1])
    // Part 2
    fmt.Println("Max calories among top-three elves:", caloriesOfElves[len(caloriesOfElves)-1] + caloriesOfElves[len(caloriesOfElves)-2] + caloriesOfElves[len(caloriesOfElves)-3])
    // Part 2
}
