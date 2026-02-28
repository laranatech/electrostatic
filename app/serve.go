package app

import (
	"errors"
	"log"
	"net/http"
	"os"
	"path"
	"strings"

	"larana.tech/go/electrostatic/app/hotreload"
	"larana.tech/go/electrostatic/templates"
)

func (a *App) ServePages() {
	if a.Cfg.Hotreload {
		wsHandler, err := hotreload.GetWSHandler(a.Root)

		if err != nil {
			log.Fatal(err)
			return
		}

		http.HandleFunc("/ws", wsHandler)
	}

	for _, entry := range a.Cfg.Catalogs.Entries {
		http.HandleFunc(entry.Path, func(w http.ResponseWriter, r *http.Request) {
			pages, err := a.ScanCatalogPages(&entry)

			if err != nil {
				log.Println(err)
				err = a.RespondWithError(w, r, 500)
				if err != nil {
					log.Fatal(err)
				}
				return
			}

			result, err := templates.FormatCatalogPage(a.Root, &entry, pages, a.Cfg)

			if err != nil {
				log.Println(err)

				err = a.RespondWithError(w, r, 500)

				if err != nil {
					log.Fatal(err)
				}
				return
			}

			if a.Cfg.Hotreload {
				result = hotreload.Inject(result)
			}

			w.WriteHeader(200)
			w.Header().Add("Content-Type", "text/html")
			w.Write([]byte(result))
		})
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		paths, err := ScanAllFilepaths(a.Root)

		if err != nil {
			log.Println(err)

			err = a.RespondWithError(w, r, 500)

			if err != nil {
				log.Fatal(err)
			}

			return
		}

		var filepath string

		for _, v := range paths {
			if FormatFilepathToRoute(a.Root, v) == r.URL.Path {
				filepath = v
				break
			}
		}

		if filepath == "" {
			err = a.ServeStatic(w, r)
			if err == nil {
				return
			}

			log.Println(err)

			err = a.RespondWithError(w, r, 404)

			if err != nil {
				log.Fatal(err)
			}

			return
		}

		page, err := ReadPageFile(a.Root, filepath, a.Cfg)

		if err != nil {
			log.Println(err)
			err = a.RespondWithError(w, r, 500)

			if err != nil {
				log.Fatal(err)
			}

			return
		}

		tmpltPath := path.Join(a.Root, a.Cfg.DefaultTemplate)

		tmp, err := templates.Read(tmpltPath)

		if err != nil {
			log.Println(err)

			err = a.RespondWithError(w, r, 500)

			if err != nil {
				log.Fatal(err)
			}

			return
		}

		result := templates.FormatTemplate(tmp, page, a.Cfg)

		if a.Cfg.Hotreload {
			result = hotreload.Inject(result)
		}

		w.WriteHeader(200)
		w.Header().Add("Content-Type", "text/html")
		w.Write([]byte(result))
	})
}

func (a *App) ServeStatic(w http.ResponseWriter, r *http.Request) error {
	cleanPath := path.Clean(r.URL.Path)

	if strings.Contains(cleanPath, "..") {
		return errors.New("cleanPath contains '..'")
	}

	filePath := path.Join(a.Root, "/public", cleanPath)

	info, err := os.Stat(filePath)

	if errors.Is(err, os.ErrNotExist) {
		info, err = os.Stat(filePath + ".html")

		if err == nil {
			http.ServeFile(w, r, filePath+".html")
			return nil
		}
	}

	if err != nil {
		return err
	}

	if info.IsDir() {
		indexFile := path.Join(a.Root, filePath, "index.html")

		if _, err := os.Stat(indexFile); errors.Is(err, os.ErrNotExist) {
			return err
		}

		http.ServeFile(w, r, indexFile)
		return nil
	}

	http.ServeFile(w, r, filePath)

	return nil
}
