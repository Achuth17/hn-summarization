package hn

import (
	"context"
	"fmt"
	"hn-exploration/fetcher"
	"strings"

	"github.com/alexferrari88/gohn/pkg/gohn"
	"github.com/alexferrari88/gohn/pkg/processors"
)

func GetEncodedStoryTopComments(id int) string {
	hn, err := gohn.NewClient(nil)
	if err != nil {
		panic(err)
	}
	ctx := context.Background()
	var story *gohn.Item
	story, _ = hn.Items.Get(ctx, id)
	if story == nil {
		panic("No story found")
	}
	commentsMap, err := hn.Items.FetchAllDescendants(ctx, story, processors.UnescapeHTML())
	if err != nil {
		panic(err)
	}
	storyWithComments := gohn.Story{
		Parent:          story,
		CommentsByIdMap: commentsMap,
	}
	storyWithComments.SetCommentsPosition()
	orderedIDs, err := storyWithComments.GetOrderedCommentsIDs()
	if err != nil {
		panic(err)
	}
	return convertTopCommentsIntoText(orderedIDs, commentsMap, storyWithComments)
}

func GetEncodedStoryContent(id int) string {
	ctx := context.Background()
	hn, err := gohn.NewClient(nil)
	if err != nil {
		panic(err)
	}
	var story *gohn.Item
	story, _ = hn.Items.Get(ctx, id)
	if story == nil {
		panic("No story found")
	}
	postStr, _ := fetcher.FetchURLContent(*story.URL)
	return postStr
}

func convertTopCommentsIntoText(orderedIDs []int, commentsMap map[int]*gohn.Item, storyWithComments gohn.Story) string {
	topLevelComments := []*gohn.Item{}
	var sb strings.Builder
	for _, id := range orderedIDs {
		comment := commentsMap[id]
		isTopLevel, _ := storyWithComments.IsTopLevelComment(comment)
		if isTopLevel && comment.Text != nil {
			topLevelComments = append(topLevelComments, comment)
		}
	}
	for _, comment := range topLevelComments {
		sb.WriteString(fmt.Sprintf("%s: %s\n\n", *comment.By, *comment.Text))
	}
	return sb.String()
}
