package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Day4() {
    file, err := os.Open("./input4.txt")
    if err != nil {
        panic(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    inclusive_range := 0
    for scanner.Scan() {
        assignedSectionString := scanner.Text()
        // assignedSectionArray will be an array of two elems eg: ["8-9", "10-14"]
        assignedSectionArray := strings.FieldsFunc(assignedSectionString, func(r rune) bool {
            if r == ',' {
                return true
            }
            return false
        })

        sections := []int{}
        for _, elfRangeString := range assignedSectionArray {
            // again elfRange will be an array of two elems eg: ["7", "8"]
            elfRange := strings.FieldsFunc(elfRangeString, func(r rune) bool {
                if r == '-' {
                    return true
                }
                return false
            })

            for _, elfRangeInt := range elfRange {
                i, e := strconv.Atoi(elfRangeInt)
                if e != nil { panic(e) }
                sections = append(sections, i)
            }
        }

        fmt.Println("Ranges:", sections)

        if sections[0] <= sections[2] && sections[1] >= sections[3] {
            inclusive_range++
            fmt.Println("Inclusive")
        } else if sections[0] >= sections[2] && sections[1] <= sections[3] {
            inclusive_range++
            fmt.Println("Inclusive")
        }

    }

    fmt.Println("\n-------\n", "Total inclusive range pairs:", inclusive_range)
}
