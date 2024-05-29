package layouts

import (
	"fmt"
	"html/template"
	"net/http"
	"sync"
)

var (
	// templates holds the parsed templates with their names as keys.
	templates map[string]*template.Template
	// templatesLock is a read-write mutex to synchronize access to the templates map.
	templatesLock sync.RWMutex
)

// Init initializes the templates map by parsing the provided template files along with the layout file.
// It takes a map of template file names (templateFiles) and a layout file name (layoutFile) as arguments.
// If any template fails to parse, it returns an error.
func Init(templateFiles map[string]string, layoutFile string) error {
	// Acquire a write lock to ensure exclusive access to the templates map.
	templatesLock.Lock()
	defer templatesLock.Unlock()

	// Initialize the templates map.
	templates = make(map[string]*template.Template)
	for name, file := range templateFiles {
		// Parse the template file along with the layout file.
		tmpl, err := template.ParseFiles(file, layoutFile)
		if err != nil {
			return fmt.Errorf("failed to parse template %s: %w", name, err)
		}
		// Store the parsed template in the map.
		templates[name] = tmpl
	}
	return nil
}

// RenderTemplate renders the specified template with the given data and writes the output to the http.ResponseWriter.
// It takes an http.ResponseWriter (w), the name of the template to render (tmplName), the name of the layout to use (layout),
// and the data to pass to the template (data) as arguments.
func RenderTemplate(w http.ResponseWriter, tmplName string, layout string, data any) {
	// Acquire a read lock to allow concurrent reads of the templates map.
	templatesLock.RLock()
	defer templatesLock.RUnlock()

	// Retrieve the template from the map.
	tmpl, ok := templates[tmplName]
	if !ok {
		http.Error(w, "Template not found", http.StatusInternalServerError)
		return
	}

	// Execute the template with the provided layout and data.
	err := tmpl.ExecuteTemplate(w, layout, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
