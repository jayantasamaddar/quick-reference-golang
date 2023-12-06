package std

import (
	"html/template"
	"os"
)

const tmpl = `
<!DOCTYPE html>
<html>
<head>
    <title>{{ index .Title }}</title>
</head>
<body>
    <h1>{{ .Heading }}</h1>
    <p>{{ .Content }}</p>
</body>
</html>
`

func HTMLTemplatesDemo() {

	// Create a new HTML template with the given name.
	t, err := template.New("example").Parse(tmpl)
	if err != nil {
		panic(err)
	}

	// Struct to insert data into the template
	data := struct {
		Title   string
		Heading string
		Content string
	}{
		Title:   "Sample Page",
		Heading: "Welcome to Go Templates",
		Content: "This is an example of using Go's HTML templates.",
	}

	// Use the Execute method of the parsed template to fill in the placeholders with the data and write the result to an output stream.
	// (e.g., an HTTP response or a file).
	// In this example, we're writing the rendered template to the standard output (os.Stdout), but you can use any io.Writer interface.
	// E.g.: An HTTP response writer or a file writer, depending on your use case.
	err = t.Execute(os.Stdout, data)
	if err != nil {
		panic(err)
	}

}
