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
	ConfigFileName                   string
	GameWindowPosition               ScreenPosition
	RepelImageRectangle              image.Rectangle
	NoBattleReferenceImageRectangle  image.Rectangle
	PokemonReferenceImageRectangle   image.Rectangle
	AchievmentReferenceImaeRectangle image.Rectangle
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

	c.saveImage("repel_off.png", c.RepelImageRectangle)
}

func (c *Config) ConfigNoBattleReferenceImage() {
	fmt.Println("Click on the upper left corner of the no battle reference image")
	upperLeftCorner := GetClickPosition()

	fmt.Println("Click on the bottom right corner of the no battle reference image")
	bottomRightCorner := GetClickPosition()

	c.NoBattleReferenceImageRectangle = image.Rectangle{
		Min: image.Point{X: upperLeftCorner.X, Y: upperLeftCorner.Y},
		Max: image.Point{X: bottomRightCorner.X, Y: bottomRightCorner.Y},
	}

	c.writeToFile()

	c.saveImage("no_battle_reference.png", c.NoBattleReferenceImageRectangle)
}

func (c *Config) ConfigPokemonReferenceImage() {
	fmt.Println("Click on the upper left corner of the pokemon reference image")
	upperLeftCorner := GetClickPosition()

	fmt.Println("Click on the bottom right corner of the pokemon reference image")
	bottomRightCorner := GetClickPosition()

	c.PokemonReferenceImageRectangle = image.Rectangle{
		Min: image.Point{X: upperLeftCorner.X, Y: upperLeftCorner.Y},
		Max: image.Point{X: bottomRightCorner.X, Y: bottomRightCorner.Y},
	}

	c.writeToFile()

	c.saveImage("pokemon_reference.png", c.PokemonReferenceImageRectangle)
}

func (c *Config) ConfigAchievmentPopupImage() {
	fmt.Println("Click on the upper left corner of the achievment location")
	upperLeftCorner := GetClickPosition()

	fmt.Println("Click on the bottom right corner of the achievment location")
	bottomRightCorner := GetClickPosition()

	c.AchievmentReferenceImaeRectangle = image.Rectangle{
		Min: image.Point{X: upperLeftCorner.X, Y: upperLeftCorner.Y},
		Max: image.Point{X: bottomRightCorner.X, Y: bottomRightCorner.Y},
	}

	c.writeToFile()

	c.saveImage("achievment_reference.png", c.AchievmentReferenceImaeRectangle)
}

func (c *Config) saveImage(filename string, rec image.Rectangle) {
	img, err := screenshot.CaptureRect(rec)
	if err != nil {
		panic(err)
	}

	filepath := fmt.Sprintf("images/%s", filename)
	file, err := os.Create(filepath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	err = png.Encode(file, img)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s saved\n", filename)
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
