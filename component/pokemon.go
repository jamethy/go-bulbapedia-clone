package component

import (
	"context"

	"go-bulbapedia-clone/js"
	"go-bulbapedia-clone/poke"
)

type Pokemon struct {
	js.Entity
	// todo link to previous
	Name js.Text
	// todo types
	Image js.Image
	// todo line break
	// todo link to next
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
	loader.Start(context.Background())
	parentDiv := js.CreateDivWithProps(js.TagProps{
		ClassName: "centered-container",
	})
	parentDiv.AppendChild(loader)
	return parentDiv
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
		Image:  image,
	}
}
