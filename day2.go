package main

import (
	"bufio"
	"fmt"
	//"math"
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
    points1 := 0
    points2 := 0

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        round := scanner.Text()

        // Part 1
        // points1 += weapons[string(round[2])]
        // if weapons[string(round[0])] == weapons[string(round[2])] {
        //     points1 += 3
        //     continue
        // } 

        // if math.Abs(float64(weapons[string(round[0])] - weapons[string(round[2])])) > 1 {
        //     if weapons[string(round[0])] >= weapons[string(round[2])] {
        //         points1 += 6
        //     }
        // } else {
        //     if weapons[string(round[0])] <= weapons[string(round[2])] {
        //         points1 += 6
        //     }

        // }

        // Part 2
        if string(round[2]) == "X" { // need to lose
            points2 += 0 // losing
            points2 += wrappedSub(weapons[string(round[0])])
        } else if string(round[2]) == "Y" { // draw
            points2 += 3
            points2 += weapons[string(round[0])] // same weapon since we're drawing
        } else if string(round[2]) == "Z" { // win
            points2 += 6
            points2 += wrappedAdd(weapons[string(round[0])])
        }
    }

    fmt.Println(points1)
    fmt.Println(points2)
}

func wrappedAdd(start int) int {
    if start == 3 { return 1 }
    start += 1
    return start
}

func wrappedSub(start int) int {
    if start == 1 { return 3 }
    start -= 1
    return start
}
