package hotreload

import (
	_ "embed"
	"fmt"
	"log"
	"net/http"
	"os"
	"path"
	"strings"
	"time"

	"github.com/gorilla/websocket"
)

//go:embed "hotreload.js"
var hotreloadScript string

func Inject(page string) string {
	return strings.Replace(
		page,
		"</body>",
		fmt.Sprintf("<script>%s</script></body>", hotreloadScript),
		1,
	)
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type file struct {
	path       string
	size       int64
	lastupdate time.Time
}

var ch = make(chan bool)

var files = make([]file, 0, 100)

func scan(root string) ([]file, error) {
	entries, err := os.ReadDir(root)

	if err != nil {
		return nil, err
	}

	tFiles := make([]file, 0, 100)

	for _, entry := range entries {
		filePath := path.Join(root, entry.Name())

		if entry.IsDir() {
			d, err := scan(filePath)

			if err != nil {
				return nil, err
			}

			tFiles = append(tFiles, d...)
			continue
		}

		info, err := entry.Info()

		if err != nil {
			return nil, err
		}

		f := file{
			path:       filePath,
			size:       info.Size(),
			lastupdate: info.ModTime(),
		}

		tFiles = append(tFiles, f)
	}

	return tFiles, nil
}

func compare(a, b []file) (bool, error) {
	if len(a) != len(b) {
		return true, nil
	}

	for i, fa := range a {
		fb := b[i]

		if fa.path != fb.path {
			return true, nil
		}

		if fa.lastupdate != fb.lastupdate {
			return true, nil
		}

		if fa.size != fb.size {
			return true, nil
		}
	}
	return false, nil
}

func watch(root string) chan bool {
	go func() {
		for {
			time.Sleep(2 * time.Second)
			tFiles, err := scan(root)

			if err != nil {
				log.Println(err)
				break
			}

			if len(files) == 0 {
				files = tFiles
				continue
			}

			res, err := compare(files, tFiles)

			files = tFiles

			if err != nil {
				log.Println(err)
				break
			}

			if res {
				ch <- res
			}
		}
	}()

	return ch
}

func GetWSHandler(root string) (http.HandlerFunc, error) {
	c := watch(root)

	return func(w http.ResponseWriter, r *http.Request) {
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Println(err)
			return
		}
		defer conn.Close()

		for v := range c {
			if !v {
				continue
			}

			if err := conn.WriteMessage(websocket.TextMessage, []byte("reload")); err != nil {
				break
			}

			conn.Close()
			break
		}
	}, nil
}
