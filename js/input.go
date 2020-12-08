package js

type (
	Input struct {
		Entity
	}

	InputProps struct {
		TagProps
		Type string
	}
)

func CreateInputWithProps(props InputProps) Input {
	b := Input{Entity{Obj: CreateElement("input")}}
	b.SetProps(props)
	return b
}

func (i *Input) SetType(t string) {
	i.SetValue("type", t)
}

func (i *Input) SetProps(props InputProps) {
	i.Entity.SetProps(props.TagProps)
	i.SetType(props.Type)
}

func (i *Input) GetValue() string {
	return i.Obj.Get("value").String()
}