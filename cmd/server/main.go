package main

import (
	"html/template"
	"io"
	"log"
	"net/http"

	"github.com/labstack/echo"
)

type TemplateRenderer struct {
	templates *template.Template
}

func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

type Todo struct {
	Name string
}

type Content struct {
	Todos []Todo
}

func main() {
	r := echo.New()

	tmpl, err := template.ParseGlob("./resources/views/*.html")
	if err != nil {
		log.Fatalf("unable to parse glob %e\n", err)
	}

	r.Renderer = &TemplateRenderer{
		templates: tmpl,
	}

	todos := []Todo{
		{
			Name: "test 1",
		},
		{
			Name: "test 2",
		},
		{
			Name: "test 3",
		},
	}
	data := Content{
		Todos: todos,
	}

	r.GET("/", func(c echo.Context) error {
		return c.Render(http.StatusOK, "index.html", data)
	})
	r.Static("/css", "resources/css");

	r.Logger.Fatal(r.Start(":1234"))
}
