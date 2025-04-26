package config

import (
	"encoding/json"
	"fmt"
	"image"
	"image/png"
	"os"

	"github.com/kbinani/screenshot"
	gohook "github.com/robotn/gohook"
)

type ScreenPosition struct {
	X int
	Y int
}

type Config struct {
	ConfigFileName      string
	GameWindowPosition  ScreenPosition
	RepelImageRectangle image.Rectangle
}

func NewConfig() Config {
	config := Config{
		ConfigFileName: "config.json",
	}
	config.readFromFile()
	return config
}

func (c *Config) writeToFile() error {
	data, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return err
	}

	err = os.WriteFile(c.ConfigFileName, data, 0644)
	if err != nil {
		return err
	}

	return nil
}

func (c *Config) readFromFile() error {
	data, err := os.ReadFile(c.ConfigFileName)
	if err != nil {
		return err
	}

	err = json.Unmarshal(data, c)
	if err != nil {
		return err
	}

	return nil
}

func (c *Config) ConfigGameWindowPos() {
	fmt.Println("Click on the game window to register its position")

	c.GameWindowPosition = GetClickPosition()

	c.writeToFile()
}

func (c *Config) ConfigRepelOffImage() {
	fmt.Println("Click on the upper left corner of the repel message")
	upperLeftCorner := GetClickPosition()

	fmt.Println("Click on the bottom right corner of the repel message")
	bottomRightCorner := GetClickPosition()

	c.RepelImageRectangle = image.Rectangle{
		Min: image.Point{X: upperLeftCorner.X, Y: upperLeftCorner.Y},
		Max: image.Point{X: bottomRightCorner.X, Y: bottomRightCorner.Y},
	}

	c.writeToFile()

	c.saveRepelScreenshot()
}

func GetClickPosition() ScreenPosition {
	var position ScreenPosition

	eventHook := gohook.Start()
	for e := range eventHook {
		if e.Kind == gohook.MouseDown {
			if e.Button == gohook.MouseMap["left"] {
				position = ScreenPosition{X: int(e.X), Y: int(e.Y)}
				fmt.Printf("Position x:%d y:%d registered\n", e.X, e.Y)
				break
			}
		}
	}

	return position
}

func (c *Config) saveRepelScreenshot() {
	repelOffImage, err := screenshot.CaptureRect(c.RepelImageRectangle)
	if err != nil {
		panic(err)
	}
	fileName := "images/repel_off.png"

	file, err := os.Create(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	err = png.Encode(file, repelOffImage)
	if err != nil {
		panic(err)
	}
}
