package eventstream

import (
	"bufio"
	"context"
	"net/http"
	"time"
)

func subscribe(ctx context.Context, client *http.Client, url string, since time.Time, handler func(evt *Event)) error {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url+"?since="+since.UTC().Format(time.RFC3339), nil)

	if err != nil {
		return err
	}

	req.Header.Set("Cache-Control", "no-cache")
	req.Header.Set("Accept", "text/event-stream")
	req.Header.Set("Connection", "keep-alive")
	res, err := client.Do(req)

	if err != nil {
		return err
	}

	defer res.Body.Close()
	reader := bufio.NewReader(res.Body)

	evt := new(Event)

	for {
		line, err := reader.ReadBytes('\n')

		if err != nil {
			return err
		}

		if len(line) <= 1 {
			continue
		}

		body := string(line)
		err = evt.SetID(body)

		if err != nil {
			err = evt.SetData(body)
		}

		if len(evt.ID) > 0 && len(evt.Data) > 0 && err == nil {
			handler(evt)
			evt = new(Event)
		}
	}
}
