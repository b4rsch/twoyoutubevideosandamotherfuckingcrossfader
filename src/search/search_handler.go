package search

import (
	"flag"
	"fmt"
	"google.golang.org/api/googleapi/transport"
	"google.golang.org/api/youtube/v3"
	"log"
	"net/http"
)

var maxResults = flag.Int64("max-results", 25, "Max YouTube results")

const developerKey = ""

type Search struct {
	Pattern string
}

func SearchHandler(term string) map[string]string {
	query := flag.String("query", term, "Search term")
	flag.Parse()

	client := &http.Client{
		Transport: &transport.APIKey{Key: developerKey},
	}

	service, err := youtube.New(client)
	if err != nil {
		log.Fatalf("Error creating new YouTube client: %v", err)
	}

	// Make the API call to YouTube.
	call := service.Search.List([]string{"id,snippet"}).
		Q(*query).
		MaxResults(*maxResults).
		VideoEmbeddable("true")
	response, err := call.Do()
	// Group video, channel, and playlist results in separate lists.
	videos := make(map[string]string)
	if err != nil {
		fmt.Println("error occurred", err)
		return videos
	}

	// Iterate through each item and add it to the correct list.
	for _, item := range response.Items {
		if item.Id.Kind == "youtube#video" {
			videos[item.Id.VideoId] = item.Snippet.Title
		}
	}

	return videos
}

// Print the ID and title of each result in a list as well as a name that
// identifies the list. For example, print the word section name "Videos"
// above a list of video search results, followed by the video ID and title
// of each matching video.
//func printIDs(sectionName string, matches map[string]string) {
//	fmt.Printf("%v:\n", sectionName)
//	for id, title := range matches {
//		fmt.Printf("[%v] %v\n", id, title)
//	}
//	fmt.Printf("\n\n")
//}
