package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	eventstream "github.com/wikimedia-enterprise/wmf-event-stream-sdk-go"
)

func main() {
	client := eventstream.NewClient()

	stream := client.RevisionScore(context.Background(), time.Now().UTC(), func(evt *eventstream.RevisionScore) error {
		fmt.Println(evt.Data.Database)
		return errors.New("hello world")
	})

	for err := range stream.Sub() {
		log.Println(err)
	}
}
