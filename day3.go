package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Day3Part1() {
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

type Set struct {
    contains  map[string]bool
}

func NewSet() *Set {
    return &Set{contains: map[string]bool{}}
}

func (set *Set) add(item string) {
    set.contains[item] = true
}

func (set *Set) bulkAdd(items string) {
    for _, char := range items {
        set.add(string(char))
    }
}

func (set *Set) remove(item string) {
    set.contains[item] = false
}

func (a Set) intersection(b Set) []string {
    var commonChars []string

    for key := range a.contains {
        if b.contains[key] == a.contains[key] {
            commonChars = append(commonChars, key)
        }
    }
    return commonChars
}

func Day3Part1WithSets() {
    file, err := os.Open("./input3.txt")
    if err != nil {
        panic(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    priorities := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
    priorities_sum := 0

    for scanner.Scan() {
        rucksack := scanner.Text()

        compartment_1 := NewSet()
        compartment_2 := NewSet()
        compartment_1.bulkAdd(rucksack[:len(rucksack)/2])
        compartment_2.bulkAdd(rucksack[len(rucksack)/2:])

        commonChar := compartment_1.intersection(*compartment_2)
        priorities_sum += strings.Index(priorities, string(commonChar[0])) + 1
    }

    fmt.Println(priorities_sum)
}

func Day3Part2WithSets() {
    file, err := os.Open("./input3.txt")
    if err != nil {
        panic(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    priorities := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
    priorities_sum := 0

    for scanner.Scan() {
      elf1 := scanner.Text()
      scanner.Scan()
      elf2 := scanner.Text()
      scanner.Scan()
      elf3 := scanner.Text()
      elf1Set := NewSet()
      elf2Set := NewSet()
      elf3Set := NewSet()
      elf1Set.bulkAdd(elf1)
      elf2Set.bulkAdd(elf2)
      elf3Set.bulkAdd(elf3)

      intersection1 := NewSet()
      for _, elem := range elf2Set.intersection(*elf1Set) {
          intersection1.add(elem)
      }

      badge := elf3Set.intersection(*intersection1)[0]
      priorities_sum += strings.Index(priorities, string(badge)) + 1
    }

    fmt.Println(priorities_sum)
}
