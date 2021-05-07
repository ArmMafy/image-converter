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

func imageDecoder(file *os.File, imagePath Converter) (image.Image, error) {
	var err error
	var img image.Image
	if imagePath.Extension == ".jpeg" {
		img, err = jpeg.Decode(file)
	} else {
		img, err = png.Decode(file)
	}
	if err != nil {
		fmt.Println(err)
	}
	return img, err
}

func imageEncoder(newImage *os.File, imgSet *image.RGBA, imagePath Converter) error {
	if imagePath.Extension == ".jpeg" {
		jpeg.Encode(newImage, imgSet, nil)
	} else {
		png.Encode(newImage, imgSet)
	}
	return nil
}

func Worker(wg *sync.WaitGroup, ch chan Converter, destinationDirectory string) {
	defer wg.Done()

	for {
		imagePath, ok := <-ch
		if !ok {
			break
		}
		imagePath.DestinationPath = destinationDirectory
		file, err := os.Open(path.Join(imagePath.SourcePath, imagePath.Name))
		if err != nil {
			fmt.Println(err)
		}
		defer file.Close()
		img, err := imageDecoder(file, imagePath)
		if err != nil {
			fmt.Println(err)
		}
		b := img.Bounds()
		imgSet := image.NewRGBA(b)

		for y := 0; y < b.Max.Y; y++ {
			for x := 0; x < b.Max.X; x++ {
				imgSet.Set(x, y, color.GrayModel.Convert(img.At(x, y)))
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
