package actions

import (
	"fmt"
	"time"

	"github.com/go-vgo/robotgo"
)

const walkDelay = 250 * time.Millisecond

func walk(direction string) {
	robotgo.KeyDown(direction)
	time.Sleep(walkDelay)
	robotgo.KeyUp(direction)
}

func WalkLeft() {
	walk("left")
}

func WalkRight() {
	walk("right")
}

func WalkUp() {
	walk("up")
}

func WalkDown() {
	walk("down")
}

func Press(keyname string) {
	fmt.Printf("pressing %s\n", keyname)
	robotgo.KeyDown(keyname)
	time.Sleep(75 * time.Millisecond)
	robotgo.KeyUp(keyname)
}

func UseRepel() {
	Press("space")
	time.Sleep(500 * time.Millisecond)
	Press("x")
	time.Sleep(500 * time.Millisecond)
	Press("enter")
	time.Sleep(500 * time.Millisecond)
	Press("up")
	time.Sleep(500 * time.Millisecond)
	Press("up")
	time.Sleep(500 * time.Millisecond)
	Press("down")
	time.Sleep(500 * time.Millisecond)
	Press("down")
	time.Sleep(500 * time.Millisecond)
	Press("x")
	time.Sleep(1000 * time.Millisecond)
	Press("x")
	time.Sleep(1000 * time.Millisecond)
	Press("x")
	time.Sleep(1000 * time.Millisecond)
	Press("x")
	time.Sleep(1000 * time.Millisecond)
	Press("z")
	time.Sleep(1000 * time.Millisecond)
	Press("z")
	time.Sleep(1000 * time.Millisecond)
	Press("space")
}

func RunAway() {
	time.Sleep(500 * time.Millisecond)
	Press("space")
	time.Sleep(100 * time.Millisecond)
	Press("right")
	time.Sleep(100 * time.Millisecond)
	Press("down")
	time.Sleep(100 * time.Millisecond)
	Press("x")
	time.Sleep(50 * time.Millisecond)
	Press("space")
	time.Sleep(100 * time.Millisecond)
	Press("z")
	time.Sleep(100 * time.Millisecond)
	Press("z")
	time.Sleep(100 * time.Millisecond)
	Press("z")
	time.Sleep(100 * time.Millisecond)
}

func ThrowPokeball() {
	time.Sleep(500 * time.Millisecond)
	Press("space")
	time.Sleep(100 * time.Millisecond)
	Press("down")
	time.Sleep(1000 * time.Millisecond)
	Press("x")
	time.Sleep(1500 * time.Millisecond)
	Press("right")
	time.Sleep(1000 * time.Millisecond)
	Press("x")
	time.Sleep(500 * time.Millisecond)
	Press("x")
	time.Sleep(500 * time.Millisecond)
	Press("space")
	time.Sleep(3000 * time.Millisecond)
}

func ThrowPokeballAgain() {
	time.Sleep(1000 * time.Millisecond)
	Press("x")
	time.Sleep(500 * time.Millisecond)
	Press("x")
	time.Sleep(500 * time.Millisecond)
	Press("space")
	time.Sleep(500 * time.Millisecond)
	Press("x")
	time.Sleep(500 * time.Millisecond)
	Press("x")
	time.Sleep(500 * time.Millisecond)
	Press("space")
	time.Sleep(3000 * time.Millisecond)
}

func ExitCapture() {
	Press("z")
	time.Sleep(750 * time.Millisecond)
	Press("z")
	time.Sleep(750 * time.Millisecond)
	Press("z")
	time.Sleep(750 * time.Millisecond)
	Press("z")
	time.Sleep(750 * time.Millisecond)
}

func ResetGame() {
	Press("h")
	time.Sleep(1000 * time.Millisecond)
	Press("x")
	time.Sleep(1500 * time.Millisecond)
	Press("x")
	time.Sleep(1500 * time.Millisecond)
	Press("x")
	time.Sleep(1500 * time.Millisecond)
	Press("x")
	time.Sleep(1500 * time.Millisecond)
	Press("x")
	time.Sleep(2000 * time.Millisecond)
}
