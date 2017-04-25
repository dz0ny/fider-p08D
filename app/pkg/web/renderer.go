package web

import (
	"fmt"
	"html/template"
	"io"

	"os"

	"io/ioutil"

	"github.com/WeCanHearYou/wechy/app/pkg/env"
	"github.com/labstack/echo"
)

//HTMLRenderer renderer
type HTMLRenderer struct {
	templates map[string]*template.Template
	logger    echo.Logger
}

// NewHTMLRenderer creates a new HTMLRenderer
func NewHTMLRenderer(logger echo.Logger) *HTMLRenderer {
	renderer := &HTMLRenderer{nil, logger}
	renderer.templates = make(map[string]*template.Template)

	renderer.add("index.html")
	renderer.add("404.html")
	renderer.add("500.html")

	return renderer
}

//Render a template based on parameters
func (r *HTMLRenderer) add(name string) *template.Template {
	base := env.Path("/views/base.html")
	file := env.Path("/views", name)
	tpl, err := template.ParseFiles(base, file)
	if err != nil {
		panic(err)
	}

	r.templates[name] = tpl
	return tpl
}

//Render a template based on parameters
func (r *HTMLRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	ctx := Context{Context: c}

	tmpl, ok := r.templates[name]
	if !ok {
		panic(fmt.Errorf("The template '%s' does not exist", name))
	}

	if env.IsDevelopment() {
		tmpl = r.add(name)
	}

	m := data.(echo.Map)
	m["authEndpoint"] = os.Getenv("AUTH_ENDPOINT")
	m["tenant"] = ctx.Tenant()
	m["env"] = env.Current()

	if ctx.IsAuthenticated() {
		m["__User"] = ctx.User()
	}

	files, _ := ioutil.ReadDir("dist/js")
	if len(files) > 0 {
		m["__JavaScriptBundle"] = files[0].Name()
	}

	return tmpl.Execute(w, m)
}