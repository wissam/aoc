package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)
var global_counter int = 0

func main() {

    file, err := os.Open("input.txt")
    if err != nil {
        fmt.Println(err)
        return
    }
    defer file.Close()
    var safe int = 0
    var unsafe int = 0
    var status bool 

    fileScanner := bufio.NewScanner(file)
    for fileScanner.Scan() {
        line := fileScanner.Text()
        temp := strings.Fields(line)
        tempi := make([]int, len(temp))
        // Convert string slice to int slices
        for i := 0; i < len(temp); i++ {
            tempi[i], err = strconv.Atoi(temp[i])
            if err != nil {
                fmt.Println(err)
            }
        }
        // check if sorted ascending or descending
         status, safe, unsafe,tempi = check(tempi,safe,unsafe) 
         // fmt.Println("Status: ", status)
        if !status {
            fmt.Println("Tempi: ", tempi)
        }

    }
    fmt.Println("Unsafe: ", unsafe)
    fmt.Println("Safe: ", safe)
    fmt.Println("Total: ", safe+unsafe)
    fmt.Println("Global Counter: ", global_counter)

}




func check(tempi []int,safe int, unsafe int) (bool,int,int,[]int) {
    ordered, dampened, tempi := isOrdered(tempi)
        if ordered  {
            gradual,tempi := isGradual(tempi, dampened)
            if gradual  {
                safe++
                return true, safe, unsafe, tempi
            } else {
                unsafe++
                return false, safe, unsafe, tempi
            }
        } else {
            unsafe++
            return false, safe, unsafe, tempi
        }
}

func isGradual(intArr []int, dampened bool) (bool, []int) {
    for i := 0; i < len(intArr)-1; i++ {
        lediff := math.Abs(float64(intArr[i+1]) - float64(intArr[i]))
        if lediff < 1 || lediff > 3 {
            if dampened {
                return false, intArr
            } else {
                intArr = GradualDampener(intArr)
                break
            }
        }
    }
    for i := 0; i < len(intArr)-1; i++ {
        lediff := math.Abs(float64(intArr[i+1]) - float64(intArr[i]))
        if lediff < 1 || lediff > 3 {
            // fmt.Println("Gradual: ", intArr)
            global_counter++
            return false, intArr
        }
    }
    return true, intArr
}

func isOrdered(intArr []int) (bool, bool, []int) {
    if slices.IsSorted(intArr) {
        return true, false, intArr
    } else {
        // fmt.Println("pre reverse: ", intArr)
        slices.Reverse(intArr)
        // fmt.Println("post reverse: ", intArr)
        if slices.IsSorted(intArr) {
            return true, false, intArr
        } else {
            newIntArr := OrderDampener(intArr)
            if slices.IsSorted(newIntArr) {
                // fmt.Println("intArr: ", intArr)
                // fmt.Println("newIntArr: ", newIntArr)
                return true, true, newIntArr
            } else {
                return false, true , newIntArr
            }
        }
    }
}

func GradualDampener(intArr []int) []int {
    // fmt.Println("pregraddamp: ", intArr)
    for i := 0; i < len(intArr)-1; i++ {
        lediff := math.Abs(float64(intArr[i+1]) - float64(intArr[i]))
        if lediff < 1 || lediff > 3 {
            intArr =  append(intArr[:i], intArr[i+1:]...)
            // fmt.Println("GradualDampener: ", intArr)
            return intArr
        }
    }
    return intArr
}

func OrderDampener(intArr []int) []int {
    for i := 0; i < len(intArr)-1; i++ {
        if intArr[i] <= intArr[i+1] {
            // fmt.Println("went here")
            if i < len(intArr)-2 {
                return append(intArr[:i+1], intArr[i+2:]...)
            } else {
                return intArr[:i+1]
            }
        }
    }
    // fmt.Println("OrderedDamper: ", intArr)
    return intArr
} 
