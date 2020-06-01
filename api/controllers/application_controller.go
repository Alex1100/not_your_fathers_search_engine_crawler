package controllers

import (
	"net/http"

	links "not_your_fathers_search_engine_crawler/api/controllers/links"
)

// GetCrawlFromSource calls links.CrawlFromSource
func GetCrawlFromSource(w http.ResponseWriter, r *http.Request) {
	links.CrawlFromSource(w, r)
}
