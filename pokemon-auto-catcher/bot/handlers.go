package bot

import (
	"fmt"
	"image"
	"image/png"
	"os"
	"time"

	"github.com/go-vgo/robotgo"
	"github.com/iagotito/go-bot/pokemon-auto-catcher/actions"
	"github.com/kbinani/screenshot"
)

func handleClickingGameWindow(b *Bot) State {
	fmt.Println("clicking game window")
	robotgo.MoveMouse(b.Config.GameWindowPosition.X, b.Config.GameWindowPosition.Y)
	robotgo.Click("left")
	time.Sleep(50 * time.Millisecond)
	robotgo.KeyDown("space")
	time.Sleep(100 * time.Millisecond)
	robotgo.KeyUp("space")
	return CheckingConditions
}

func handleWalkingLeft(b *Bot) State {
	fmt.Println("walking left")

	actions.WalkLeft()

	b.NextWalkingState = WalkingRight
	return CheckingConditions
}

func handleWalkingRight(b *Bot) State {
	fmt.Println("walking right")

	actions.WalkRight()

	b.NextWalkingState = WalkingLeft
	return CheckingConditions
}

func handleCheckingConditions(b *Bot) State {
	fmt.Println("checking conditions")

	repelOff := checkRepelOff(b)
	if repelOff {
		return UsingRepel
	}

	battleStarted := checkBattle(b)
	if battleStarted {
		return Battling
	}

	return b.NextWalkingState
}

func checkBattle(b *Bot) bool {
	fmt.Println("checking battle")
	file, err := os.Open("images/no_battle_reference.png")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	savedImage, err := png.Decode(file)
	if err != nil {
		panic(err)
	}

	currentImage, err := screenshot.CaptureRect(b.Config.NoBattleReferenceImageRectangle)
	if err != nil {
		panic(err)
	}

	isSimilar, err := compareImages(savedImage, currentImage)
	if err != nil {
		panic(err)
	}

	return !isSimilar
}

func handleUsingRepel(b *Bot) State {
	fmt.Println("using repel")

	b.RepelsUsed++
	if b.RepelsUsed > 2 {
		return ResetingGame
	}

	actions.UseRepel()

	return WalkingLeft
}

func handleBattling(b *Bot) State {
	fmt.Println("battling")

	time.Sleep(500 * time.Millisecond)
	actions.Press("x")
	time.Sleep(100 * time.Millisecond)
	//if checkDesiredPokemon(b) && checkAchievmentPopup(b) {
	if checkAchievmentPopup(b) {
		return AchievmentEnabled
	}
	return RunningAway
}

func checkAchievmentPopup(b *Bot) bool {
	fmt.Println("checking achievment")

	// return true if the popup shows (the check will be false)
	return !checkImage(b.Config.AchievmentReferenceImaeRectangle, "achievment_reference.png")
}

func handleAchievmentEnabled(b *Bot) State {
	b.Stop = true
	return AchievmentEnabled
}

func handleCapturing(b *Bot) State {
	fmt.Println("capturing")

	actions.ThrowPokeball()

	if !checkDesiredPokemon(b) {
		actions.ExitCapture()
		return ResetingGame
	}

	return ThrowingPokeballAgain
}

func handleThrowingPokeballAgain(b *Bot) State {
	fmt.Println("throwing pokeball again")

	actions.ThrowPokeballAgain()

	// if the pokemon image is not the desired, it means the pokemon was captured
	time.Sleep(time.Second)
	if !checkDesiredPokemon(b) {
		actions.ExitCapture()
		return ResetingGame
	}

	return ThrowingPokeballAgain
}

func handleRunningAway(b *Bot) State {
	fmt.Println("running away")

	actions.RunAway()

	return CheckingConditions
}

func handleResetingGame(b *Bot) State {
	fmt.Println("reseting")

	actions.ResetGame()

	b.RepelsUsed = 0

	return CheckingConditions
}

func checkImage(currentRectangle image.Rectangle, referenceImageName string) bool {
	//robotgo.Move(
	//currentRectangle.Min.X,
	//currentRectangle.Min.Y,
	//)
	//time.Sleep(time.Second)
	//robotgo.Move(
	//currentRectangle.Max.X,
	//currentRectangle.Max.Y,
	//)
	//time.Sleep(time.Second)
	file, err := os.Open(fmt.Sprintf("images/%s", referenceImageName))
	if err != nil {
		panic(err)
	}
	defer file.Close()

	savedImage, err := png.Decode(file)
	if err != nil {
		panic(err)
	}

	currentImage, err := screenshot.CaptureRect(currentRectangle)
	if err != nil {
		panic(err)
	}

	currentPokemonReference := "images/current_reference.png"
	currentPokemonFile, err := os.Create(currentPokemonReference)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	err = png.Encode(currentPokemonFile, currentImage)
	if err != nil {
		panic(err)
	}

	isSimilar, err := compareImages(savedImage, currentImage)
	if err != nil {
		panic(err)
	}

	return isSimilar
}

func checkDesiredPokemon(b *Bot) bool {
	fmt.Println("checking pokemon")

	return checkImage(b.Config.PokemonReferenceImageRectangle, "pokemon_reference.png")
}

func checkRepelOff(b *Bot) bool {
	fmt.Println("checking repel")

	return checkImage(b.Config.RepelImageRectangle, "repel_off.png")
}

func compareImages(img1, img2 image.Image) (bool, error) {
	if img1.Bounds() != img2.Bounds() {
		return false, nil // Images must have the same dimensions
	}

	bounds := img1.Bounds()
	width, height := bounds.Max.X, bounds.Max.Y
	totalPixels := width * height
	differentPixels := 0

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			r1, g1, b1, a1 := img1.At(x, y).RGBA()
			r2, g2, b2, a2 := img2.At(x, y).RGBA()

			// Compare pixel values (RGBA)
			if r1 != r2 || g1 != g2 || b1 != b2 || a1 != a2 {
				differentPixels++
			}
		}
	}

	// Calculate similarity percentage
	similarity := 1.0 - float64(differentPixels)/float64(totalPixels)
	return similarity >= 0.99, nil
}
