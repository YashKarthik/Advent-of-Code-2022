package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Day3() {
    file, err := os.Open("./input3.txt")
    if err != nil {
        panic(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    priorities := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
    //characterMap := map[string]bool{}
    commons := map[string]int{}

    for scanner.Scan() {
        rucksack := scanner.Text()
        compartment_1 := rucksack[:len(rucksack)/2]
        compartment_2 := rucksack[len(rucksack)/2:]

        //for _, elem := range compartment_1 {
        //    characterMap[string(elem)] = true
        //}

        //for _, elem := range compartment_2 {
        //    if characterMap[string(elem)] == true {
        //        commons[string(elem)] += 1
        //    }
        //}
        for _, s := range compartment_2 {
            if strings.Index(compartment_1, string(s)) != -1 {
                fmt.Println(compartment_1, compartment_2, string(s))
                commons[string(s)] += 1
                break
            }
        }

    }
    priorities_sum := 0
    for str, times := range commons {
        priority := strings.Index(priorities, string(str)) + 1
        priorities_sum += priority * times
    }

    fmt.Println(priorities_sum)
}
