package main

import (
	"context"
	"errors"
	"log"
	"time"

	eventstream "github.com/wikimedia-enterprise/wmf-event-stream-sdk-go"
)

func main() {
	client := eventstream.NewClient()

	dt := time.Date(2014, 1, 1, 0, 0, 0, 0, time.UTC)
	stream := client.PageChange(context.Background(), dt, func(evt *eventstream.PageChange) error {
		log.Printf("page title: %s", evt.Data.Page.PageTitle)
		return errors.New("hello world")
	})

	for err := range stream.Sub() {
		log.Println(err)
	}

}
