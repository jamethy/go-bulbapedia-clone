package js

import (
	"syscall/js"
)

type (
	Button struct {
		Entity
	}
	ButtonProps struct {
		TagProps
		InnerText string
		OnClick   func()
	}
)

func CreateButton(innerText string) Button {
	b := Button{Entity{Obj: CreateElement("button")}}
	b.SetInnerText(innerText)
	return b
}

func CreateButtonWithProps(props ButtonProps) Button {
	b := Button{Entity{Obj: CreateElement("button")}}
	b.SetProps(props)
	return b
}

func (b *Button) SetOnClick(f func()) {
	if f == nil {
		b.Obj.Set("onclick", nil)
		return
	}
	b.Obj.Set("onclick", js.FuncOf(func(_ js.Value, _ []js.Value) interface{} {
		f()
		return nil
	}))
}

func (b *Button) SetProps(props ButtonProps) {
	b.Entity.SetProps(props.TagProps)
	b.SetInnerText(props.InnerText)
	b.SetOnClick(props.OnClick)
}