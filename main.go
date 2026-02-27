package main

import (
	"flag"
	"log"
	"net/http"
	"time"

	"larana.tech/go/electrostatic/config"
	"larana.tech/go/electrostatic/content"
	"larana.tech/go/electrostatic/export"
	"larana.tech/go/electrostatic/pages"
	"larana.tech/go/electrostatic/sitemap"
)

func main() {
	mode := flag.String("m", "export", "electrostatic mode. Can be `init`, `serve` or `export`(default)")
	root := flag.String("r", "", "content directory path")
	port := flag.String("p", ":3030", "Serving port for SSMG (SSR) mode")
	dist := flag.String("d", "./dist", "Output directory")
	hotreloadEnabled := flag.Bool("h", false, "enable hot reload")
	flag.Parse()

	if *root == "" {
		log.Println("You have to specify the content path with `-r` flag ")
		return
	}

	if *mode == "init" {
		initRoot(*root)
		return
	}

	cfg, err := config.Read(*root)

	if err != nil {
		log.Fatal(err)
	}

	switch *mode {
	case "serve":
		serve(*root, *port, cfg, *hotreloadEnabled)
		return
	case "export":
		exportSite(*root, *dist, cfg)
		return
	}

	log.Fatal("Invalid mode: ", *mode)
}

func serve(root, port string, cfg *config.Config, hotreloadEnabled bool) {
	log.Println("Serving on", port)

	sitemap.ServeSitemap(cfg)
	pages.ServePages(root, cfg, hotreloadEnabled)

	err := http.ListenAndServe(port, nil)

	if err != nil {
		log.Fatal(err)
	}
}

func exportSite(root, dist string, cfg *config.Config) {
	t0 := time.Now()

	log.Println("exporting")

	err := export.Export(root, dist, cfg)

	if err != nil {
		log.Println(err)
		return
	}

	log.Println("done in", time.Since(t0))
}

func initRoot(root string) {
	t0 := time.Now()

	log.Println("Init", root)

	content.InitializeContentTemplate(root)
	log.Println("Done in", time.Since(t0))
}
