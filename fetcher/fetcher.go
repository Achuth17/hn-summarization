package fetcher

import (
	"fmt"
	"hn-exploration/gemini"
	"hn-exploration/prompts"
	"io"
	"net/http"
)

func FetchURLContent(url string) (string, error) {
	htmlContent, err := fetchHTML(url)
	if err != nil {
		fmt.Println("Error fetching HTML:", err)
		return "", err
	}
	htmlExtractPrompt := prompts.ConvertHTMLIntoTextPrompt(htmlContent)
	articleContent := gemini.InvokeGemini(htmlExtractPrompt)
	// fmt.Println("-------")
	// fmt.Println(articleContent)
	// fmt.Println("-------")
	return articleContent, nil
}

func fetchHTML(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}
