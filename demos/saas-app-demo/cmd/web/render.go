package main

// Render templates
import (
	"embed"
	"fmt"
	"html/template"
	"net/http"
	"strings"
)

// Data to pass to templates
type templateData struct {
	StringMap map[string]string
	IntMap    map[string]int
	FloatMap  map[string]float32
	Data      map[string]interface{}
	// For Cross-Site Request Forgery protection
	CSRFToken string
	// Message that displays once and then goes away
	Flash string
	// Display warning to the user
	Warning string
	// Error
	Error string
	// Whether user is authenticated
	IsAuthenticated int
	API             string
	CSSVersion      string
}

var functions = template.FuncMap{}

// Go's Embed functionality. Appeared in standard library in Go 1.16

//go:embed templates
var templateFS embed.FS

// Add Default template data
func (app *application) addDefaultData(td *templateData, r *http.Request) *templateData {
	td.API = app.config.api
	return td
}

// Render the template
func (app *application) renderTemplate(w http.ResponseWriter, r *http.Request, page string, td *templateData, partials ...string) error {
	var t *template.Template
	var err error
	templateToRender := fmt.Sprintf("templates/%s.page.gohtml", page)

	_, templateInMap := app.templateCache[templateToRender]

	// Use templateCache only in production
	if templateInMap && app.config.env == "production" {
		t = app.templateCache[templateToRender]
	} else {
		// If templateCache doesn't exist, map the partials, parse the template and add it to the templateCache map
		t, err = app.parseTemplate(partials, page, templateToRender)
		if err != nil {
			app.errorLog.Println(err)
			return err
		}
	}

	// Check if template data was passed. If not passed create a reference to template data. Add default data.
	if td == nil {
		td = &templateData{}
	}
	td = app.addDefaultData(td, r)

	// Execute applies a parsed template to the specified data object, writing the output to wr
	err = t.Execute(w, td)
	if err != nil {
		app.errorLog.Println(err)
		return err
	}

	return nil
}

// Parse the templates if templateCache doesn't exist and add to the `app.templateCache`
func (app *application) parseTemplate(partials []string, page string, templateToRender string) (*template.Template, error) {
	var t *template.Template
	var err error

	// build partials: map partials with the partial file names and
	if len(partials) > 0 {
		for i, x := range partials {
			partials[i] = fmt.Sprintf("templates/%s.partial.gohtml", x)
		}
	}

	if len(partials) > 0 {
		t, err = template.New(fmt.Sprintf("%s.page.gohtml", page)).Funcs(functions).ParseFS(templateFS, "templates/base.layout.gohtml", strings.Join(partials, ","), templateToRender)
	} else {
		// Case when there are no partials associated with templates
		t, err = template.New(fmt.Sprintf("%s.page.gohtml", page)).Funcs(functions).ParseFS(templateFS, "templates/base.layout.gohtml", templateToRender)
	}

	if err != nil {
		app.errorLog.Println(err)
		return nil, err
	}

	// Add to the templateCache
	app.templateCache[templateToRender] = t
	return t, nil
}
