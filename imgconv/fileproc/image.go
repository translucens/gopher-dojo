package fileproc

import (
	"bufio"
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

	buffered := bufio.NewReader(infd)

	switch inputFileType {
	case "jpg":
		return jpeg.Decode(buffered)
	case "png":
		return png.Decode(buffered)
	}
	return nil, errors.New("unsupported file type")
}

func writeImage(img image.Image, outputPath string, config options.Options) error {

	ofd, err := os.OpenFile(outputPath, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0755)
	if err != nil {
		return err
	}
	defer ofd.Close()

	buffered := bufio.NewWriter(ofd)
	defer buffered.Flush()

	switch config.OutputFileType {
	case "jpg":
		options := jpeg.Options{Quality: config.JPEGQuality}
		err = jpeg.Encode(ofd, img, &options)
	case "png":
		encoder := png.Encoder{CompressionLevel: png.CompressionLevel(-config.PNGLevel)}
		encoder.Encode(buffered, img)
	}
	if err != nil {
		return err
	}

	return nil
}
