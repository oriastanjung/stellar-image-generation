package usecase

import (
	"fmt"
	"image"
	"image/jpeg"
	"log"
	"net/http"
	"os"

	// We need this to properly decode PNG, GIF, etc., if the image isn't strictly JPEG:
	_ "image/gif"
	_ "image/png"

	// Third-party WebP library for encoding
	"github.com/chai2010/webp"
)

// DownloadAndSaveImages downloads the file from imageURL and saves both a .jpeg and .webp in ./public/.
func DownloadAndSaveImages(imageURL, baseFileName string) error {
	// 1. Get the file from the URL
	resp, err := http.Get(imageURL)
	if err != nil {
		return fmt.Errorf("failed to download image: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	// 2. Decode the image (any format supported by Go’s image package)
	img, _, err := image.Decode(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to decode image: %w", err)
	}

	// 3. Create public folder if it doesn’t exist
	if err := os.MkdirAll("../public", 0755); err != nil {
		return fmt.Errorf("failed to create public folder: %w", err)
	}

	// 4. Save as JPEG
	jpegPath := fmt.Sprintf("../public/%s.jpeg", baseFileName)
	outJpeg, err := os.Create(jpegPath)
	if err != nil {
		return fmt.Errorf("failed to create JPEG file: %w", err)
	}
	defer outJpeg.Close()

	// Encode image to JPEG
	if err = jpeg.Encode(outJpeg, img, nil); err != nil {
		return fmt.Errorf("failed to encode JPEG: %w", err)
	}
	log.Printf("Saved JPEG to %s\n", jpegPath)

	// 5. Save as WebP
	webpPath := fmt.Sprintf("../public/%s.webp", baseFileName)
	outWebp, err := os.Create(webpPath)
	if err != nil {
		return fmt.Errorf("failed to create WebP file: %w", err)
	}
	defer outWebp.Close()

	// Encode image to WebP (adjust Options for quality, lossless, etc.)
	// Quality can be 0-100, with 75-90 typically decent.
	if err = webp.Encode(outWebp, img, &webp.Options{Lossless: false, Quality: 80}); err != nil {
		return fmt.Errorf("failed to encode WebP: %w", err)
	}
	log.Printf("Saved WebP to %s\n", webpPath)

	return nil
}
