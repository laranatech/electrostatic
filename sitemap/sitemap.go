package sitemap

import (
	"log"
	"net/http"

	"larana.tech/go/electrostatic/config"
)

const SitemapRoute = "/sitemap.xml"

func ServeSitemap(cfg *config.Config) {
	http.HandleFunc(SitemapRoute, func(w http.ResponseWriter, r *http.Request) {
		// TODO: issue-6
		// read content dir recursively
		// add all entries
		// markup xml
		log.Println("Sitemap requested")

		w.WriteHeader(404)
	})
}
