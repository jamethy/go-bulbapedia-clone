package js

type (
	Image struct {
		Entity
	}
	ImageProps struct {
		TagProps
		Src string
	}
)

func CreateImageWithProps(props ImageProps) Image {
	b := Image{Entity{Obj: CreateElement("img")}}
	b.SetProps(props)
	return b
}

func (i *Image) SetSrc(src string) {
	i.SetValue("src", src)
}

func (i *Image) SetProps(props ImageProps) {
	i.Entity.SetProps(props.TagProps)
	i.SetSrc(props.Src)
}