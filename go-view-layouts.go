package layouts

import (
	"html/template"
	"net/http"
	"sync"
)

var (
	templates     map[string]*template.Template
	templatesLock sync.Mutex
)

// initializes the templates map by parsing the provided template files with the layout file.
// Returns an error if any template fails to parse.
func Init(templateFiles map[string]string, layoutFile string) error {
	templatesLock.Lock()
	defer templatesLock.Unlock()

	templates = make(map[string]*template.Template)
	for name, file := range templateFiles {
		tmpl, err := template.ParseFiles(file, layoutFile)
		if err != nil {
			return err
		}
		templates[name] = tmpl
	}
	return nil
}

// renders the specified template with the given data and writes the output to the http.ResponseWriter.
func RenderTemplate(w http.ResponseWriter, tmplName string, layout string, data any) {
	templatesLock.Lock()
	defer templatesLock.Unlock()

	tmpl, ok := templates[tmplName]
	if !ok {
		http.Error(w, "Template not found", http.StatusInternalServerError)
		return
	}

	err := tmpl.ExecuteTemplate(w, layout, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
