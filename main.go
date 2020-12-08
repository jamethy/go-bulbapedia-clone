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

	// set up routes for content

	js.GlobalRouter.AddRouteWithParams("/pokemon/{id:\\d+}", func(m map[string]string) js.ValueHolder {
		id, _ := strconv.Atoi(m["id"])
		return component.PokemonPage(id)
	})

	// temp - load bulbasaur for home page
	js.GlobalRouter.AddRoute("/", func() js.ValueHolder {
		return component.PokemonPage(1)
	})

	js.GlobalRouter.Parent = js.CreateDiv()
	js.GlobalRouter.LoadRoute(js.GetCurrentRoute())

	// finally create the body
	js.Body.ReplaceChildren(
		createHeader(),
		js.GlobalRouter.Parent,
		createFooter(),
	)

	// this keeps lib.wasm running so js functions are still available
	// not completely clear why necessary
	c := make(chan struct{}, 0)
	<-c
}

func createHeader() js.Entity {
	header := js.CreateTag("header")
	header.SetInnerHTML(
		// language=HTML
		"<h1>Bulbapedia-Clone</h1>",
	)
	return header
}

// createFooter using go templates because why not
func createFooter() js.Entity {
	type test struct {
		ValueOne string
	}

	tmpl, _ := template.New("footer").Parse(
		// language=GoHTML
		"<span>Brought to you by: {{.ValueOne}}</span>",
	)

	var sb = strings.Builder{}
	_ = tmpl.Execute(&sb, test{ValueOne: "me"})

	div := js.CreateDiv()
	div.SetInnerHTML(sb.String())
	return div
}
