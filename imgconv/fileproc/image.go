package fileproc

import (
	"errors"
	"image"
	"image/jpeg"
	"image/png"
	"os"

	"github.com/translucens/gopher-dojo/imgconv/options"
)

// ConvertImages converts images to specified filetype
func ConvertImages(filelist []string, config options.Options) error {

	for _, v := range filelist {
		err := convertImage(v, config)
		if err != nil {
			return err
		}
	}
	return nil
}

func convertImage(filepath string, config options.Options) error {

	img, err := openImage(filepath, config.InputFileType)
	if err != nil {
		return err
	}

	re, err := FileTypeRegexp(config.InputFileType)
	if err != nil {
		return err
	}

	outputPath := re.ReplaceAllString(filepath, "."+config.OutputFileType)
	_, err = os.Stat(outputPath)
	if err == nil && !config.Overwrite {
		os.Stderr.WriteString("Skipped existing file: " + outputPath + "\n")
		return nil
	}

	err = writeImage(img, outputPath, config)
	if err != nil {
		return err
	}

	return nil
}

func openImage(filepath string, inputFileType string) (image.Image, error) {
	infd, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer infd.Close()

	switch inputFileType {
	case "jpg":
		return jpeg.Decode(infd)
	case "png":
		return png.Decode(infd)
	}
	return nil, errors.New("unsupported file type")
}

func writeImage(img image.Image, outputPath string, config options.Options) error {

	ofd, err := os.OpenFile(outputPath, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0755)
	if err != nil {
		return err
	}

	switch config.OutputFileType {
	case "jpg":
		options := jpeg.Options{Quality: config.JPEGQuality}
		err = jpeg.Encode(ofd, img, &options)
	case "png":
		encoder := png.Encoder{CompressionLevel: png.CompressionLevel(-config.PNGLevel)}
		encoder.Encode(ofd, img)
	}
	if err != nil {
		return err
	}

	return nil
}
