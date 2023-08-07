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

	for {
		screenBounds := screenshot.GetDisplayBounds(1)
		img, err := screenshot.Capture(screenBounds.Min.X, screenBounds.Min.Y, screenBounds.Dx(), screenBounds.Dy())
		if err != nil {
			log.Println("Error capturing screenshot:", err)
			continue
		}

		found := isColorFound(img, targetColor)
		if found {
			fmt.Println("Target color found!")
			// Trigger a notification here (e.g., send an email, display a message)
		} else {
			fmt.Println("Target color not found.")
		}

		time.Sleep(5 * time.Second) // Wait for 5 seconds before capturing the next screenshot
	}
}

func isColorFound(img image.Image, targetColor color.Color) bool {
	bounds := img.Bounds()
	for x := bounds.Min.X; x < bounds.Max.X; x++ {
		for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
			if img.At(x, y) == targetColor {
				return true
			}
		}
	}
	return false
}
