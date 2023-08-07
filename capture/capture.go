package main

import (
	"fmt"
	"image"
	"image/color"
	"log"
	"time"

	"github.com/kbinani/screenshot"
)

func main() {
	targetColor := color.RGBA{R: 0x9c, G: 0x27, B: 0xb0, A: 0xff}

	var previousScreen image.Image

	for {
		screenBounds := screenshot.GetDisplayBounds(0)
		img, err := screenshot.Capture(screenBounds.Min.X, screenBounds.Min.Y, screenBounds.Dx(), screenBounds.Dy())
		if err != nil {
			log.Println("Error capturing screenshot:", err)
			continue
		}

		if previousScreen != nil && isColorChange(previousScreen, img, targetColor) {
			fmt.Println("New purple color detected!")
			// Trigger a notification here (e.g., send an email, display a message)
		}

		previousScreen = img

		time.Sleep(5 * time.Second) // Wait for 5 seconds before capturing the next screenshot
	}
}

func isColorChange(prevScreen, currentScreen image.Image, targetColor color.Color) bool {
	bounds := prevScreen.Bounds()
	for x := bounds.Min.X; x < bounds.Max.X; x++ {
		for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
			prevColor := prevScreen.At(x, y)
			currColor := currentScreen.At(x, y)

			if prevColor != currColor && currColor == targetColor {
				return true
			}
		}
	}
	return false
}
