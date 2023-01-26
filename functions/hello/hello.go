// Sends a simple message to the client
package hello

import (
	"html/template"
	"net/http"
	"time"

	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
)

func init() {
	functions.HTTP("HelloHTTP", HelloHTTP)
}

var htmlTemplate = `<html>
  <head>
    <title>Hello from Go</title>
	</head>
	<body>
		<p>Hello there!</p>
		<p>Here it is currently {{ .Date }}</p>
	</body>
</html>
	`

// HelloHTTP is an HTTP Cloud Function with a request parameter.
func HelloHTTP(w http.ResponseWriter, r *http.Request) {
	var my_template *template.Template = template.New("hello")
	my_template.Parse(htmlTemplate)
	dt := time.Now()
	render_data := struct {
		Date string
	}{
		Date: dt.Local().Format("02-01-2006 15:04:05"),
	}
	my_template.Execute(w, render_data)
}
