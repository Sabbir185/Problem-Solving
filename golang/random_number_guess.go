package main

import (
	"fmt"
	"math/rand"
	"time"
)

func random(min int, max int) int {
	return rand.Intn(max-min) + min
}

func main() {
	rand.Seed(time.Now().UnixNano())
	number := random(0, 100)

	for {
		var user_input int
		fmt.Print("Enter your number: ")
		fmt.Scan(&user_input)

		if user_input > number {
			fmt.Println("Enter lower number")
		} else if user_input < number {
			fmt.Println("Enter bigger number")
		} else {
			fmt.Println("Congratulation!...")
			break
		}
	}
}
