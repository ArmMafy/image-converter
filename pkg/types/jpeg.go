package types

// import (
// 	"education/images/pkg/converter"
// 	"fmt"
// )

// type JpegImage struct {
// 	Path string
// 	source
// 	destination
// 	name
// }

// func NewJpegImage(path string) (converter.Convertible, error) {
// 	return &JpegImage{Path: path}, nil
// }

// func check(err error) {
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// }

// func (imagePath *JpegImage) ConvertToGrey() {
// file, err := os.Open(imagePath.Path)
// check(err)
// defer file.Close()
// img, err := jpeg.Decode(file)
// check(err)

// b := img.Bounds()
// imgSet := image.NewRGBA(b)
// for y := 0; y < b.Max.Y; y++ {
// 	for x := 0; x < b.Max.X; x++ {
// 		imgSet.Set(x, y, color.GrayModel.Convert(img.At(x, y)))
// 	}
// }
// outFile, err := os.Create()
// check(err)
// jpeg.Encode(outFile, imgSet, nil)
