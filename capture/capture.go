package main

import (
	"fmt"
	"image"
	"image/color"
	"log"
	"os/exec"
	"time"

	"github.com/gen2brain/beeep"
	"github.com/kbinani/screenshot"
)

var (
	previousPurplePixelCount = 0
	previousBluePixelCount   = 0
)

func main() {
	purpleColor := color.RGBA{R: 0x9c, G: 0x27, B: 0xb0, A: 0xff}
	blueColor := color.RGBA{R: 0x11, G: 0x66, B: 0xbb, A: 0xff}

	for {
		screenBounds := screenshot.GetDisplayBounds(0)
		img, err := screenshot.Capture(screenBounds.Min.X, screenBounds.Min.Y, screenBounds.Dx(), screenBounds.Dy())
		if err != nil {
			log.Println("Error capturing screenshot:", err)
			continue
		}

		purplePixelCount := getPixelCount(img, purpleColor)
		bluePixelCount := getPixelCount(img, blueColor)

		if (purplePixelCount > previousPurplePixelCount) || (bluePixelCount > previousBluePixelCount) {
			fmt.Println("Color found:", getColorName(purplePixelCount > previousPurplePixelCount, bluePixelCount > previousBluePixelCount))
			displayDesktopNotification("Color Detected", getColorName(purplePixelCount > previousPurplePixelCount, bluePixelCount > previousBluePixelCount)+" color was detected on the screen!")
			playAlarmSound()
		}

		previousPurplePixelCount = purplePixelCount
		previousBluePixelCount = bluePixelCount

		time.Sleep(5 * time.Second) // Wait for 5 seconds before capturing the next screenshot
	}
}

func getPixelCount(img image.Image, targetColor color.Color) int {
	count := 0
	bounds := img.Bounds()
	for x := bounds.Min.X; x < bounds.Max.X; x++ {
		for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
			if img.At(x, y) == targetColor {
				count++
			}
		}
	}
	return count
}

func displayDesktopNotification(title, message string) {
	err := beeep.Notify(title, message, "")
	if err != nil {
		log.Println("Error displaying desktop notification:", err)
	}
}

func playAlarmSound() {
	cmd := exec.Command("aplay", "alarm.wav") // Change "alarm.wav" to your audio file
	err := cmd.Run()
	if err != nil {
		log.Println("Error playing alarm sound:", err)
	}
}

func getColorName(purple, blue bool) string {
	if purple && blue {
		return "Purple and Blue"
	} else if purple {
		return "Purple"
	} else if blue {
		return "Blue"
	} else {
		return "Unknown"
	}
}
