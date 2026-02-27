package pages

import (
	"errors"
	"log"
	"net/http"
	"os"
	"path"
	"strings"
)

func ServePages(root string) {
	http.HandleFunc("/articles", func(w http.ResponseWriter, r *http.Request) {
		result, err := FormatPageList(root)

		if err != nil {
			w.WriteHeader(500)
			return
		}

		w.WriteHeader(200)
		w.Header().Add("Content-Type", "text/html")
		w.Write([]byte(result))
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		paths, err := ScanAllFilepaths(root)

		if err != nil {
			w.WriteHeader(500)
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
			w.WriteHeader(404)
			return
		}

		page, err := ReadPageFile(root, filepath)

		if err != nil {
			w.WriteHeader(500)
			return
		}

		tmp, err := ReadTemplateFile(root)

		if err != nil {
			w.WriteHeader(500)
			return
		}

		result := FormatTemplate(tmp, page)

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
