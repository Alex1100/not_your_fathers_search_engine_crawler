package subscriber

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"not_your_fathers_search_engine_crawler/config"
	"not_your_fathers_search_engine_crawler/pkg/services/crawler"
	"not_your_fathers_search_engine_crawler/pkg/services/publisher"
	"os"
	"sync"

	"cloud.google.com/go/pubsub"
	"google.golang.org/api/option"
)

// PullCrawlFromSourceMsgs receives messages, crawls and publishes
// list of crawled links found from original src
func PullCrawlFromSourceMsgs(w io.Writer, projectID, subID string) error {
	ctx := context.Background()
	client, err := pubsub.NewClient(ctx, projectID, option.WithCredentialsFile(os.Getenv("google_app_path")))
	if err != nil {
		return fmt.Errorf("pubsub.NewClient: %v", err)
	}

	// Consume 10 messages.
	var mu sync.Mutex
	received := 0
	sub := client.Subscription(subID)
	cctx, cancel := context.WithCancel(ctx)
	appConfig := config.ReadConfig()
	pubSubConfig := appConfig.PubSubConfig

	err = sub.Receive(cctx, func(ctx context.Context, msg *pubsub.Message) {
		fmt.Fprintf(w, "Got message: %q\n", string(msg.Data))

		links := crawler.StartCrawlProcess(string(msg.Data))
		pubBuf := new(bytes.Buffer)
		err := publisher.PublishLinks(pubBuf, pubSubConfig.ProjectID, pubSubConfig.Topics.UpsertLink, links)

		if err != nil {
			fmt.Println("Epic failure, you should probably look into it: ", err)
		}

		msg.Ack()
		mu.Lock()
		defer mu.Unlock()

		received++
		if received == 10 {
			cancel()
		}
	})
	if err != nil {
		return fmt.Errorf("Receive: %v", err)
	}
	return nil
}
