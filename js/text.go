package js

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
	t := Text{}
	if props.Tag == "" {
		t.Obj = CreateTextNode(props.InnerText)
	} else {
		t.Entity = CreateTag(props.Tag)
	}
	t.SetProps(props)
	return t
}
func (t *Text) SetProps(props TextProps) {
	t.Entity.SetProps(props.TagProps)
	t.SetInnerHTML(props.InnerText)
}
