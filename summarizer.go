package main

import (
	"flag"
	"fmt"
	"hn-summarization/gemini"
	"hn-summarization/hn"
	"hn-summarization/prompts"
)

func main() {
	postId := flag.Int("postId", 41540984, "Enter a valid hackernews post id.")

	flag.Parse()

	encodedTopComments := hn.GetEncodedStoryTopComments(*postId)
	encodedPostContent := hn.GetEncodedStoryContent(*postId)
	summarizationPrompt := prompts.GenerateSummarizationPrompt(encodedPostContent, encodedTopComments)
	postSummary := gemini.InvokeGemini(summarizationPrompt)
	fmt.Println(postSummary)
}
