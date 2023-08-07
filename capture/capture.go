package main

import (
	"fmt"
	"image/color"
	"log"
	"time"

	"github.com/kbinani/screenshot"
)

func main() {
	targetColor := color.RGBA{R: 0x9c, G: 0x27, B: 0xb0, A: 0xff}

	for {
		img, err := screenshot.CaptureScreen()
		if err != nil {
			log.Println("Error capturing screenshot:", err)
			continue
		}

		found := isColorFound(img, targetColor)
		if found {
			fmt.Println("Purple color found!")
			// Trigger a notification here (e.g., send an email, display a message)
		} else {
			fmt.Println("Purple color not found.")
		}

		time.Sleep(5 * time.Second) // Wait for 5 seconds before capturing the next screenshot
	}
}

func isColorFound(img screenshot.Image, targetColor color.Color) bool {
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