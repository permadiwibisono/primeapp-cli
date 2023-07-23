package main

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
)

func Test_isPrime(t *testing.T) {
	tests := []struct {
		title          string
		input          int
		expectedResult bool
		expectedMsg    string
	}{
		{"zero_number", 0, false, "0 is a not prime number by definition!"},
		{"one_number", 1, false, "1 is a not prime number by definition!"},
		{"negative_number", -7, false, "Negative numbers are not prime by definition!"},
		{"prime_number", 7, true, "7 is a prime number!"},
		{"not_prime_number", 8, false, "8 is a not prime number! It is divisible by 2"},
	}

	for _, entity := range tests {
		result, msg := isPrime(entity.input)
		if entity.expectedResult != result {
			t.Errorf("%s: expected \"%t\" but got \"%t\"", entity.title, entity.expectedResult, result)
		}
		if entity.expectedMsg != msg {
			t.Errorf("%s: expected \"%s\" but got \"%s\"", entity.title, entity.expectedMsg, msg)
		}
	}
}

func Test_prompt(t *testing.T) {
	// save a copy of io.Stdout
	oldStdout := os.Stdout

	// create a read and write pipe
	r, w, _ := os.Pipe()

	// set os.Stdout to our write pipe
	os.Stdout = w

	prompt()

	_ = w.Close()

	// reset os.Stdout to what it was before
	os.Stdout = oldStdout

	// read the output of our prompt() func from our read pipe
	out, _ := io.ReadAll(r)

	// perform our test
	if string(out) != "-> " {
		t.Errorf("incorrect prompt: expected \"-> \" but got \"%s\"", string(out))
	}
}

func Test_intro(t *testing.T) {
	// save a copy of io.Stdout
	oldStdout := os.Stdout

	// create a read and write pipe
	r, w, _ := os.Pipe()

	// set os.Stdout to our write pipe
	os.Stdout = w

	intro()

	_ = w.Close()

	// reset os.Stdout to what it was before
	os.Stdout = oldStdout

	// read the output of our intro() func from our read pipe
	out, _ := io.ReadAll(r)

	// perform our test
	if !strings.Contains(string(out), "Enter a whole number") {
		t.Errorf("intro text is incorrect: got \"%s\"", string(out))
	}
}

func Test_checkInput(t *testing.T) {
	tests := []struct {
		title    string
		input    string
		expected string
	}{
		{title: "empty_input", input: "", expected: "Please enter a whole number!"},
		{title: "random_alphabet_input", input: "lorem ipsum", expected: "Please enter a whole number!"},
		{title: "quit_input", input: "q", expected: ""},
		{title: "quit_uppercase_input", input: "Q", expected: ""},
		{title: "number_input", input: "7", expected: "7 is a prime number!"},
		{title: "decimal_input", input: "1.1", expected: "Please enter a whole number!"},
	}
	for _, entity := range tests {
		res, _ := checkInput(entity.input)
		if !strings.EqualFold(res, entity.expected) {
			t.Errorf("%s: expected \"%s\", but got \"%s\"", entity.title, entity.expected, res)
		}
	}
}

func Test_readInput(t *testing.T) {
	// to test this function, we need a channel, and an instance of an io.Reader
	doneChan := make(chan bool)
	// save a copy of io.Stdout
	oldStdout := os.Stdout

	// create a read and write pipe
	r, w, _ := os.Pipe()

	// set os.Stdout to our write pipe
	os.Stdout = w

	// create a ref to a bytes.Buffer
	var stdin bytes.Buffer

	stdin.Write([]byte("1\nq\n"))

	go readInput(&stdin, doneChan)
	<-doneChan
	close(doneChan)

	_ = w.Close()

	// reset os.Stdout to what it was before
	os.Stdout = oldStdout

	// read the output of our prompt() func from our read pipe
	out, _ := io.ReadAll(r)

	expected := "1 is a not prime number by definition!"

	// perform our test
	if !strings.Contains(string(out), expected) {
		t.Errorf("incorrect: expected contains \"%s\" but got \"%s\"", expected, string(out))
	}
}
