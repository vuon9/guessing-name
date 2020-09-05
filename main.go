package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func main() {
	println("Guess the number!")

	rand.Seed(time.Now().UnixNano())
	secretNumber := rand.Intn(101)

	for {
		println("Please input your guess")

		var guess int
		scanner := bufio.NewScanner(os.Stdin)
		if scanner.Scan() {
			guessInt64, err := strconv.ParseInt(scanner.Text(), 10, 64)
			if err != nil {
				fmt.Printf("Your input is not a number!")
				continue
			}
			guess = int(guessInt64)
		}

		if guess > secretNumber {
			fmt.Println("Too big!")
		}

		if guess < secretNumber {
			fmt.Println("Too small!")
		}

		if guess == secretNumber {
			fmt.Println("You win!")
			break
		}

	}
}
