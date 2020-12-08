package main

import (
	"strconv"

	"go-bulbapedia-clone/component"
	"go-bulbapedia-clone/js"
)

func main() {

	router := js.Router{}

	router.AddRoute("/pokemon/{id:\\d+}", func(m map[string]string) js.ValueHolder {
		id, _ := strconv.Atoi(m["id"])
		return component.LoadPokemonById(id)
	})

	// temp - load bulbasaur for home page
	router.AddRoute("/", func(m map[string]string) js.ValueHolder {
		return component.LoadPokemonById(1)
	})

	// this is called on page load
	// todo figure out how to navigate within js
	page := router.GetPage(js.GetCurrentRoute())
	if page != nil {
		js.Body.AppendChild(page)
	} else {
		js.Body.AppendChild(js.Entity{Obj: js.CreateTextNode("404 not found")})
	}

	// this keeps lib.wasm running so js functions are still available
	// not completely clear why necessary
	c := make(chan struct{}, 0)
	<-c
}
