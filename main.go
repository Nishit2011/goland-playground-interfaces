package main

import "fmt"

type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}

//print function for english greeting
func printGreeting(eb englishBot) {
	fmt.Println(eb.getGreeting())
}

//print function for spanish greeting
func printGreeting(sb spanishBot) {
	fmt.Println(sb.getGreeting())
}

//creating two bots to generate english and spanish greetings
//the receiver value here is eb
func (englishBot) getGreeting() string {
	//very custom logic for generating an english greeting
	return "Hi There!"
}

//the receiver value here is sb
func (spanishBot) getGreeting() string {
	//very custom logic for generating an english greeting
	return "Hola!"
}
