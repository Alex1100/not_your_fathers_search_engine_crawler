package links

import (
	"bytes"
	"fmt"
	"net/http"

	config "not_your_fathers_search_engine_crawler/config"
	"not_your_fathers_search_engine_crawler/pkg/services/subscriber"
)

// CrawlFromSource crawls and publishes links
func CrawlFromSource(w http.ResponseWriter, r *http.Request) {
	subBuf := new(bytes.Buffer)
	appConfig := config.ReadConfig()
	pubSubConfig := appConfig.PubSubConfig

	errGettingSrc := subscriber.PullCrawlFromSourceMsgs(subBuf, pubSubConfig.ProjectID, pubSubConfig.Topics.UpsertLink)
	if errGettingSrc != nil {
		fmt.Println("Epic failure, you should probably look into it: ", errGettingSrc)
	}

	fmt.Println("LINKS PUBLISHED IN GOOGLE CLOUD PLATFORM'S PUB/SUB `upsert_link` TASK")
}
