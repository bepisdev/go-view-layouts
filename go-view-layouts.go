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

func Init(template_files map[string]string, layout_file string) map[string]*template.Template {
	templates = make(map[string]*template.Template)
	for _, file := range template_files {
    		templates[template_files[0]] = template.Must(template.ParseFiles(template_files[1], layout_file))
  	}
	return templates
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
