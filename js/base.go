package js

import "syscall/js"

var Window = js.Global()
var Document = Window.Get("document")
var Body = Entity{Obj: Document.Get("body")}
var GlobalObject = Window.Get("Object")

// wrappers of global js functions
func CreateElement(tag string) js.Value {
	return Document.Call("createElement", tag)
}

func CreateTextNode(data string) js.Value {
	return Document.Call("createTextNode", data)
}

type (
	// Entity container for any javascript entity
	Entity struct {
		Obj js.Value
	}

	// Object typical js object
	Object map[string]interface{}

	// properties for creating base tag
	TagProps struct {
		ID        string
		Style     Object
		ClassName string
	}

	// anything that can produce a js entity
	ValueHolder interface {
		GetJSValue() js.Value
	}
)

func (e Entity) GetJSValue() js.Value {
	return e.Obj
}

func (e *Entity) SetInnerText(innerText string) {
	e.Obj.Set("innerHTML", innerText)
}

func (e *Entity) AppendChild(obj ValueHolder) {
	if obj != nil {
		e.Obj.Call("appendChild", obj.GetJSValue())
	}
}

func (e *Entity) appendChildren(objs []ValueHolder) {
	for _, obj := range objs {
		e.AppendChild(obj)
	}
}

func (e *Entity) ClearChildren() {
	e.Obj.Call("replaceChildren")
}

func (e *Entity) ReplaceChildren(objs ...ValueHolder) {
	e.ClearChildren()
	e.appendChildren(objs)
}

func (e *Entity) SetValue(key, value string) {
	if value == "" {
		e.Obj.Delete(key)
	} else {
		e.Obj.Set(key, value)
	}
}

func (e *Entity) SetId(id string) {
	e.SetValue("id", id)
}

func (e *Entity) SetStyle(style Object) {
	GlobalObject.Call("assign", e.Obj.Get("style"), style.raw())
}

func (e *Entity) SetClassName(className string) {
	e.SetValue("className", className)
}

func (e *Entity) SetProps(props TagProps) {
	e.SetId(props.ID)
	e.SetStyle(props.Style)
	e.SetClassName(props.ClassName)
}

func (r *Object) raw() map[string]interface{} {
	if r == nil {
		return nil
	}
	return *r
}

// some simple tags

func CreateDiv() Entity {
	return Entity{Obj: CreateElement("div")}
}

func CreateDivWithProps(props TagProps) Entity {
	o := CreateDiv()
	o.SetProps(props)
	return o
}

func CreateBreak() Entity {
	return Entity{Obj: CreateElement("br")}
}
