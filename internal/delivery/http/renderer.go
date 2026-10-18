package http

import (
	"html/template"
	"io"
	"path/filepath"
	"strings"

	"github.com/labstack/echo/v4"
)

// TemplateRenderer implements echo.Renderer interface using html/template
type TemplateRenderer struct {
	templates map[string]*template.Template
}

// NewTemplateRenderer compiles all templates and partials
func NewTemplateRenderer(rootDir string) (*TemplateRenderer, error) {
	templates := make(map[string]*template.Template)

	funcMap := template.FuncMap{
		"upper": strings.ToUpper,
		"lower": strings.ToLower,
		"dict": func(values ...interface{}) (map[string]interface{}, error) {
			m := make(map[string]interface{})
			for i := 0; i < len(values); i += 2 {
				key, ok := values[i].(string)
				if !ok {
					continue
				}
				if i+1 < len(values) {
					m[key] = values[i+1]
				}
			}
			return m, nil
		},
		"add": func(a, b int) int {
			return a + b
		},
		"formatStep": func(n int) string {
			if n < 10 {
				return "0" + string(rune('0'+n))
			}
			return string(rune('0' + n))
		},
	}

	layoutPath := filepath.Join(rootDir, "layouts", "base.html")
	partialsGlob := filepath.Join(rootDir, "partials", "*.html")
	pagesGlob := filepath.Join(rootDir, "pages", "*.html")

	pages, err := filepath.Glob(pagesGlob)
	if err != nil {
		return nil, err
	}

	for _, page := range pages {
		pageName := filepath.Base(page)
		files := []string{layoutPath, page}
		partialFiles, _ := filepath.Glob(partialsGlob)
		files = append(files, partialFiles...)

		tmpl, err := template.New(pageName).Funcs(funcMap).ParseFiles(files...)
		if err != nil {
			return nil, err
		}
		templates[pageName] = tmpl
	}

	return &TemplateRenderer{
		templates: templates,
	}, nil
}

// Render renders an HTML template
func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	tmpl, ok := t.templates[name]
	if !ok {
		return echo.ErrNotFound
	}
	// The layout defines "base", which calls "{{template "content" .}}"
	return tmpl.ExecuteTemplate(w, "base", data)
}

