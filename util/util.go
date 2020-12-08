package util

import (
	"fmt"
	"io"
	"log"
	"net/url"
	"regexp"
	"runtime/debug"
	"strings"
)

func SafeClose(closer io.Closer) {
	if closer != nil {
		err := closer.Close()
		if err != nil {
			log.Println("ERROR: failed to close closer: ", err)
			debug.PrintStack()
		}
	}
}

var partFinder = regexp.MustCompile("{[^}]+}")
func RouteToRegex(route string) *regexp.Regexp {
	route = partFinder.ReplaceAllStringFunc(route, func(part string) string {
		part = part[1:len(part)-1]
		innerRegex := "[^/]+"
		partsOfPart := strings.SplitN(part, ":", 2)
		if len(partsOfPart) > 1 {
			innerRegex = partsOfPart[1]
		}
		return fmt.Sprintf("(?P<%s>%s)", partsOfPart[0], innerRegex)
	})

	return regexp.MustCompile("^" + route + "$")
}

func ParseRoute(routeRegex *regexp.Regexp, currentRoute string) (map[string]string, bool) {
	routeParts := strings.SplitN(currentRoute, "?", 2)
	matches := routeRegex.FindStringSubmatch(routeParts[0])
	if matches == nil {
		return nil, false
	}

	res := make(map[string]string, len(routeRegex.SubexpNames()))
	for i, s := range routeRegex.SubexpNames() {
		if i != 0 {
			res[s] = matches[i]
		}
	}

	if len(routeParts) > 1 {
		pairs := strings.Split(routeParts[1], "&")
		for _, p := range pairs {
			splitPair := strings.Split(p, "=")
			if len(splitPair) != 2 {
				res[p] = "true"
			} else {
				key, err := url.QueryUnescape(splitPair[0])
				if err != nil {
					continue
				}
				value, err := url.QueryUnescape(splitPair[1])
				if err != nil {
					continue
				}
				res[key] = value
			}
		}
	}
	return res, true
}