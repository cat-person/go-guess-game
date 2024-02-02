package main

import (
	"fmt"
	"math/rand"
	"strconv"
)


const biggestPossibleNumberToGuess = 100;

func main() {
    fmt.Println("Please Guess the number in range 0..99")
	var numberToGuess = rand.Intn(biggestPossibleNumberToGuess)
	fmt.Println("The number is %d", numberToGuess)
	var guessCount = 0
	
	guessedRight := false
	for !guessedRight {
		fmt.Printf("Try to guess number\n")
		var guess string

		_, err := fmt.Scanln(&guess)

		if err != nil {
			fmt.Println(err)
		} else {
			guessInt, err := strconv.Atoi(guess)
			if err != nil {
				fmt.Println(err)
			} else {
				if guessInt < 0 || guessInt >= biggestPossibleNumberToGuess {
					fmt.Println("Please choose positive number below 100")
				} else {
					guessCount ++
		
					var message string 
					guessedRight, message = evaluateAnswer(numberToGuess, guessInt)
			
					fmt.Println(message);
				}	
			}
		}		
	} 
	fmt.Printf("Guess the number with %d attempts\n", guessCount)
}

func evaluateAnswer(answer, guess int) (quessedRight bool, message string) {
	// This function need to return 2 values 
	// 1) guessedRight boolean value which should be true, if guess is equal to answer and false otherwise
	// 2) message should indicate if guess is bigger or smaller than the answer and  

	return answer == guess, "Logic hasn't been implemented"
}