package main

import (
	"fmt"
	"os"
	"time"

	"golang.org/x/term"
)

var ClearScreenAnsi = "\033[H\033[2J"
var EndOfTextAscii = 3

func main() {
	enableRawMode()
}

func enableRawMode() {
	clearScreen()
	oldState, err := term.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		panic(err)
	}

	defer term.Restore(int(os.Stdin.Fd()), oldState)

	inputCh := make(chan byte)
	go func() {
		b := make([]byte, 1)
		for {
			os.Stdin.Read(b)
			inputCh <- b[0]
		}
	}()

	ticker := time.NewTicker(100 * time.Millisecond)
	defer ticker.Stop()

	for {
		select {
		case key := <-inputCh:
			if key == 'q' || key == byte(EndOfTextAscii) {
				return
			}
			handleInput(key)
		case <-ticker.C:
			updateGameState()
		}
	}
}

func handleInput(key byte) {
	fmt.Printf("Key pressed:%q\n", key)
}

func clearScreen() {
	fmt.Println(ClearScreenAnsi)
}

func updateGameState() {
	fmt.Println("gs tick")
}
