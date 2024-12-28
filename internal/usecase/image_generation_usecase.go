package usecase

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/oriastanjung/stellar/internal/config"
)

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

// RequestBody represents the full JSON payload.
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

func GenerateImage(prompt string) (string, string, error) {
	cfg := config.LoadEnv()
	// Build the request body based on the Rust example.
	body := RequestBody{
		Messages: []Message{
			{
				ID:      cfg.IMAGE_API_GENERATION_MESSAGE_ID,
				Content: prompt, // This is the user prompt we pass in
				Role:    "user",
			},
		},
		ID:                    cfg.IMAGE_API_GENERATION_MESSAGE_ID,
		PreviewToken:          nil,
		UserID:                nil,
		CodeModelMode:         true,
		AgentMode:             AgentMode{Mode: true, ID: cfg.IMAGE_GENERATION_MODEL, Name: "Image Generation"},
		TrendingAgentMode:     map[string]interface{}{},
		IsMicMode:             false,
		UserSystemPrompt:      nil,
		MaxTokens:             1024,
		PlaygroundTopP:        nil,
		PlaygroundTemp:        nil,
		IsChromeExt:           false,
		GithubToken:           "",
		ClickedAnswer2:        false,
		ClickedAnswer3:        false,
		ClickedForceWebSearch: false,
		VisitFromDelta:        false,
		MobileClient:          false,
		UserSelectedModel:     nil,
		Validated:             cfg.IMAGE_API_GENERATION_VALIDATED,
		ImageGenerationMode:   true,
		WebSearchModePrompt:   false,
		DeepSearchMode:        false,
		Domains:               nil,
	}

	// Serialize the body to JSON
	jsonData, err := json.Marshal(body)
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
	req.Header.Set("accept-language", "en-US,en;q=0.9")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("origin", cfg.IMAGE_API_GENERATION_ORIGIN)
	req.Header.Set("priority", "u=1, i")
	referrer := fmt.Sprintf("%s/agent/%s", cfg.IMAGE_API_GENERATION_ORIGIN, cfg.IMAGE_GENERATION_MODEL)
	req.Header.Set("referer", referrer)

	// Create an HTTP client and send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", "", fmt.Errorf("error sending request: %w", err)
	}
	defer resp.Body.Close()

	// Read response
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", "", fmt.Errorf("error reading response: %w", err)
	}

	imageURL, filename, err := extractImageURLAndFilename(string(respBody))
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return "", "", err
	}

	return imageURL, filename, nil
}

func extractImageURLAndFilename(respBody string) (string, string, error) {
	// Step 1: Find the start of the image URL using "![]("
	startIndex := strings.Index(respBody, "![](")
	if startIndex == -1 {
		return "", "", fmt.Errorf("image URL not found in response body")
	}
	startIndex += len("![](") // Move index to the actual start of the URL

	// Step 2: Find the closing parenthesis ')'
	endIndex := strings.Index(respBody[startIndex:], ")")
	if endIndex == -1 {
		return "", "", fmt.Errorf("closing parenthesis for image URL not found")
	}
	endIndex += startIndex // Adjust endIndex relative to the full string

	// Step 3: Extract the URL
	imageURL := respBody[startIndex:endIndex]

	// Step 4: Extract the filename
	lastSlashIndex := strings.LastIndex(imageURL, "/")
	if lastSlashIndex == -1 || !strings.HasSuffix(imageURL, ".jpeg") {
		return "", "", fmt.Errorf("invalid image URL format")
	}
	filenameWithExt := imageURL[lastSlashIndex+1:] // Get the filename after the last '/'

	// Step 5: Remove the .jpeg extension
	filename := strings.TrimSuffix(filenameWithExt, ".jpeg")
	return imageURL, filename, nil
}
