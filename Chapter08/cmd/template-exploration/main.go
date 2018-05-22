package main

import (
	"html/template"
	"log"
	"os"
	"time"

	uuid "github.com/satori/go.uuid"
)

const tmpl = `
<html>
	<head>
		<title>{{.Title}}</title>
	</head>
	<body>
		<table>
			<th>
				<td>Reminder Name</td>
				<td>Reminder ID</td>
				<td>Reminder Due</td>
			</th>
			{{ range .Reminders }}
			<tr>
				<td>{{ .Name }}</td>
				<td>{{ .ID }}</td>
				<td>{{ .Due }}</td>
			</tr>
			{{ else }}
			<tr>
				<td colspan=3>No rows!</td>
			</tr>
			{{ end }}
	</body>
</html>
`

type Reminder struct {
	ID   uuid.UUID
	Name string
	Due  time.Time
}

const Week = time.Duration(time.Hour * 24 * 7)

func mustTime(d time.Time, err error) time.Time {
	if err != nil {
		log.Fatalf("failed must condition: %s\n", err.Error())
	}
	return d
}

func main() {
	reminders := []Reminder{
		Reminder{
			ID:   uuid.NewV4(),
			Name: "Oil Change",
			Due:  time.Now().Add(13 * Week),
		},
		Reminder{
			ID:   uuid.NewV4(),
			Name: "Birthday Party",
			Due:  mustTime(time.Parse("2006-01-02", "2020-01-01")),
		},
	}

	t, err := template.New("reminders").Parse(tmpl)
	if err != nil {
		log.Fatalf("failed to parse template: %s\n", err.Error())
	}

	tmplData := struct {
		Reminders []Reminder
		Title     string
	}{reminders, "Reminders Page"}

	// render the template output based on dynamic data
	err = t.Execute(os.Stdout, tmplData)
	if err != nil {
		log.Fatalf("failed to render template: %s\n", err.Error())
	}
}
