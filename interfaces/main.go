package main

import "fmt"

type bot interface {
	getGreeting() string
}
type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}
	PrintGreeting(eb)
	PrintGreeting(sb)
}

func PrintGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func (spanishBot) getGreeting() string {
	// custom logic for spanishBot
	return "Hola!"
}

func (englishBot) getGreeting() string {
	// custom logic for englishBot
	return "Hi there"
}
