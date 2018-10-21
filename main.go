package main

import "fmt"

//interface being used to define what different type of function with return type it should have
//and anyone that matches the above decription becomes a type bot
//therefore english and spanish bot both become type bot as well i.e. we dont write englishBot implements bot

//interfaces are implicit so we dont need to define any specific links b/w interfaces and the bots
type bot interface {

	//a method getGreeting that return type string becomes of type bot
	//spansh and english bot both have function getGreeting becomes a bot type as well
	getGreeting() string
}

type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
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
