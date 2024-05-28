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

// FIXME: This is the original render function from Wikara. Modify this to 
// be a) more generic b) consume user defined templates.
// func RenderTemplate(w http.ResponseWriter, tmplname string, p *page.Page) {
// 	templatesLock.Lock()
// 	defer templatesLock.Unlock()

// 	tmpl, ok := templates[tmplname+".html"]
// 	if !ok {
// 		http.Error(w, "Template not found", http.StatusInternalServerError)
// 		return
// 	}

// 	err := tmpl.ExecuteTemplate(w, "base", p)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 	}
// }
