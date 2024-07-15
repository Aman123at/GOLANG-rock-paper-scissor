package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

func main() {
	fmt.Println("Hello, Welcome to rock, paper, scissor game!")

	// all options
	var options = []string{"rock", "paper", "scissor"}

	rand.NewSource(int64(time.Now().Nanosecond()))

	maxRounds := 10

	// generate random number
	for {
		if maxRounds == 0 {
			break
		}

		// create a reader
		reader := bufio.NewReader(os.Stdin)

		fmt.Println("> Choose one option between rock, paper, scissor : ")

		input, err := reader.ReadString('\n')

		if err != nil {
			panic(err)
		}

		var random = rand.Intn(3)

		trimmedInput := strings.TrimSpace(input)

		randomOption := options[random]

		fmt.Printf("> You entered : %s \n", trimmedInput)

		fmt.Printf("> CPU entered : %s \n", randomOption)

		if trimmedInput == randomOption {
			fmt.Println("Draw!")
		} else if (trimmedInput == "rock" && randomOption == "scissor") || (trimmedInput == "paper" && randomOption == "rock") || (trimmedInput == "scissor" && randomOption == "paper") {
			fmt.Println("You wins!")
			break
		} else {
			fmt.Println("CPU wins!")
			break
		}

		maxRounds--
	}
}
