package app

import (
	"fmt"
	"html/template"
	"net/http"
	"path"

	"larana.tech/go/electrostatic/app/hotreload"
	"larana.tech/go/electrostatic/templates"
)

func (a *App) RespondWithError(
	w http.ResponseWriter,
	r *http.Request,
	errorCode int,
) error {
	w.WriteHeader(errorCode)

	filepath := path.Join(a.Root, fmt.Sprintf("/%v.md", errorCode))

	page, err := ReadPageFile(a.Root, filepath, a.Cfg)

	if err != nil {
		return err
	}

	tmpltPath := path.Join(a.Root, a.Cfg.DefaultTemplate)

	tmp, err := templates.Read(tmpltPath)

	if err != nil {
		return err
	}

	t := template.Must(template.New("page").Parse(tmp))

	result, err := templates.FormatPage(t, page, a.Cfg)

	if a.Cfg.Hotreload {
		result = hotreload.Inject(result)
	}

	w.Header().Add("Content-Type", "text/html")
	w.Write([]byte(result))

	return nil
}
