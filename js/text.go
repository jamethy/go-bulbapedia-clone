package js

import "syscall/js"

type (
	Text struct {
		Entity
	}
	TextProps struct {
		TagProps
		InnerText string
		Tag       string
	}
)

func CreateTextWithProps(props TextProps) Text {
	var obj js.Value
	if props.Tag == "" {
		obj = CreateTextNode(props.InnerText)
	} else {
		obj = CreateElement(props.Tag)
	}
	t := Text{Entity{Obj: obj}}
	t.SetProps(props)
	return t
}
func (t *Text) SetProps(props TextProps) {
	t.Entity.SetProps(props.TagProps)
	t.SetInnerText(props.InnerText)
}
