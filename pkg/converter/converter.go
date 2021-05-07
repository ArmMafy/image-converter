package converter

// type Convertible interface {
// 	GetImages()
// 	ConvertToGrey()
// }

type Converter struct {
	Name            string
	Extension       string
	SourcePath      string
	DestinationPath string
}

func NewConverter(Name string) *Converter {
	n := Converter{Name: Name}
	return &n
}
