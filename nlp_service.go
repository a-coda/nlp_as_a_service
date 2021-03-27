package main

import (
	"flag"
	"html/template"
	"log"
	"net/http"
	"github.com/jdkato/prose"
)

var addr = flag.String("addr", ":8080", "http service address") 

var templ = template.Must(template.New("nlp").Parse(templateStr))

func main() {
	flag.Parse()
	http.Handle("/", http.HandlerFunc(NLPService))
	err := http.ListenAndServe(*addr, nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}

func NLPService(w http.ResponseWriter, req *http.Request) {
	templ.Execute(w, analyzeThis(req.FormValue("text")))
}

func analyzeThis(text string) *prose.Document {
	doc, _ := prose.NewDocument(text)
	return doc
}

const templateStr = `
<html>
<head>
<title>NLP as a Service</title>
</head>
<body>
<h1>NLP as a Service</h1>
<h2>Text to analyze</h2>
<form action="/" name=nlp_form method="GET">
    <input maxLength=1024 size=80 name=text value="" title="Text to analyze">
    <input type=submit value="Show analysis" name=nlp>
</form>
<h2>Analysis</h2>
<h3>Entities</h3>
{{if .Entities}}
  {{range .Entities}} 
    {{.Text}} ({{.Label}})<br/>
  {{end}}
{{end}}
<h3>Sentences</h3>
{{if .Sentences}}
  {{range .Sentences}}
    {{.Text}}<br/>
  {{end}}
{{end}}
<h3>Tokens</h3>
{{if .Tokens}}
<table>
<tr>
  <th>Text</th>
  <th>POS Tag</th>
  <th>IOB</th>
</tr>
  {{range .Tokens}}
<tr style="text-align:center">
  <td>{{.Text}}</td>
  <td>{{.Tag}}</td>
  <td>{{.Label}}</td>
</tr>
  {{end}}
</table>
{{end}}
</body>
</html>
`
