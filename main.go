package main

import (
	"flag"
	"log"
	"time"

	"larana.tech/go/electrostatic/app"
	"larana.tech/go/electrostatic/config"
	"larana.tech/go/electrostatic/content"
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
		InitRoot(*root)
		return
	}

	cfg, err := config.Read(*root)

	if err != nil {
		log.Fatal(err)
	}

	if *hotreloadEnabled {
		cfg.Hotreload = *hotreloadEnabled
	}

	a := app.New(*root, *port, *dist, cfg)

	switch *mode {
	case "serve":
		a.Serve()
		return
	case "export":
		a.ExportSite()
		return
	}

	log.Fatal("Invalid mode: ", *mode)
}

func InitRoot(root string) {
	t0 := time.Now()

	log.Println("Init", root)

	content.InitializeTemplate(root)
	log.Println("Done in", time.Since(t0))
}
