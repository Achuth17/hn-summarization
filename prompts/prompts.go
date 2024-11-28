package prompts

import (
	"bufio"
	"os"
	"strings"
)

func GenerateSummarizationPrompt(post string, topComments string) string {
	promptTemplate := fetchPromptTemplate("prompts/transform-post-to-summary.txt")
	promptStr := strings.Replace(promptTemplate, "{@post-top-comments}", topComments, 1)
	promptStr = strings.Replace(promptStr, "{@post-content}", post, 1)
	return promptStr
}

func ConvertHTMLIntoTextPrompt(htmlContent string) string {
	promptTemplate := fetchPromptTemplate("prompts/transform-html-to-text.txt")
	promptStr := strings.Replace(promptTemplate, "{@post-html-content}", htmlContent, 1)
	return promptStr
}

func fetchPromptTemplate(promptFileName string) string {
	var sb strings.Builder
	file, err := os.Open(promptFileName)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err = file.Close(); err != nil {
			panic(err)
		}
	}()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			sb.WriteString("\n")
		} else {
			sb.WriteString(line)
		}
	}
	return sb.String()
}
