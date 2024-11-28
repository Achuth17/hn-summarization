package gemini

import (
	"context"
	"encoding/json"
	"log"
	"os"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

func InvokeGemini(prompt string) string {
	ctx := context.Background()
	// Access your API key as an environment variable (see "Set up your API key" above)
	client, err := genai.NewClient(ctx, option.WithAPIKey(os.Getenv("API_KEY")))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()
	model := client.GenerativeModel("gemini-1.5-flash")
	resp, err := model.GenerateContent(ctx, genai.Text(prompt))
	if err != nil {
		log.Fatal(err)
	}
	return getResponseAsString(resp)
}

func getResponseAsString(resp *genai.GenerateContentResponse) string {
	respStr := ""
	for _, cand := range resp.Candidates {
		if cand.Content != nil {
			for _, part := range cand.Content.Parts {
				text, err := parseString(part.(genai.Text))
				if err != nil {
					log.Fatal(err)
				}
				respStr += text
			}
		}
	}
	return respStr
}

func parseString(text genai.Text) (string, error) {
	content, _ := json.Marshal(text)

	var responseContentString string
	err := json.Unmarshal(content, &responseContentString)
	if err != nil {
		return "", err
	}

	return responseContentString, nil
}
