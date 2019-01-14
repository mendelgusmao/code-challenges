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

type renderer struct {
	templateNames []string
}

func NewRenderer(templateNames ...string) renderer {
	return renderer{templateNames: templateNames}
}

func (rd renderer) Do(w http.ResponseWriter, r *http.Request, data interface{}) {
	t, ok := templates[rd.templateNames[0]]

	if !ok {
		layout, err := layoutTemplate.Clone()

		if err != nil {
			log.Printf("template.Render: %s", err)
			http.Error(w, "", http.StatusInternalServerError)
		}

		files := make([]string, len(rd.templateNames))

		for index, name := range rd.templateNames {
			files[index] = path.Join(config.Frontend.TemplatesDir, name+extension)
		}

		t, err = layout.
			Funcs(funcMap(w, r)).
			ParseFiles(files...)

		if err != nil {
			log.Printf("template.Render: %s", err)
			http.Error(w, "", http.StatusInternalServerError)
		}

		templates[rd.templateNames[0]] = t
	}

	if err := t.Funcs(funcMap(w, r)).ExecuteTemplate(w, "application", data); err != nil {
		log.Printf("template.Render: %s", err)
		http.Error(w, "", http.StatusInternalServerError)
	}
}
