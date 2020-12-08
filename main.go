package main

import (
	"strconv"

	"go-bulbapedia-clone/component"
	"go-bulbapedia-clone/js"
)

func main() {

	js.GlobalRouter.AddRoute("/pokemon/{id:\\d+}", func(m map[string]string) js.ValueHolder {
		id, _ := strconv.Atoi(m["id"])
		return component.PokemonPage(id)
	})

	// temp - load bulbasaur for home page
	js.GlobalRouter.AddRoute("/", func(m map[string]string) js.ValueHolder {
		return component.PokemonPage(1)
	})

	// this is called on page load
	js.GlobalRouter.LoadRoute(js.GetCurrentRoute())

	// this keeps lib.wasm running so js functions are still available
	// not completely clear why necessary
	c := make(chan struct{}, 0)
	<-c
}
