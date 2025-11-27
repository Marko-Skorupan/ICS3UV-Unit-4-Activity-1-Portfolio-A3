/*
 * @author Marko Skorupan
 * @version 1.0.0
 * @date 2025-11-26
 * @fileoverview Validate end > start and display range.
 */

package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
)

func main() {
    reader := bufio.NewReader(os.Stdin)

    fmt.Print("Enter Start Value: ")
    startStr, _ := reader.ReadString('\n')
    startStr = strings.TrimSpace(startStr)
    start, _ := strconv.Atoi(startStr)

    fmt.Print("Enter End Value: ")
    endStr, _ := reader.ReadString('\n')
    endStr = strings.TrimSpace(endStr)
    end, _ := strconv.Atoi(endStr)

    for end <= start {
        fmt.Print("Sorry, your ending value must be larger. Enter another ending value: ")
        endStr, _ = reader.ReadString('\n')
        endStr = strings.TrimSpace(endStr)
        end, _ = strconv.Atoi(endStr)
    }

    for i := start; i <= end; i++ {
        fmt.Println(i)
    }
}
