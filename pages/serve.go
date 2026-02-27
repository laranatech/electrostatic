package pages

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"path"
	"strings"

	"larana.tech/go/electrostatic/config"
	"larana.tech/go/electrostatic/pages/hotreload"
)

func ServePages(root string, cfg *config.Config, hotreloadEnabled bool) {
	if hotreloadEnabled {
		wsHandler, err := hotreload.GetWSHandler(root)

		if err != nil {
			log.Fatal(err)
			return
		}

		http.HandleFunc("/ws", wsHandler)
	}

	for _, entry := range cfg.Catalogs.Entries {
		http.HandleFunc(entry.Path, func(w http.ResponseWriter, r *http.Request) {
			result, err := FormatPageList(root, &entry, cfg)

			if err != nil {
				log.Println(err)

				err = serveError(root, w, r, 500, hotreloadEnabled, cfg)

				if err != nil {
					log.Fatal(err)
				}
				return
			}

			if hotreloadEnabled {
				result = hotreload.Inject(result)
			}

			w.WriteHeader(200)
			w.Header().Add("Content-Type", "text/html")
			w.Write([]byte(result))
		})
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		paths, err := ScanAllFilepaths(root)

		if err != nil {
			log.Println(err)

			err = serveError(root, w, r, 500, hotreloadEnabled, cfg)

			if err != nil {
				log.Fatal(err)
			}

			return
		}

		var filepath string

		for _, v := range paths {
			if FormatFilepathToRoute(root, v) == r.URL.Path {
				filepath = v
				break
			}
		}

		if filepath == "" {
			err = serveStatic(root, w, r)
			if err == nil {
				return
			}

			log.Println(err)

			err = serveError(root, w, r, 404, hotreloadEnabled, cfg)

			if err != nil {
				log.Fatal(err)
			}

			return
		}

		page, err := ReadPageFile(root, filepath, cfg)

		if err != nil {
			log.Println(err)
			err = serveError(root, w, r, 500, hotreloadEnabled, cfg)

			if err != nil {
				log.Fatal(err)
			}

			return
		}

		tmpltPath := path.Join(root, cfg.DefaultTemplate)

		tmp, err := ReadTemplateFile(tmpltPath)

		if err != nil {
			log.Println(err)

			err = serveError(root, w, r, 500, hotreloadEnabled, cfg)

			if err != nil {
				log.Fatal(err)
			}

			return
		}

		result := FormatTemplate(tmp, page, cfg)

		if hotreloadEnabled {
			result = hotreload.Inject(result)
		}

		w.WriteHeader(200)
		w.Header().Add("Content-Type", "text/html")
		w.Write([]byte(result))
	})
}

func serveStatic(root string, w http.ResponseWriter, r *http.Request) error {
	cleanPath := path.Clean(r.URL.Path)

	if strings.Contains(cleanPath, "..") {
		return errors.New("cleanPath contains '..'")
	}

	filePath := path.Join(root, "/public", cleanPath)

	info, err := os.Stat(filePath)
	if err != nil {
		return err
	}

	if info.IsDir() {
		indexFile := path.Join(root, filePath, "index.html")

		if _, err := os.Stat(indexFile); errors.Is(err, os.ErrNotExist) {
			return err
		}

		http.ServeFile(w, r, indexFile)
		return nil
	}

	http.ServeFile(w, r, filePath)

	return nil
}

func serveError(
	root string,
	w http.ResponseWriter,
	r *http.Request,
	errorCode int,
	hotreloadEnabled bool,
	cfg *config.Config,
) error {
	w.WriteHeader(errorCode)

	filepath := path.Join(root, fmt.Sprintf("/%v.md", errorCode))

	page, err := ReadPageFile(root, filepath, cfg)

	if err != nil {
		return err
	}

	tmpltPath := path.Join(root, cfg.DefaultTemplate)

	tmp, err := ReadTemplateFile(tmpltPath)

	if err != nil {
		return err
	}

	result := FormatTemplate(tmp, page, cfg)

	if hotreloadEnabled {
		result = hotreload.Inject(result)
	}

	w.Header().Add("Content-Type", "text/html")
	w.Write([]byte(result))

	return nil
}
