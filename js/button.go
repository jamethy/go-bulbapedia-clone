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

func CreateButtonWithProps(props ButtonProps) Button {
	b := Button{CreateTag("button")}
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
	b.SetInnerHTML(props.InnerText)
	b.SetOnClick(props.OnClick)
}
