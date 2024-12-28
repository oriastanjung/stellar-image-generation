package main

import (
	"fmt"
	"log"

	"github.com/oriastanjung/stellar/internal/usecase"
)

func main() {
	// Example of how you might build or receive a user prompt at runtime.
	// In this example, we’re just hard-coding the details for demonstration.
	coreSubject := "A majestic flying castle"
	keyDescriptors := "crystalline spires, glowing stained-glass windows"
	environment := "high above the clouds"
	style := "surreal fantasy style, soft pastel color palette"
	moodTone := "epic and uplifting"
	composition := "wide shot, symmetrical"
	additionalInstructions := "illuminated by bright golden sunlight"

	// Build the final prompt string
	prompt := fmt.Sprintf(
		"Core Subject: %s\nKey Descriptors: %s\nEnvironment: %s\nStyle: %s\nMood/Tone: %s\nComposition: %s\nAdditional Instructions: %s",
		coreSubject,
		keyDescriptors,
		environment,
		style,
		moodTone,
		composition,
		additionalInstructions,
	)

	// Call our usecase function to send the request
	imageURL, filename, err := usecase.GenerateImage(prompt)
	if err != nil {
		log.Fatalf("Failed to generate image: %v", err)
	}
	fmt.Println("Image URL:", imageURL)
	if err := usecase.DownloadAndSaveImages(imageURL, filename); err != nil {
		log.Fatalf("Failed to download and save image: %v", err)
	}
	fmt.Println("Response from API:")
	fmt.Println(imageURL)
	fmt.Println("Downloaded and saved image:", filename)
}
