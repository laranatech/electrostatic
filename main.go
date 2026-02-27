package main

import (
	"flag"
	"log"
	"net/http"
	"time"

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
	flag.Parse()

	if *root == "" {
		log.Println("You have to specify the content path with `-r` flag ")
		return
	}

	switch *mode {
	case "serve":
		serve(*root, *port)
		return
	case "export":
		exportSite(*root, *dist)
		return
	case "init":
		initRoot(*root)
	}

	log.Fatal("Invalid mode: ", *mode)
}

func serve(root, port string) {
	log.Println("Serving on", port)

	sitemap.ServeSitemap()
	pages.ServePages(root)

	err := http.ListenAndServe(port, nil)

	if err != nil {
		log.Fatal(err)
	}
}

func exportSite(root, dist string) {
	t0 := time.Now()

	log.Println("exporting")

	err := export.Export(root, dist)

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
