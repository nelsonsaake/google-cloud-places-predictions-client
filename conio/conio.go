package conio

import (
	"fmt"
	"log"

	"github.com/eiannone/keyboard"
)

var (
	ch       = make(chan string)
	text     = ""
	onChange func(string) error
)

func StartListening() {
	if err := keyboard.Open(); err != nil {
		log.Fatal(err)
	}
}

func StopListening() {
	defer keyboard.Close()
}

func AddListener() {
	// prompt
	print(">> ")

	// read an print
	for {
		char, key, err := keyboard.GetKey()
		if err != nil {
			log.Fatal(err)
		}

		if key == keyboard.KeyEsc {
			break
		}

		fmt.Printf("%c", char)
	}
}
