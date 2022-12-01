package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	fmt.Println("I guessing a number between 0 and 100")
	fmt.Println("Find this number")
	omg := time.Now().Unix()
	rand.Seed(omg)
	target := rand.Intn(100) + 1
	
	reader := bufio.NewReader(os.Stdin)
	succes := false
	for guess := 0; guess < 10; guess++ {
		fmt.Println("attempts left:", 10-guess)
		fmt.Print("Enter you number: ")
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		input = strings.TrimSpace(input)

		guess, err := strconv.Atoi(input)
		if err != nil {
			log.Fatal(err)
		}

		if guess < target {
			fmt.Println("Your number is smallest")
		} else if guess > target {
			fmt.Println("You nuber bigger")
		} else {
			succes = true
			fmt.Print("You find it. Congrats!")
			break
		}
	}
	if !succes {
		fmt.Println("You lose. Number is:", target)
	}
}
