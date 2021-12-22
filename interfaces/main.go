package main

import "fmt"

//If you are a type in this program that has a function called getGreeting, you are now an honorary member of type 'bot'
type bot interface {
	getGreeting() string
}

// Now that you are an honorary member of type 'bot', you can now call this function called printGreeting()
func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}

func (englishBot) getGreeting() string {
	// VERY custom logic for generating an english greeting
	return "Hi there!"
}

func (spanishBot) getGreeting() string {
	// VERY custom logic for generating an english greeting
	return "Hola!"
}
