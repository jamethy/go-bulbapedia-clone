package component

import "go-bulbapedia-clone/js"

type Error struct {
	js.Entity
}

func CreateError(err error) Error {
	return Error{Entity: js.Entity{Obj: js.CreateTextNode(err.Error())}}
}
