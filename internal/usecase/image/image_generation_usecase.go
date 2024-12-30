package usecase

import (
	"bytes"
	"encoding/json"
	"fmt"
	"image"
	"image/jpeg"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/chai2010/webp"
	"github.com/oriastanjung/stellar/internal/config"
)

// ImageUseCase defines the contract for your core logic.
type ImageUseCase interface {
	GenerateImage(prompt string) (string, string, error)
	DownloadAndSaveImages(imageURL, filename string) error
}

// imageUseCase is the concrete struct implementing the ImageUseCase interface.
type imageUseCase struct{}

// NewImageUseCase creates a new instance of imageUseCase.
func NewImageUseCase() ImageUseCase {
	return &imageUseCase{}
}

// Message represents a single message in the request body.
type Message struct {
	ID      string `json:"id"`
	Content string `json:"content"`
	Role    string `json:"role"`
}

// AgentMode represents the agent mode data structure.
type AgentMode struct {
	Mode bool   `json:"mode"`
	ID   string `json:"id"`
	Name string `json:"name"`
}

// RequestBody represents the full JSON payload for the external service.
type RequestBody struct {
	Messages              []Message              `json:"messages"`
	ID                    string                 `json:"id"`
	PreviewToken          interface{}            `json:"previewToken"`
	UserID                interface{}            `json:"userId"`
	CodeModelMode         bool                   `json:"codeModelMode"`
	AgentMode             AgentMode              `json:"agentMode"`
	TrendingAgentMode     map[string]interface{} `json:"trendingAgentMode"`
	IsMicMode             bool                   `json:"isMicMode"`
	UserSystemPrompt      interface{}            `json:"userSystemPrompt"`
	MaxTokens             int                    `json:"maxTokens"`
	PlaygroundTopP        interface{}            `json:"playgroundTopP"`
	PlaygroundTemp        interface{}            `json:"playgroundTemperature"`
	IsChromeExt           bool                   `json:"isChromeExt"`
	GithubToken           string                 `json:"githubToken"`
	ClickedAnswer2        bool                   `json:"clickedAnswer2"`
	ClickedAnswer3        bool                   `json:"clickedAnswer3"`
	ClickedForceWebSearch bool                   `json:"clickedForceWebSearch"`
	VisitFromDelta        bool                   `json:"visitFromDelta"`
	MobileClient          bool                   `json:"mobileClient"`
	UserSelectedModel     interface{}            `json:"userSelectedModel"`
	Validated             string                 `json:"validated"`
	ImageGenerationMode   bool                   `json:"imageGenerationMode"`
	WebSearchModePrompt   bool                   `json:"webSearchModePrompt"`
	DeepSearchMode        bool                   `json:"deepSearchMode"`
	Domains               interface{}            `json:"domains"`
}

// GenerateImage calls an external API to produce an image URL/filename based on the prompt.
func (uc *imageUseCase) GenerateImage(prompt string) (string, string, error) {
	cfg := config.LoadEnv()

	// Build the request body
	requestPayload := RequestBody{
		Messages: []Message{
			{
				ID:      cfg.IMAGE_API_GENERATION_MESSAGE_ID,
				Content: prompt,
				Role:    "user",
			},
		},
		ID:                  cfg.IMAGE_API_GENERATION_MESSAGE_ID,
		CodeModelMode:       true,
		AgentMode:           AgentMode{Mode: true, ID: cfg.IMAGE_GENERATION_MODEL, Name: "Image Generation"},
		TrendingAgentMode:   map[string]interface{}{},
		IsMicMode:           false,
		MaxTokens:           1024,
		Validated:           cfg.IMAGE_API_GENERATION_VALIDATED,
		ImageGenerationMode: true,
	}

	// Serialize the body to JSON
	jsonData, err := json.Marshal(requestPayload)
	if err != nil {
		return "", "", fmt.Errorf("error marshaling JSON: %w", err)
	}

	// Create a new HTTP request
	req, err := http.NewRequest("POST", cfg.IMAGE_API_GENERATION_URL, bytes.NewBuffer(jsonData))
	if err != nil {
		return "", "", fmt.Errorf("error creating request: %w", err)
	}

	// Set up headers
	req.Header.Set("Accept", "*/*")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("origin", cfg.IMAGE_API_GENERATION_ORIGIN)
	referrer := fmt.Sprintf("%s/agent/%s", cfg.IMAGE_API_GENERATION_ORIGIN, cfg.IMAGE_GENERATION_MODEL)
	req.Header.Set("referer", referrer)

	// Send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", "", fmt.Errorf("error sending request: %w", err)
	}
	defer resp.Body.Close()

	// Read response
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", "", fmt.Errorf("error reading response: %w", err)
	}

	// Extract the image URL and filename
	imageURL, filename, err := uc.extractImageURLAndFilename(string(respBody))
	if err != nil {
		return "", "", err
	}

	return imageURL, filename, nil
}

// extractImageURLAndFilename is a private helper function to parse the
// response body and retrieve the image URL + filename from markdown syntax (![](...)).
func (uc *imageUseCase) extractImageURLAndFilename(respBody string) (string, string, error) {
	// Step 1: Find the start of the image URL using "![]("
	startIndex := strings.Index(respBody, "![](")
	if startIndex == -1 {
		return "", "", fmt.Errorf("image URL not found in response body")
	}
	startIndex += len("![](")

	// Step 2: Find the closing parenthesis ')'
	endIndex := strings.Index(respBody[startIndex:], ")")
	if endIndex == -1 {
		return "", "", fmt.Errorf("closing parenthesis for image URL not found")
	}
	endIndex += startIndex

	// Step 3: Extract the URL
	imageURL := respBody[startIndex:endIndex]

	// Step 4: Extract the filename
	lastSlashIndex := strings.LastIndex(imageURL, "/")
	if lastSlashIndex == -1 || !strings.HasSuffix(imageURL, ".jpeg") {
		return "", "", fmt.Errorf("invalid image URL format: %s", imageURL)
	}
	filenameWithExt := imageURL[lastSlashIndex+1:] // e.g. "example123.jpeg"
	filename := strings.TrimSuffix(filenameWithExt, ".jpeg")

	return imageURL, filename, nil
}

func (uc *imageUseCase) DownloadAndSaveImages(imageURL, baseFileName string) error {
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
