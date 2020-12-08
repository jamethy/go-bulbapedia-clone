// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package ctxhttp provides helper functions for performing context-aware HTTP requests.
package ctxhttp // import "golang.org/x/net/context/ctxhttp"

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"go-bulbapedia-clone/util"
)

const GET = "GET"

type Params interface {
	IsParams()
}

type RequestBody interface {
	IsRequestBody()
}

type ResponseBody interface {
	IsResponseBody()
}

// Do sends an HTTP request with the provided http.Client and returns
// an HTTP response.
//
// If the client is nil, http.DefaultClient is used.
//
// The provided ctx must be non-nil. If it is canceled or times out,
// ctx.Err() will be returned.
func Do(ctx context.Context, client *http.Client, req *http.Request) (*http.Response, error) {
	if client == nil {
		client = http.DefaultClient
	}
	resp, err := client.Do(req.WithContext(ctx))
	// If we got an error, and the context has been canceled,
	// the context's error is probably more useful.
	if err != nil {
		select {
		case <-ctx.Done():
			err = ctx.Err()
		default:
		}
	}
	return resp, err
}

// GetPage issues a GET request via the Do function.
func Get(ctx context.Context, client *http.Client, url string) (*http.Response, error) {
	req, err := http.NewRequest(GET, url, nil)
	if err != nil {
		return nil, err
	}
	return Do(ctx, client, req)
}

// GetJson issues a GET request via the GetPage function and unmarshals the body
func GetJson(ctx context.Context, client *http.Client, url string, body ResponseBody) (*http.Response, error) {
	res, err := Get(ctx, client, url)
	if err != nil {
		return nil, fmt.Errorf("failed to make GET request: %w", err)
	}

	if res.StatusCode != 200 {
		return res, fmt.Errorf("received %d", res.StatusCode)
	}

	defer util.SafeClose(res.Body)
	err = json.NewDecoder(res.Body).Decode(body)
	if err != nil {
		return res, fmt.Errorf("failed to unmarshal body: %w", err)
	}
	return res, nil
}

// GetJson issues a GET request via the GetPage function and unmarshals the body
func GetJsonWithParams(ctx context.Context, client *http.Client, url string, params Params, body ResponseBody) (*http.Response, error) {
	// add query parameters
	if params != nil {
		var err error
		url, err = AddQueryParameters(url, params)
		if err != nil {
			return nil, fmt.Errorf("cannot add params: %w", err)
		}
	}
	return GetJson(ctx, client, url, body)
}
