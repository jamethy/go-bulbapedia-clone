package poke

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"go-bulbapedia-clone/ctxhttp"
)

const baseUrl = "https://pokeapi.co/api/v2"

func GetById(ctx context.Context, id int) (p Pokemon, err error) {
	url := baseUrl + "/pokemon/" + strconv.Itoa(id)

	time.Sleep(time.Second) // fake wait time to test loading
	_, err = ctxhttp.GetJsonWithParams(ctx, http.DefaultClient, url, nil, &p)
	if err != nil {
		return p, fmt.Errorf("problem with api: %w", err)
	}
	return p, nil
}

func GetByName(ctx context.Context, name string) (p Pokemon, err error) {
	url := baseUrl + "/pokemon/" + name

	_, err = ctxhttp.GetJsonWithParams(ctx, http.DefaultClient, url, nil, &p)
	if err != nil {
		return p, fmt.Errorf("problem with api: %w", err)
	}
	return p, nil
}
