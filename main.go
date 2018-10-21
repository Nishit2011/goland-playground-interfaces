package main

type englishBot struct{}
type spanishBot struct{}

func main() {

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
