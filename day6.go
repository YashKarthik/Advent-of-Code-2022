package main

import (
	"bufio"
	"fmt"
	"os"
)

func Day6() {

    file, err := os.Open("./input6.txt")
    if err != nil {
        panic(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    scanner.Scan()
    signalBuffer := scanner.Text()

    for index := range signalBuffer {
        markerTest := NewSet()
        markerTest.bulkAdd(signalBuffer[index:index+14])

        count := 0
        for _, val := range markerTest.contains {
            if val == true { count++ }
        }

        if count == 14 {
            fmt.Println(index+14)
            return
        }
    }
}
