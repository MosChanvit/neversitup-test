package main

import (
	"fmt"
	"log"
	"neversitup/service"
)

func main() {
	// Test cases 1
	fmt.Println("Test cases 1")
	fmt.Println(service.Permute("a"))
	fmt.Println(service.Permute("ab"))
	fmt.Println(service.Permute("abc"))
	fmt.Println(service.Permute("aabb"))
	fmt.Println("")

	// // Test cases 2
	fmt.Println("Test cases 2")
	fmt.Println(service.FindOdd([]int{7}))
	fmt.Println(service.FindOdd([]int{0}))
	fmt.Println(service.FindOdd([]int{1, 1, 2}))
	fmt.Println(service.FindOdd([]int{0, 1, 0, 1, 0}))
	fmt.Println(service.FindOdd([]int{1, 2, 2, 3, 3, 3, 4, 3, 3, 3, 2, 2, 1}))
	fmt.Println("")

	// Test cases 3
	fmt.Println("Test cases 3")
	fmt.Println(service.CountSmileys([]string{":)", ";(", ";}", ":-D"}))       // Output: 2
	fmt.Println(service.CountSmileys([]string{";D", ":-(", ":-)", ";~)"}))     // Output: 3
	fmt.Println(service.CountSmileys([]string{";]", ":[", ";*", ":$", ";-D"})) // Output: 1
	fmt.Println("")

	log.Println("END")

}
