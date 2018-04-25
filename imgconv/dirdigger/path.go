package dirdigger

import (
	"errors"
	"os"
	"path/filepath"
	"regexp"
)

var (
	// JPEGRegExp is regular expression of JPEG file
	jpegRegExp *regexp.Regexp
	// PNGRegExp is regular expression of JPEG file
	pngRegExp *regexp.Regexp
)

func init() {
	jpegRegExp = regexp.MustCompile("(?i)\\.jpe?g")
	pngRegExp = regexp.MustCompile("(?i)\\.png")
}

// GenerateImageFileListFromDirs returns image file path list in multiple directories.
// "jpg" or "png" are supported.
func GenerateImageFileListFromDirs(dirs []string, imagetype string) ([]string, error) {
	var allFileList []string

	for _, v := range dirs {
		fileList, err := generateImageFileList(v, imagetype)

		if err != nil {
			return nil, err
		}
		allFileList = append(allFileList, fileList...)
	}
	return allFileList, nil
}

// GenerateImageFileList returns image file path list. "jpg" or "png" are supported.
func generateImageFileList(rootdir string, imagetype string) ([]string, error) {

	re, err := FileTypeRegexp(imagetype)
	if err != nil {
		return nil, err
	}
	return generateFileList(rootdir, re)
}

// GenerateFileList returns file path list of files which comply with the regular expression
func generateFileList(rootdir string, filter *regexp.Regexp) ([]string, error) {

	var fileList []string

	err := filepath.Walk(rootdir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() && filter.MatchString(info.Name()) {
			fileList = append(fileList, path)
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return fileList, nil
}

// FileTypeRegexp returns regular expression of filetype
func FileTypeRegexp(imagetype string) (*regexp.Regexp, error) {

	switch imagetype {
	case "jpg":
		return jpegRegExp, nil
	case "png":
		return pngRegExp, nil
	default:
		return nil, errors.New("Unsupported format")
	}

}
