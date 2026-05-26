package main

import (
	"fmt"
	"os"

	"golang.org/x/term"
)

var ClearScreenAnsi = "\033[H\033[2J"

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

	b := make([]byte, 1)
	for {
		if b[0] == 'q' || b[0] == 3 {
			break
		}

		os.Stdin.Read(b)
		fmt.Printf("Key pressed: %q\n", b[0])
	}
}

func clearScreen() {
	fmt.Println(ClearScreenAnsi)
}
