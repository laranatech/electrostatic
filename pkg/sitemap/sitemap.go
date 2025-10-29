package sitemap

import (
	"fmt"
	"net/http"
)

const SitemapRoute = "/sitemap.xml"

func ServeSitemap() {
	http.HandleFunc(SitemapRoute, func(w http.ResponseWriter, r *http.Request) {
		// read content dir recursively
		// add all entries
		// markup xml
		fmt.Println("Sitemap requested")

		w.WriteHeader(404)
	})
}
