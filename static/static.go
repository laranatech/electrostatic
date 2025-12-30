package static

import (
	"net/http"
)

func ServeStatic(root string) {
	fsAssets := http.FileServer(http.Dir(root + "/assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fsAssets))

	publicAssets := http.FileServer(http.Dir(root + "/public"))
	http.Handle("/public/", http.StripPrefix("/public/", publicAssets))
}
