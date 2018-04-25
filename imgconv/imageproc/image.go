package imageproc

import (
	"bufio"
	"errors"
	"image"
	"image/jpeg"
	"image/png"
	"io"
	"os"

	"github.com/translucens/gopher-dojo/imgconv/dirdigger"
	"github.com/translucens/gopher-dojo/imgconv/options"
)

// ConvertImages converts images to specified filetype
func ConvertImages(filelist []string, config options.Options) error {

	for _, v := range filelist {
		err := ConvertImage(v, config)
		if err != nil {
			return err
		}
	}
	return nil
}

// ConvertImage converts single image
func ConvertImage(filepath string, config options.Options) error {

	img, err := OpenImageFile(filepath, config.InputFileType)
	if err != nil {
		os.Stderr.WriteString("Skipped file: " + filepath + "; " + err.Error() + "\n")
		return nil
	}

	re, err := dirdigger.FileTypeRegexp(config.InputFileType)
	if err != nil {
		return err
	}

	outputPath := re.ReplaceAllString(filepath, "."+config.OutputFileType)
	_, err = os.Stat(outputPath)
	if err == nil && !config.Overwrite {
		os.Stderr.WriteString("Skipped existing file: " + outputPath + "\n")
		return nil
	}

	err = WriteImageFile(img, outputPath, config)
	if err != nil {
		return err
	}

	return nil
}

// OpenImageFile opens image file
func OpenImageFile(filepath string, inputFileType string) (image.Image, error) {
	infd, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer infd.Close()

	return OpenImage(infd, inputFileType)
}

// OpenImage reads image from io.Reader
func OpenImage(reader io.Reader, inputFileType string) (image.Image, error) {

	buffered := bufio.NewReader(reader)

	switch inputFileType {
	case "jpg":
		return jpeg.Decode(buffered)
	case "png":
		return png.Decode(buffered)
	}
	return nil, errors.New("unsupported file type")
}

// WriteImageFile writes image to file
func WriteImageFile(img image.Image, outputPath string, config options.Options) error {

	ofd, err := os.OpenFile(outputPath, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0755)
	if err != nil {
		return err
	}
	defer ofd.Close()

	return WriteImage(img, ofd, config)
}

// WriteImage writes image to io.Writer
func WriteImage(img image.Image, writer io.Writer, config options.Options) error {

	buffered := bufio.NewWriter(writer)
	defer buffered.Flush()

	switch config.OutputFileType {
	case "jpg":
		options := jpeg.Options{Quality: config.JPEGQuality}
		return jpeg.Encode(buffered, img, &options)
	case "png":
		encoder := png.Encoder{CompressionLevel: png.CompressionLevel(-config.PNGLevel)}
		return encoder.Encode(buffered, img)
	}
	return errors.New("Unsupported image type")
}
