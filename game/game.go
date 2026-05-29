package game

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/kiing-dom/megamath/common"
)

type GameState struct {
	Question     string
	CorrectIndex int
	Answers      []Answer
	Lives        int
	Points       int
}

type Answer struct {
	Value    int
	Position common.Vec2
}

func (gs *GameState) NewQuestion() {
	source := rand.NewSource(time.Now().UnixNano())
	localRand := rand.New(source)

	ops := []string{"+", "-", "*"}
	a := localRand.Intn(20) + 1
	b := localRand.Intn(20) + 1
	op := ops[localRand.Intn(len(ops))]

	var correct int
	switch op {
	case "+":
		correct = a + b
	case "-":
		correct = a - b
	case "*":
		correct = a * b
	}

	gs.Question = fmt.Sprintf("%d %s %d = ?", a, op, b)

	// generate different answers
	numAnswers := 4
	correctIndex := localRand.Intn(numAnswers)

	gs.Answers = make([]Answer, numAnswers)
	for i := 0; i < numAnswers; i++ {
		if i == correctIndex {
			gs.Answers[i] = Answer{Value: correct}
		} else {
			wrong := correct + localRand.Intn(10) - 5
			for wrong == correct {
				wrong = correct + rand.Intn(10) - 5
			}

			gs.Answers[i] = Answer{Value: wrong}
		}
	}

	gs.CorrectIndex = correctIndex
}
