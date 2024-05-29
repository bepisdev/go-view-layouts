package layouts

import (
	"fmt"
	"html/template"
	"net/http"
	"sync"
)

var (
	// parsed templates with their names as keys
	templates map[string]*template.Template
	// read-write mutex to synchronize access to templates
	templatesLock sync.RWMutex
)

// initializes the templates map by parsing the provided template files with the layout file.
// It returns an error if any template fails to parse.
func Init(templateFiles map[string]string, layoutFile string) error {
	// Acquire a write lock to ensure exclusive access to templates map
	templatesLock.Lock()
	defer templatesLock.Unlock()

	// Initialize the templates map
	templates = make(map[string]*template.Template)
	for name, file := range templateFiles {
		// Parse the template file along with the layout file
		tmpl, err := template.ParseFiles(file, layoutFile)
		if err != nil {
			return fmt.Errorf("failed to parse template %s: %w", name, err)
		}
		// Store the parsed template in the map
		templates[name] = tmpl
	}
	return nil
}

// renders the specified template with the given data and writes the output to the http.ResponseWriter.
func RenderTemplate(w http.ResponseWriter, tmplName string, layout string, data any) {
	// Acquire a read lock to allow concurrent reads of templates map
	templatesLock.RLock()
	defer templatesLock.RUnlock()

	// Retrieve the template from the map
	tmpl, ok := templates[tmplName]
	if !ok {
		http.Error(w, "Template not found", http.StatusInternalServerError)
		return
	}

	// Execute the template with the provided layout and data
	err := tmpl.ExecuteTemplate(w, layout, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
