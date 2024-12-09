package main

import (
    "os"
    "bufio"
    "strings"
    "strconv"
    "sort"
    "fmt"
)

func Abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}

func Counter(slice []int, val int) (result int) {
    for _, v := range slice {
        if v == val {
            result += 1
        }
    }
    return result
}

func Load(path string) (left, right []int){
    data, err := os.Open(path)

    if err != nil {
        return left, right
    }
    scanner := bufio.NewScanner(data)

    for scanner.Scan() {
        line := strings.Split(scanner.Text(), "   ")

        part1, _ := strconv.Atoi(line[0])
        part2, _ := strconv.Atoi(line[1])
        left = append(left, part1)
        right = append(right, part2)
    }
    return left, right
}

func Distance(path string) (result int) {
    pole1, pole2 := Load(path)
    sort.Slice(pole1, func(i, j int) bool {
        return pole1[i] < pole1[j]
    })
    sort.Slice(pole2, func(i, j int) bool {
        return pole2[i] < pole2[j]
    })

    var sum int = 0
    for i, _ := range pole1 {
        sum += Abs(pole1[i] - pole2[i])   
    }
    return sum
}

func Similarity(path string) (result int) {
    left, right := Load(path)
    for _, v := range left {
        result += Counter(right, v) * v
    }
    return result
}

func main(){
    fmt.Printf("%v is total distance\n", Distance("input"))
    fmt.Printf("%v is total similarity\n", Similarity("input"))
}
