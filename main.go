package main

import (
	"html/template"
	"strconv"
	"strings"

	"go-bulbapedia-clone/component"
	"go-bulbapedia-clone/js"
)

// trying to clone https://bulbapedia.bulbagarden.net/wiki/Bulbasaur_(Pok%C3%A9mon)
func main() {

	// create header, content, and footer
	header := js.CreateTag("header")
	header.SetInnerHTML(
		// language=HTML
		"<h1>Bulbapedia-Clone</h1>",
	)

	content := js.CreateDiv()
	footer := js.CreateTag("footer")

	js.Body.ReplaceChildren(
		header,
		content,
		footer,
	)

	// set up routes for content
	js.GlobalRouter.Parent = content

	js.GlobalRouter.AddRoute("/pokemon/{id:\\d+}", func(m map[string]string) js.ValueHolder {
		id, _ := strconv.Atoi(m["id"])
		return component.PokemonPage(id)
	})

	// temp - load bulbasaur for home page
	js.GlobalRouter.AddRoute("/", func(m map[string]string) js.ValueHolder {
		return component.PokemonPage(1)
	})

	js.GlobalRouter.AddRoute("/template-test", templateTest)

	// this is called on page load
	js.GlobalRouter.LoadRoute(js.GetCurrentRoute())

	// this keeps lib.wasm running so js functions are still available
	// not completely clear why necessary
	c := make(chan struct{}, 0)
	<-c
}

func templateTest(_ map[string]string) js.ValueHolder {
	div := js.CreateDiv()
	type test struct {
		ValueOne string
	}

	tmpl, _ := template.New("test").Parse(
		// language=GoHTML
		"<h3>This is a test: {{.ValueOne}}</h3>",
	)

	var sb = strings.Builder{}
	_ = tmpl.Execute(&sb, test{ValueOne: "value one"})

	div.SetInnerHTML(sb.String())
	return div
}
