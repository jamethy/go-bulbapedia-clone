package js

import (
	"regexp"
	"strings"

	"go-bulbapedia-clone/util"
)

type (
	PageFunc func(map[string]string) ValueHolder

	Router struct {
		routes     map[string]PageFunc
		routeRegex map[string]*regexp.Regexp
	}
)

// GetCurrentRoute gets the location of the current page minus the host
func GetCurrentRoute() string {
	href := Window.Get("location").Get("href").String()
	if strings.Count(href, "/") <= 2 {
		return "/"
	}
	return "/" + strings.SplitN(href, "/", 4)[3]
}

// AddRoute add a page for a route
// route can contain variables with regex, such as /page/my-{variable}/with-id-{id:\\d+}
func (r *Router) AddRoute(route string, pageFunc PageFunc) {
	if r.routes == nil {
		r.routes = make(map[string]PageFunc)
		r.routeRegex = make(map[string]*regexp.Regexp)
	}
	r.routes[route] = pageFunc
	r.routeRegex[route] = util.RouteToRegex(route)
}

// GetPage for currentRoute
func (r *Router) GetPage(currentRoute string) ValueHolder {
	for route, rgx := range r.routeRegex {
		if values, ok := util.ParseRoute(rgx, currentRoute); ok {
			return r.routes[route](values)
		}
	}
	return nil
}
