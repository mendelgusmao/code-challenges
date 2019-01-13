package templates

import (
	"html/template"
	"log"
	"net/http"
	"path"

	"bitbucket.org/mendelgusmao/me_gu/frontend/config"
)

const extension = ".html"

var (
	templates      = make(map[string]*template.Template)
	layoutTemplate *template.Template
)

func init() {
	config.AfterLoad(loadLayoutTemplates)
}

func loadLayoutTemplates(c *config.Specification) error {
	tmpl, err := template.
		New("application").
		Funcs(funcMap(nil, nil)).
		ParseGlob(path.Join(c.TemplatesDir, "layout", "*"+extension))

	if err != nil {
		return err
	}

	layoutTemplate = tmpl

	return nil
}

func Render(w http.ResponseWriter, r *http.Request, name string, data interface{}) {
	t, ok := templates[name]

	if !ok {
		layout, err := layoutTemplate.Clone()

		if err != nil {
			log.Printf("template.Render: %s", err)
			http.Error(w, "", http.StatusInternalServerError)
		}

		t, err = layout.
			Funcs(funcMap(w, r)).
			ParseFiles(path.Join(config.Frontend.TemplatesDir, name+extension))

		if err != nil {
			log.Printf("template.Render: %s", err)
			http.Error(w, "", http.StatusInternalServerError)
		}

		templates[name] = t
	}

	if err := t.Funcs(funcMap(w, r)).ExecuteTemplate(w, "application", data); err != nil {
		log.Printf("template.Render: %s", err)
		http.Error(w, "", http.StatusInternalServerError)
	}
}
