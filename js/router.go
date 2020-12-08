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
		Parent     Entity
	}
)

var GlobalRouter = Router{
	Parent: Body,
}

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

func (r *Router) LoadRoute(route string) {
	page := GlobalRouter.GetPage(route)
	// todo create wrapper so only inner stuff is loaded

	r.Parent.ClearChildren()
	if page != nil {
		r.Parent.AppendChild(page)
	} else {
		route = "/not-found"
		r.Parent.AppendChild(Entity{Obj: CreateTextNode("404 not found")})
	}
	Window.Get("history").Call("pushState", (&Object{}).raw(), "", route)
}

func RouteLoader(route string) func() {
	return func() {
		GlobalRouter.LoadRoute(route)
	}
}