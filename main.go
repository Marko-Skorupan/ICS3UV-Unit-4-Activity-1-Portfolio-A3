/*
 * @author Marko Skorupan
 * @version 1.0.0
 * @date 2025-11-26
 * @fileoverview Validate end > start and display range.
 */

package main

import "fmt"

func main() {
	var start, end int

	fmt.Print("Enter Start Value: ")
		fmt.Scan(&start)

		fmt.Print("Enter End Value: ")
		fmt.Scan(&end)

		for end <= start {
			fmt.Print("Sorry, ending value must be larger. Enter another ending value: ")
			fmt.Scan(&end)
		}

		for counter := start; counter <= end; counter++ {
			fmt.Println(counter)
		}
	fmt.Println("\nDone.")
}