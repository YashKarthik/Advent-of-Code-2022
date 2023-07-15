package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func Day2() {
    file, err := os.Open("./input2.txt")
    if err != nil {
        panic(err)
    }
    defer file.Close()

    weapons := map[string]int{
        "A": 1,
        "B": 2,
        "C": 3,
        "X": 1,
        "Y": 2,
        "Z": 3,
    }
    points := 0

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        round := scanner.Text()

        points += weapons[string(round[2])]
        if weapons[string(round[0])] == weapons[string(round[2])] {
            points += 3
            continue
        } 

        if math.Abs(float64(weapons[string(round[0])] - weapons[string(round[2])])) > 1 {
            if weapons[string(round[0])] >= weapons[string(round[2])] {
                points += 6
            }
        } else {
            if weapons[string(round[0])] <= weapons[string(round[2])] {
                points += 6
            }

        }
    }

    fmt.Println(points)
}
