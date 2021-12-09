package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
    "strconv"
)

func main() {
    file, err := os.Open("../inputs/day1.txt")

    if err != nil {
        log.Fatal(err)
    }

    defer file.Close()

    scanner := bufio.NewScanner(file)
    scanner.Split(bufio.ScanWords)

    var arr []int

    for scanner.Scan() {
        x, err := strconv.Atoi(scanner.Text())

        if err != nil {
            log.Fatal(err)
        }

        arr = append(arr, x)
    }

    var count = 0

    prev := arr[0]

    for i := 1; i < len(arr); i ++ {

        curr := arr[i]

        if curr > prev {
            count++
        }

        prev = curr
    }

    fmt.Println(count)



}

