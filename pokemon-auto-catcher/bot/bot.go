package bot

import (
	"fmt"
	"log"
	"time"

	"github.com/go-vgo/robotgo"
	"github.com/iagotito/go-bot/pokemon-auto-catcher/config"
	gohook "github.com/robotn/gohook"
)

type State int

const (
	ClickingGameWindow State = iota
	WalkingLeft
	WalkingRight
	CheckingConditions
	UsingRepel
)

type StateHandler func(b *Bot) State

var stateHandlers = map[State]StateHandler{
	ClickingGameWindow: handleClickingGameWindow,
	WalkingLeft:        handleWalkingLeft,
	WalkingRight:       handleWalkingRight,
	CheckingConditions: handleCheckingConditions,
	UsingRepel:         handleUsingRepel,
}

type Bot struct {
	CurrentState     State
	DefaultDelay     time.Duration
	Stop             bool
	Config           config.Config
	NextWalkingState State
}

func NewBot() *Bot {
	return &Bot{
		CurrentState:     ClickingGameWindow,
		DefaultDelay:     100 * time.Millisecond,
		Stop:             false,
		Config:           config.NewConfig(),
		NextWalkingState: WalkingLeft,
	}
}

func (b *Bot) Configure() {
	fmt.Println("What do you want to configure?")
	fmt.Println("1. Game window position")
	fmt.Println("2. Repel off message rectangle")

	var input string
	_, err := fmt.Scanln(&input)
	if err != nil {
		log.Fatal(err)
	}

	switch input {
	case "1":
		b.Config.ConfigGameWindowPos()
	case "2":
		b.Config.ConfigRepelOffImage()
	default:
		fmt.Println("Invalid option")
	}
}

func (b *Bot) Run() {
	go func() {
		eventHook := gohook.Start()
		for e := range eventHook {
			if e.Kind == gohook.KeyDown && e.Keychar == 'q' {
				fmt.Println("Stoping bot.")
				b.Stop = true
			}
		}
	}()

	time.Sleep(time.Second)

	for {
		if b.Stop {
			robotgo.KeyDown("space")
			time.Sleep(100 * time.Millisecond)
			robotgo.KeyUp("space")
			break
		}
		handler, _ := stateHandlers[b.CurrentState]

		b.CurrentState = handler(b)

		time.Sleep(50 * time.Millisecond)
	}
}
