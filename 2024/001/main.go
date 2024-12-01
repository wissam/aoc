package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

//Warning: I am NOT a real programmer. I am pretty sure there are better ways to do this!!!
func main() {
    file, err := os.Open("input.txt")
    var intArr1 = make([]int, 0)
    var intArr2 = make([]int, 0)
    total := 0
    if err != nil {
        fmt.Println(err)
        return
    }
    defer file.Close()
    
    fileScanner := bufio.NewScanner(file)
    for fileScanner.Scan() {
        line := fileScanner.Text()
        temp := strings.Fields(line)
        num0, _ := strconv.Atoi(temp[0])
        num1, _ := strconv.Atoi(temp[1])
        intArr1 = append(intArr1, num0)
        intArr2 = append(intArr2, num1)

    }
    //need to sort the arrays.
    sort.Ints(intArr1)
    sort.Ints(intArr2)
    
    for i := 0; i < len(intArr1); i++ {
        // fmt.Println(intArr1[i], intArr2[i])
        distance := math.Abs(float64(intArr1[i]) - float64(intArr2[i]))
        total += int(distance)
    }
    fmt.Println(total)

    //find similarities between the two arrays
    similarities := 0
    for i := 0; i < len(intArr1); i++ {
        count := 0
        for j := 0; j < len(intArr2); j++ {
            if intArr1[i] == intArr2[j] {
               count++ 
            }
        }
        similarities += intArr1[i] * count
    }
    fmt.Println(similarities)
}
