package component

import (
	"context"
	"fmt"

	"go-bulbapedia-clone/js"
	"go-bulbapedia-clone/poke"
)

type Pokemon struct {
	js.Entity
}

func PokemonPage(id int) js.ValueHolder {
	parentDiv := js.CreateDivWithProps(js.TagProps{
		ClassName: "centered-container",
	})
	prevButton := js.CreateButtonWithProps(js.ButtonProps{
		InnerText: "Prev",
		OnClick: func() {
			if id == 1 {
				id = 152
			}
			js.GlobalRouter.LoadRoute(fmt.Sprintf("/pokemon/%d", id-1))
		},
	})
	nextButton := js.CreateButtonWithProps(js.ButtonProps{
		InnerText: "next",
		OnClick: func() {
			if id == 151 {
				id = 0
			}
			js.GlobalRouter.LoadRoute(fmt.Sprintf("/pokemon/%d", id+1))
		},
	})
	parentDiv.AppendChild(prevButton)
	parentDiv.AppendChild(LoadPokemonById(id))
	parentDiv.AppendChild(nextButton)

	return parentDiv
}

// LoadPokemonById creates a centered container that presents the requested pokemon
// shows a loading spinner while waiting on the API
func LoadPokemonById(id int) js.ValueHolder {
	loader := CreateLoadingContainer(func(ctx context.Context) (js.ValueHolder, error) {
		p, err := poke.GetById(context.Background(), id)
		if err != nil {
			return js.Entity{}, err
		}
		return CreatePokemon(p), nil
	})
	loader.SetProps(js.TagProps{
		ClassName: "poke-container",
	})
	loader.Start(context.Background())
	return loader
}

// CreatePokemon creates a pokemon display - just the image with the name above
func CreatePokemon(p poke.Pokemon) Pokemon {
	image := js.CreateImageWithProps(js.ImageProps{
		Src: p.Sprites.FrontDefault,
	})
	name := js.CreateTextWithProps(js.TextProps{
		InnerText: p.Name,
		Tag:       "h3",
		TagProps: js.TagProps{
			Style: js.Object{
				"text-align": "center",
			},
		},
	})
	parentDiv := js.CreateDiv()
	parentDiv.ReplaceChildren(
		name,
		image,
	)
	return Pokemon{
		Entity: parentDiv,
	}
}
