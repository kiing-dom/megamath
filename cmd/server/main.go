package main

import (
	"fmt"
	"os"
	"time"

	"github.com/kiing-dom/megamath/entities"
	"golang.org/x/term"
)

var ClearScreenAnsi = "\033[H\033[2J"
var MoveCursorAnsi = "\033[%d;%df"
var EndOfTextAscii = 3

func main() {
	enableRawMode()
}

func enableRawMode() {
	player := entities.Player{
		Position: entities.Vec2{X: 5, Y: 5},
		Velocity: entities.Vec2{X: 1, Y: 1},
		Char:     '@',
	}
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
			handleInput(key, &player)
		case <-ticker.C:
			updateGameState(&player)
		}
	}
}

func handleInput(key byte, p *entities.Player) {
	if key == 'w' {
		p.MoveUp()
	} else if key == 's' {
		p.MoveDown()
	} else if key == 'a' {
		p.MoveLeft()
	} else if key == 'd' {
		p.MoveRight()
	} else {
		fmt.Printf("Key pressed:%q\n", key)
	}
}

func clearScreen() {
	fmt.Println(ClearScreenAnsi)
}

func updateGameState(p *entities.Player) {
	player := fmt.Sprintf(MoveCursorAnsi, p.Position.Y, p.Position.X)
	fmt.Print(player)
}
