package main

import (
	"fmt"
	"image"
	"image/color"
	"log"
	"net/smtp"
	"time"

	"github.com/kbinani/screenshot"
)

func main() {
	targetColor := color.RGBA{R: 0x9c, G: 0x27, B: 0xb0, A: 0xff}
	emailAddress := "elaroldwolf@hotmail.com"
	smtpServer := "smtp.office365.com"
	smtpPort := 587
	smtpUsername := "elaroldwolf@hotmail.com"
	smtpPassword := "keerul1ne"

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
			sendEmailNotification(emailAddress, smtpServer, smtpPort, smtpUsername, smtpPassword)
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

func sendEmailNotification(emailAddress, smtpServer string, smtpPort int, smtpUsername, smtpPassword string) {
    auth := smtp.LoginAuth(smtpUsername, smtpPassword)
    to := []string{emailAddress}
    msg := []byte("To: " + emailAddress + "\r\n" +
        "Subject: Purple Color Detected\r\n" +
        "\r\n" +
        "Purple color was detected on the screen!\r\n")

    // Use "LOGIN" authentication
    err := smtp.SendMail(smtpServer+":"+fmt.Sprintf("%d", smtpPort), auth, smtpUsername, to, msg)
    if err != nil {
        log.Println("Error sending email:", err)
    }
}
