package main

import (
	"fmt"
	"log"

	"github.com/akhmadreiza/reilearngo/greetings"
)

func main() {
	log.SetPrefix("main: ")

	message, err := greetings.Hello("Rei")
	
	if (err != nil) {
		log.Fatal(err)
	}

	fmt.Println(message)

	//call Hellos
	names := []string {"Reiza", "Nita", "Kaindra"}

	message_hellos, err := greetings.Hellos(names)

	if (err != nil) {
		log.Fatal(err)
	}

	fmt.Println(message_hellos["Reiza"])
	fmt.Println(message_hellos["Nita"])
	fmt.Println(message_hellos["Kaindra"])
}