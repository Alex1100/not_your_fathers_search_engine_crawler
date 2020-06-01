package publisher

import (
	"context"
	"fmt"
	"io"
	"os"

	"cloud.google.com/go/pubsub"
	"google.golang.org/api/option"
)

// PublishLinks sends a message to add to a topic on
// Google Cloud Pub/Sub
func PublishLinks(w io.Writer, projectID, topicID string, payload []byte) error {
	ctx := context.Background()
	client, err := pubsub.NewClient(ctx, projectID, option.WithCredentialsFile(os.Getenv("google_app_path")))
	if err != nil {
		return fmt.Errorf("pubsub.NewClient: %v", err)
	}

	t := client.Topic(topicID)
	result := t.Publish(ctx, &pubsub.Message{
		Data: payload,
	})

	// Block until the result is returned and a server-generated
	// ID is returned for the published message.
	id, err := result.Get(ctx)
	if err != nil {
		return fmt.Errorf("Get: %v", err)
	}
	fmt.Fprintf(w, "Published a message; msg ID: %v\n", id)
	return nil
}
