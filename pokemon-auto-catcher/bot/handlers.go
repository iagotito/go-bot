package bot

import (
	"fmt"
	"time"

	"github.com/go-vgo/robotgo"
)

func handleClickingGameWindow(b *Bot) State {
	fmt.Println("clicking game window")
	robotgo.MoveMouse(b.Config.GameWindowPosition.X, b.Config.GameWindowPosition.Y)
	robotgo.Click("left")
	return WalkingLeft
}

func handleWalkingLeft(b *Bot) State {
	fmt.Println("walking left")
	robotgo.KeyDown("left")
	time.Sleep(2 * time.Second)
	robotgo.KeyUp("left")
	return WalkingRight
}

func handleWalkingRight(b *Bot) State {
	fmt.Println("walking right")
	robotgo.KeyDown("right")
	time.Sleep(2 * time.Second)
	robotgo.KeyUp("right")
	return WalkingLeft
}
