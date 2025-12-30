package pages

import (
	"net/http"
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
