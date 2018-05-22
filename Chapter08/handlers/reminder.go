package handlers

import (
	"html/template"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/labstack/echo"
	uuid "github.com/satori/go.uuid"
)

const RemindersTmpl = `
<html>
        <head>
                <title>{{.Title}}</title>
        </head>
        <body>
			<a href={{ .Reverse "login" }}>Login</a>
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

type CustomTemplate struct {
	*template.Template
}

func (ct *CustomTemplate) Render(w io.Writer, name string, data interface{},
	ctx echo.Context) error {
	return ct.ExecuteTemplate(w, name, data)
}

func CreateReminder(ctx echo.Context) error {
	return nil
}

func GetReminder(ctx echo.Context) error {
	ctx.Logger().Info("Reminder id is: ", ctx.Param("id"))
	return nil
}

func RenderReminders(c echo.Context) error {
	reminders := []Reminder{
		Reminder{
			ID:   uuid.NewV4(),
			Name: "Oil Change",
			Due:  time.Now().Add(30 * 3 * 24 * time.Hour),
		},
		Reminder{
			ID:   uuid.NewV4(),
			Name: "Birthday Party",
			Due:  mustTime(time.Parse("2006-01-02", "2020-01-01")),
		},
	}

	tmplData := struct {
		Reminders []Reminder
		Title     string
	}{reminders, "Reminders Page"}

	return c.Render(http.StatusOK, "reminders", tmplData)
}

type TmplData struct {
	Reminders []Reminder
	Title     string
	rev       func(name string, params ...interface{}) string
}

func (td TmplData) Reverse(name string, params ...interface{}) string {
	return td.rev(name, params...)
}

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

func RenderMoreReminders(c echo.Context) error {
	reminders := []Reminder{
		Reminder{
			ID:   uuid.NewV4(),
			Name: "Oil Change",
			Due:  time.Now().Add(30 * 3 * 24 * time.Hour),
		},
		Reminder{
			ID:   uuid.NewV4(),
			Name: "Birthday Party",
			Due:  mustTime(time.Parse("2006-01-02", "2020-01-01")),
		},
	}
	data := TmplData{reminders, "Reminders Page", c.Echo().Reverse}
	return c.Render(http.StatusOK, "reminders", data)
}
