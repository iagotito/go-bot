package main

import (
	"fmt"
	"log"

	"github.com/iagotito/go-bot/pokemon-auto-catcher/bot"
)

func main() {
	bot := bot.NewBot()

	fmt.Println("Choose action:")
	fmt.Println("1. Run bot")
	fmt.Println("2. Configure bot")

	var input string
	_, err := fmt.Scanln(&input)
	if err != nil {
		log.Fatal(err)
	}

	switch input {
	case "1":
		bot.Run()
	case "2":
		bot.Configure()
	default:
		fmt.Println("Invalid option")
	}
}
