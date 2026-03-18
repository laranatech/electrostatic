package app

import (
	"log"
	"net/http"
	"time"

	"larana.tech/go/electrostatic/config"
)

type Args struct {
	Version   bool
	Help      bool
	Hotreload bool
	Archive   bool
	Root      string
	Mode      string
	Port      string
	Dist      string
}

type App struct {
	Root string
	Port string
	Dist string
	Args *Args
	Cfg  *config.Config
}

func New(
	args *Args,
	cfg *config.Config,
) *App {
	a := &App{
		Args: args,
		Root: args.Root,
		Port: args.Port,
		Dist: args.Dist,
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
