package app

import (
	"log"
	"net/http"
	"time"

	"larana.tech/go/electrostatic/config"
)

type App struct {
	Root string
	Port string
	Dist string
	Cfg  *config.Config
}

func New(
	root, port, dist string,
	cfg *config.Config,
) *App {
	a := &App{
		Root: root,
		Port: port,
		Dist: dist,
		Cfg:  cfg,
	}

	return a
}

func (a *App) Serve() {
	log.Println("Serving on", a.Port)

	// TODO: issue-6
	// sitemap.ServeSitemap(cfg)
	a.ServePages()

	err := http.ListenAndServe(a.Port, nil)

	if err != nil {
		log.Fatal(err)
	}
}

func (a *App) ExportSite() {
	t0 := time.Now()

	log.Println("exporting")

	err := a.Export()

	if err != nil {
		log.Println(err)
		return
	}

	log.Println("done in", time.Since(t0))
}
