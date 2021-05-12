package converter

import (
	"fmt"
	"image"
	"image/color"
	"image/jpeg"
	"image/png"
	"os"
	"path"
	"sync"
)

// func imageDecoder(file *os.File, imagePath *Converter) (image.Image, error) {
// 	var err error
// 	var img image.Image
// 	switch imagePath.Extension {
// 	case ".jpeg":
// 		img, err = jpeg.Decode(file)
// 	case ".jpg":
// 		img, err = jpeg.Decode(file)
// 	case ".png":
// 		img, err = png.Decode(file)
// 	default:
// 		break
// 	}
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	return img, err
// }

func imageEncoder(newImage *os.File, imgSet *image.RGBA, imagePath *Converter) error {
	switch imagePath.Extension {
	case ".jpeg":
		jpeg.Encode(newImage, imgSet, nil)
	case ".jpg":
		jpeg.Encode(newImage, imgSet, nil)
	case ".png":
		png.Encode(newImage, imgSet)
	}
	return nil
}

func Worker(wg *sync.WaitGroup, ch chan *Converter, destinationDirectory string) {
	defer wg.Done()

	for {
		imagePath, ok := <-ch
		if !ok {
			break
		}

		imagePath.DestinationPath = destinationDirectory
		file, err := os.Open(imagePath.SourcePath)
		if err != nil {
			fmt.Println(err)
		}
		defer file.Close()

		img, _, err := image.Decode(file)
		if err != nil {
			fmt.Println(err)
		}

		b := img.Bounds()
		imgSet := image.NewRGBA(b)
		for z := 0; z < b.Max.Y; z++ {
			for x := 0; x < b.Max.X; x++ {
				oldPixel := img.At(x, z)
				pixel := color.GrayModel.Convert(oldPixel)
				imgSet.Set(x, z, pixel)
			}
		}

		newImage, err := os.Create(path.Join(imagePath.DestinationPath, imagePath.Name))
		if err != nil {
			fmt.Println(err)
		}
		defer newImage.Close()
		imageEncoder(newImage, imgSet, imagePath)

	}
}
