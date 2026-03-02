package main

import (
	"flag"
	"fmt"
	"log"
	"time"

	"larana.tech/go/electrostatic/app"
	"larana.tech/go/electrostatic/config"
	"larana.tech/go/electrostatic/content"
)

type Args struct {
	Version   bool
	Help      bool
	Hotreload bool
	Root      string
	Mode      string
	Port      string
	Dist      string
}

func main() {
	args := parseArgs()

	if args.Version {
		fmt.Println("0.2.2")
		return
	}

	if args.Help {
		fmt.Println("Docs: https://larana.tech/electrostatic")
		return
	}

	if args.Root == "" {
		log.Println("You have to specify the content path with `-r` flag")
		return
	}

	if args.Mode == "init" {
		InitRoot(args.Root)
		return
	}

	cfg, err := config.Read(args.Root)

	if err != nil {
		log.Fatal(err)
	}

	if args.Hotreload {
		cfg.Hotreload = args.Hotreload
	}

	a := app.New(args.Root, args.Port, args.Dist, cfg)

	switch args.Mode {
	case "serve":
		a.Serve()
		return
	case "export":
		a.ExportSite()
		return
	}

	log.Fatal("Invalid mode: ", args.Mode)
}

func parseArgs() Args {
	mode := flag.String("m", "export", "electrostatic mode. Can be `init`, `serve` or `export`(default)")
	root := flag.String("r", "", "content directory path")
	port := flag.String("p", ":3030", "Serving port for SSMG (SSR) mode")
	dist := flag.String("d", "./dist", "Output directory")

	var v bool
	flag.BoolVar(&v, "v", false, "version")
	flag.BoolVar(&v, "version", false, "version")

	var h bool
	flag.BoolVar(&h, "h", false, "help")
	flag.BoolVar(&h, "help", false, "help")

	var hr bool
	flag.BoolVar(&hr, "hr", false, "enable hot reload")
	flag.BoolVar(&hr, "hot-reload", false, "enable hot reload")

	flag.Parse()

	return Args{
		Help:      h,
		Version:   v,
		Hotreload: hr,
		Root:      *root,
		Mode:      *mode,
		Port:      *port,
		Dist:      *dist,
	}
}

func InitRoot(root string) {
	t0 := time.Now()

	log.Println("Init", root)

	content.InitializeTemplate(root)
	log.Println("Done in", time.Since(t0))
}
