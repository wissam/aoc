package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

// type partials struct {
//     start int
//     end int
//     ledo string
// }

func main() {
    // scan file
    var matches [][]string
    var index [][]int
    var counter int = 0
    file, err := os.Open("input.txt")
    if err != nil {
        fmt.Println(err)
        return
    }
    defer file.Close()
    scanner := bufio.NewScanner(file)
    // I have a feeling that I am over complicating this again and there are other ways to do it.
    // or else how can people get the answer in a few mins??
    for scanner.Scan() {
       // fmt.Println(scanner.Text())
        counter++
        fmt.Println("counter: ", counter)
        line := scanner.Text()
        fmt.Println("old line:",len(line))
        //counting don't() 
        re1 := regexp.MustCompile(`don't\(\)`)
        donts := re1.FindAllStringIndex(line,-1)
        counterdonts := len(donts)
        fmt.Println("indexdonts: ", donts)
        fmt.Println("counterdonts: ", counterdonts)
        re2 := regexp.MustCompile(`do\(\)`)
        dos := re2.FindAllStringIndex(line,-1)
        counterdos := len(dos)
        fmt.Println("indexdos: ", dos)
        fmt.Println("counterdos: ", counterdos)
        newline := ""
        r := regexp.MustCompile(`don't\(\).*?do\(\)`)
        index = r.FindAllStringIndex(line,-1)
        fmt.Println("index: ", index)
        for i, _ := range index {
            var start int
            var end int
            if i ==  0 {
                start = 0
            } else {
                start = index[i-1][1]
            }
            if i == len(index) - 1 {
                end = len(line)
            } else {
                end = index[i+1][0]
            }
            newline += line[start:index[i][0]] + line[index[i][1]:end]
        }
        // fmt.Println(newline)
        fmt.Println("new line:",len(newline))
       re := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)
        matches = append(matches, re.FindAllStringSubmatch(newline,-1)...)
    }

    total := 0
    for _, match := range matches {
        // fmt.Println(match)
        x , _ := strconv.Atoi(match[1])
        y , _ := strconv.Atoi(match[2])
        total += x * y
    }
    // fmt.Println("index: ", index)
    fmt.Println("allcounter: ", counter)

    fmt.Println("Total: ", total)
}


