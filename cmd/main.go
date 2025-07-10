package main

import (
	"fmt"
	"html/template"
	"io"
	"time"

	"github.com/gofiber/fiber/v2"
)

// TemplateEngine wraps parsed templates
type TemplateEngine struct {
    templates *template.Template
}

// Render implements fiber's Views interface
func (t *TemplateEngine) Render(w io.Writer, name string, data interface{}, layouts ...string) error {
    return t.templates.ExecuteTemplate(w, name, data)
}

func (t *TemplateEngine) Load() error {
    // No-op if you’re preloading templates; implement logic here if dynamic loading is needed
    return nil
}


func main() {
		startTime := time.Now()

	// Parse all templates in the views folder
	tmpl := template.Must(template.ParseGlob("views/*.html"))
	version := "0.0.3"

    app := fiber.New(fiber.Config{
        Views: &TemplateEngine{
            templates: tmpl,
        },
    })

	 app.Get("/", func(c *fiber.Ctx) error {
        return c.Render("index.html", fiber.Map{
			"Version": version,
        })
    })

	

	app.Get("/api", func(c *fiber.Ctx) error {
		uptime := time.Since(startTime).Round(time.Second) // rounded for readability
		
		return c.JSON(fiber.Map{
			"success": true,
			"version": fmt.Sprintf("v%s", version),
			"message": "server up and running",
			"uptime":  uptime.String(),
		})
	})

	app.Listen(":3000")
}