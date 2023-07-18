package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

//       [M]             [Z]     [V]
//       [Z]     [P]     [L]     [Z] [J]
//   [S] [D]     [W]     [W]     [H] [Q]
//   [P] [V] [N] [D]     [P]     [C] [V]
//   [H] [B] [J] [V] [B] [M]     [N] [P]
//   [V] [F] [L] [Z] [C] [S] [P] [S] [G]
//   [F] [J] [M] [G] [R] [R] [H] [R] [L]
//   [G] [G] [G] [N] [V] [V] [T] [Q] [F]
//    1   2   3   4   5   6   7   8   9
//   01234567890123456789012345678901234

func Day5() {
    file, err := os.Open("./input5.txt")
    if err != nil {
        panic(err)
    }
    defer file.Close()

    stacks := [9][]string{}
    locations := [9]int{1,5,9,13,17,21,25,29,33}

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        row := scanner.Text()

        // skip empty lines
        if len(row) == 0 { continue }

        // converts ascii drawing to stacks
        if row[0] != 'm' {
            for i, loc := range locations {
                if row[loc] == '1' {
                    break
                } else if row[loc] != ' ' {
                    stacks[i] = append(stacks[i], string(row[loc]))
                }
            }
            continue
        }

        // parse move commands into array
        movementArgs := []int{}
        movementRowSlice := strings.Fields(row) // splits with whitespace as delimiter

        for i, elem := range movementRowSlice {
            if i == 0 || i == 2 || i == 4 { continue }

            moveInt, err := strconv.Atoi(elem)
            if err != nil { panic(err) }

            movementArgs = append(movementArgs, moveInt)
        }

        // execute the move commands
        cratesToBeMoved := stacks[movementArgs[1]-1][:movementArgs[0]] // topmost crate is at [0]
        stacks[movementArgs[1]-1] = stacks[movementArgs[1]-1][movementArgs[0]:] // remove the crates from it's current stack

        reverseArray(stacks[movementArgs[2]-1])
        stacks[movementArgs[2]-1] = append(stacks[movementArgs[2]-1], cratesToBeMoved...)// put cratesToBeMoved at the top of destination stack
        reverseArray(stacks[movementArgs[2]-1])
    }

    for _, stack := range stacks {
        fmt.Print(stack[0])
    }
    fmt.Println()
}

func reverseArray(arr []string) {
    for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
        arr[i], arr[j] = arr[j], arr[i]
    }
}
