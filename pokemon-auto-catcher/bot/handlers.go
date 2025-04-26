package bot

import (
	"fmt"
	"image"
	"image/png"
	"os"
	"time"

	"github.com/go-vgo/robotgo"
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
	robotgo.KeyDown("left")
	time.Sleep(75 * time.Millisecond)
	robotgo.KeyUp("left")
	b.NextWalkingState = WalkingRight
	return CheckingConditions
}

func handleWalkingRight(b *Bot) State {
	fmt.Println("walking right")
	robotgo.KeyDown("right")
	time.Sleep(75 * time.Millisecond)
	robotgo.KeyUp("right")
	b.NextWalkingState = WalkingLeft
	return CheckingConditions
}

func handleCheckingConditions(b *Bot) State {
	fmt.Println("checking conditions")

	repelOff := checkRepelOff(b)
	if repelOff {
		return UsingRepel
	}

	return b.NextWalkingState
}

func handleUsingRepel(b *Bot) State {
	robotgo.KeyDown("space")
	time.Sleep(100 * time.Millisecond)
	robotgo.KeyUp("space")

	fmt.Println("using repel")

	fmt.Println("pressing x")
	robotgo.KeyDown("x")
	time.Sleep(100 * time.Millisecond)
	robotgo.KeyUp("x")

	time.Sleep(500 * time.Millisecond)

	fmt.Println("pressing enter")
	robotgo.KeyDown("enter")
	time.Sleep(100 * time.Millisecond)
	robotgo.KeyUp("enter")

	time.Sleep(500 * time.Millisecond)

	fmt.Println("pressing up")
	robotgo.KeyDown("up")
	time.Sleep(100 * time.Millisecond)
	robotgo.KeyUp("up")

	time.Sleep(500 * time.Millisecond)

	fmt.Println("pressing up")
	robotgo.KeyDown("up")
	time.Sleep(100 * time.Millisecond)
	robotgo.KeyUp("up")

	time.Sleep(500 * time.Millisecond)

	fmt.Println("pressing down")
	robotgo.KeyDown("down")
	time.Sleep(100 * time.Millisecond)
	robotgo.KeyUp("down")

	time.Sleep(500 * time.Millisecond)

	fmt.Println("pressing down")
	robotgo.KeyDown("down")
	time.Sleep(100 * time.Millisecond)
	robotgo.KeyUp("down")

	time.Sleep(500 * time.Millisecond)

	fmt.Println("pressing x")
	robotgo.KeyDown("x")
	time.Sleep(100 * time.Millisecond)
	robotgo.KeyUp("x")

	time.Sleep(1000 * time.Millisecond)

	fmt.Println("pressing x")
	robotgo.KeyDown("x")
	time.Sleep(100 * time.Millisecond)
	robotgo.KeyUp("x")

	time.Sleep(1000 * time.Millisecond)

	fmt.Println("pressing x")
	robotgo.KeyDown("x")
	time.Sleep(100 * time.Millisecond)
	robotgo.KeyUp("x")

	time.Sleep(1000 * time.Millisecond)

	fmt.Println("pressing x")
	robotgo.KeyDown("x")
	time.Sleep(100 * time.Millisecond)
	robotgo.KeyUp("x")

	time.Sleep(1000 * time.Millisecond)

	fmt.Println("pressing z")
	robotgo.KeyDown("z")
	time.Sleep(100 * time.Millisecond)
	robotgo.KeyUp("z")

	time.Sleep(1000 * time.Millisecond)

	fmt.Println("pressing z")
	robotgo.KeyDown("z")
	time.Sleep(100 * time.Millisecond)
	robotgo.KeyUp("z")

	time.Sleep(1000 * time.Millisecond)

	robotgo.KeyDown("space")
	time.Sleep(100 * time.Millisecond)
	robotgo.KeyUp("space")

	return WalkingLeft
}

func checkRepelOff(b *Bot) bool {
	file, err := os.Open("images/repel_off.png")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	savedImage, err := png.Decode(file)
	if err != nil {
		panic(err)
	}

	currentImage, err := screenshot.CaptureRect(b.Config.RepelImageRectangle)
	if err != nil {
		panic(err)
	}

	isSimilar, err := compareImages(savedImage, currentImage)
	if err != nil {
		panic(err)
	}

	return isSimilar
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
