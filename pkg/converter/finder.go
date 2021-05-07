package converter

import (
	"os"
	"path/filepath"
)

func GetImages(directoryPath string, imagePool chan Converter) error {

	err := filepath.Walk(directoryPath, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		fileExtension := filepath.Ext(path)
		switch fileExtension {
		case ".jpeg":
			jpegConverter := NewConverter(path)
			jpegConverter.Extension = ".jpeg"
			imagePool <- *jpegConverter
		case ".png":
			pngConverter := NewConverter(path)
			pngConverter.Extension = ".png"
			imagePool <- *pngConverter
		default:
			return nil
		}
		return nil
	})
	close(imagePool)
	if err != nil {
		return err
	}

	return nil
}
