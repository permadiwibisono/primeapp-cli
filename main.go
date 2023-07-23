package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	// print a welcome message
	intro()
	prompt()

	// create a channel to indicate when the user wants to quit
	doneChan := make(chan bool)

	// start a goroutine to read user input and run program
	go readInput(os.Stdin, doneChan)

	// block until the doneChan gets the value
	<-doneChan

	// close the channel
	close(doneChan)

	// say goodbye
	fmt.Println("Goodbye.")
}

func intro() {
	fmt.Println("Is it Prime?")
	fmt.Println("============")
	fmt.Println("Enter a whole number and we'll tell you if it is a prime number or not. Enter `q` to quit")
}

func prompt() {
	fmt.Print("-> ")
}

func readInput(reader io.Reader, doneChan chan bool) {
	scanner := bufio.NewScanner(reader)
	for {
		// read user input
		scanner.Scan()
		res, done := checkInput(scanner.Text())
		if done {
			doneChan <- true
			return
		}
		fmt.Println("result: " + res)
		prompt()
	}
}

func checkInput(input string) (string, bool) {
	// check `q` input
	if strings.EqualFold(input, "q") {
		return "", true
	}

	// try convert input to number
	n, err := strconv.Atoi(input)
	if err != nil {
		return "Please enter a whole number!", false
	}
	_, msg := isPrime(n)
	return msg, false
}

func isPrime(n int) (bool, string) {
	// 0 and 1 are not prime by definition
	if n == 0 || n == 1 {
		return false, fmt.Sprintf("%d is a not prime number by definition!", n)
	}

	// negative numbers are not prime
	if n < 0 {
		return false, "Negative numbers are not prime by definition!"
	}

	// use the modulus operator repetedly to see if we have a prime number
	for i := 2; i <= n/2; i++ {
		if n%i == 0 {
			return false, fmt.Sprintf("%d is a not prime number! It is divisible by %d", n, i)
		}
	}

	return true, fmt.Sprintf("%d is a prime number!", n)
}
