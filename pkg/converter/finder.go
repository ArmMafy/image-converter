package converter

import (
	"fmt"
	"os"
	"path/filepath"
)

func GetImages(directoryPath string, imagePool chan *Converter) error {

	directory, _ := os.Open(directoryPath)
	defer directory.Close()
	err := filepath.Walk(directoryPath, func(path string, info os.FileInfo, err error) error {
		switch filepath.Ext(path) {
		case ".jpeg":
			jpegConverter := NewConverter(info.Name())
			jpegConverter.Extension = ".jpeg"
			jpegConverter.SourcePath = path
			imagePool <- jpegConverter
		case ".jpg":
			jpegConverter := NewConverter(info.Name())
			jpegConverter.SourcePath = path
			jpegConverter.Extension = ".jpg"
			imagePool <- jpegConverter
		case ".png":
			pngConverter := NewConverter(info.Name())
			pngConverter.SourcePath = path
			pngConverter.Extension = ".png"
			imagePool <- pngConverter
		default:
			return nil
		}
		fmt.Println(path)
		return nil
	})
	close(imagePool)
	if err != nil {
		return err
	}

	return nil
}
