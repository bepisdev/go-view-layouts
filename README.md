# go-view-layouts

This Go library provides a thin wrapper around Go's `html/template` package for managing and rendering HTML templates with layouts. It ensures thread-safe access to templates, allowing concurrent use in a web server environment.

## Features
- Initialize templates with a common layout
- Render templates with thread-safe access
- Handle template parsing and rendering errors gracefully

## Installation

```bash
go get github.com/joshburnsxyz/go-view-layouts
```

## Usage

### Initialization

First, initialize the templates by providing a map of template names to file paths and the layout file path. This should be done once, typically at the start of your application.

```go
package main

import (
	"log"
	"net/http"
	layouts "github.com/joshburnsxyz/go-view-layouts"
)

func main() {
	templateFiles := map[string]string{
		"home":    "templates/home.html",
		"about":   "templates/about.html",
		"contact": "templates/contact.html",
	}

	err := layouts.Init(templateFiles, "templates/layout.html")
	if err != nil {
		log.Fatalf("Failed to initialize templates: %v", err)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		layouts.RenderTemplate(w, "home", "layout", nil)
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
```

### Rendering Templates

To render a template, call the `RenderTemplate` function with the response writer, template name, layout name, and data to be injected into the template.

```go
package main

import (
	"net/http"
	layouts "github.com/joshburnsxyz/go-view-layouts"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	data := struct {
		Title string
	}{
		Title: "Welcome Home",
	}
	layouts.RenderTemplate(w, "home", "layout", data)
}

func main() {
	// Initialize templates
	templateFiles := map[string]string{
		"home":    "templates/home.html",
		"about":   "templates/about.html",
		"contact": "templates/contact.html",
	}

	err := layouts.Init(templateFiles, "templates/layout.html")
	if err != nil {
		log.Fatalf("Failed to initialize templates: %v", err)
	}

	// Setup handlers
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
		layouts.RenderTemplate(w, "about", "layout", nil)
	})
	http.HandleFunc("/contact", func(w http.ResponseWriter, r *http.Request) {
		layouts.RenderTemplate(w, "contact", "layout", nil)
	})

	// Start server
	log.Fatal(http.ListenAndServe(":8080", nil))
}
```

## Contributing

If you find a bug or have a feature request, please open an issue on [GitHub](https://github.com/joshburnsxyz/go-view-layouts). Pull requests are welcome!
