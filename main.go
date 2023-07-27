package main

import (
	"fmt"
	"github.com/sqweek/dialog"
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"os"
)

func main() {
	// Load the image you want to overlay onto the white background
	filePath, err := dialog.File().Filter("All Files|*.*").Title("Select Image File").Load()
	if err != nil {
		panic(err)
	}
	directoryPath, err := dialog.Directory().Title("Select Directory").Browse()
	if err != nil {
		panic(err)
	}

	overlayImage, err := loadImage(filePath)
	if err != nil {
		panic(err)
	}

	// Define the dimensions for the white background image
	backgroundWidth := overlayImage.Bounds().Dx()
	backgroundHeight := overlayImage.Bounds().Dy()

	// Create the white background image
	whiteBackground := createWhiteBackground(backgroundWidth, backgroundHeight)

	// Overlay the original image onto the white background
	resultImage := overlayImages(whiteBackground, overlayImage)

	// Save the result image
	saveImageWithUniqueName(directoryPath, resultImage)
	if err != nil {
		panic(err)
	}
}

func saveImageWithUniqueName(directoryPath string, resultImage image.Image) {
	fileName := "IMG_1.jpg"
	resultImagePath := directoryPath + "\\" + fileName

	// Check if the file already exists
	_, err := os.Stat(resultImagePath)
	if err == nil {
		// File with the same name exists, find a unique name
		i := 2
		for {
			newFileName := fmt.Sprintf("IMG_%d.jpg", i)
			newResultImagePath := directoryPath + "\\" + newFileName

			_, err := os.Stat(newResultImagePath)
			if err != nil {
				// File does not exist, use this filename
				fileName = newFileName
				resultImagePath = newResultImagePath
				break
			}

			i++
		}
	}

	err = saveImage(resultImagePath, resultImage)
	if err != nil {
		panic(err)
	}
}
func loadImage(path string) (image.Image, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}

	return img, nil
}

func createWhiteBackground(width, height int) *image.RGBA {
	white := color.RGBA{R: 255, G: 255, B: 255, A: 255}
	if width < height {
		width = height
	} else {
		height = width
	}
	backgroundImage := image.NewRGBA(image.Rect(0, 0, width, height))
	draw.Draw(backgroundImage, backgroundImage.Bounds(), &image.Uniform{white}, image.Point{}, draw.Src)
	return backgroundImage
}

func overlayImages(background, overlay image.Image) image.Image {
	result := image.NewRGBA(background.Bounds())
	draw.Draw(result, background.Bounds(), background, image.Point{}, draw.Over)

	// Calculate the center point to draw the overlay image
	overlayStartX := (background.Bounds().Max.X - overlay.Bounds().Max.X) / 2
	overlayStartY := (background.Bounds().Max.Y - overlay.Bounds().Max.Y) / 2

	// Draw the overlay image at the center
	draw.Draw(result, overlay.Bounds().Add(image.Point{X: overlayStartX, Y: overlayStartY}), overlay, image.Point{}, draw.Over)

	return result
}

func saveImage(path string, img image.Image) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	return jpeg.Encode(file, img, &jpeg.Options{Quality: 100})
}
