package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/laranatech/electrostatic/pkg/content"
	"github.com/laranatech/electrostatic/pkg/export"
	"github.com/laranatech/electrostatic/pkg/pages"
	"github.com/laranatech/electrostatic/pkg/sitemap"
	"github.com/laranatech/electrostatic/pkg/static"
)

func main() {
	mode := flag.String("m", "export", "electrostatic mode. Can be `init`, `serve` or `export`(default)")
	root := flag.String("r", "", "content directory path")
	port := flag.String("p", ":3030", "Serving port for SSMG (SSR) mode")
	dist := flag.String("d", "./dist", "Output directory")
	flag.Parse()

	if *root == "" {
		fmt.Println("You have to specify the content path with `-r` flag ")
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

	fmt.Println("Invalid mode:", mode)
}

func serve(root, port string) {
	fmt.Println("Serving on", port)

	sitemap.ServeSitemap()
	static.ServeStatic(root)
	pages.ServePages(root)

	err := http.ListenAndServe(port, nil)

	if err != nil {
		log.Fatal(err.Error())
	}
}

func exportSite(root, dist string) {
	t0 := time.Now()

	fmt.Println("exporting")

	err := export.Export(root, dist)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("done in", time.Since(t0))
}

func initRoot(root string) {
	t0 := time.Now()

	fmt.Println("Init", root)

	content.InitializeContentTemplate(root)
	fmt.Println("Done in", time.Since(t0))
}
