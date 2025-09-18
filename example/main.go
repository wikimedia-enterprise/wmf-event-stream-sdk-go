package main

import (
	"context"
	"log"
	"time"

	eventstream "github.com/wikimedia-enterprise/wmf-event-stream-sdk-go"
)

func main() {
	client := eventstream.NewClient()
	client.SetUserAgent("my-useragent contact@mymail.com")

	dt := time.Date(2014, 1, 1, 0, 0, 0, 0, time.UTC)
	stream := client.PageChange(context.Background(), dt, func(evt *eventstream.PageChange) error {
		log.Printf("page title: %s", evt.Data.Page.PageTitle)
		return nil
	})

	for err := range stream.Sub() {
		log.Println(err)
	}

}
