package component

import (
	"context"

	"go-bulbapedia-clone/js"
)

type (
	// Loading is a container that shows a loading image while waiting for the API
	Loading struct {
		js.Entity
	}

	// function to load object from api
	ObjectLoader func(ctx context.Context) (js.ValueHolder, error)

	LoadingContainer struct {
		js.Entity
		Loader ObjectLoader
	}
)

func CreateLoading() Loading {
	return Loading{Entity: js.CreateDivWithProps(js.TagProps{
		ClassName: "loader",
	})}
}

func CreateLoadingContainer(f ObjectLoader) LoadingContainer {
	return LoadingContainer{
		Entity: js.CreateDiv(),
		Loader: f,
	}
}

func (l *LoadingContainer) Start(ctx context.Context) {
	l.AppendChild(CreateLoading())
	go func() {
		res, err := l.Loader(ctx)
		l.ClearChildren()
		if err != nil {
			l.AppendChild(CreateError(err))
		} else {
			l.AppendChild(res)
		}
	}()
}