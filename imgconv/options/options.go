package options

import (
	"errors"
	"flag"
	"image/jpeg"
	"image/png"
)

// Options contains settings of image converter
type Options struct {
	Overwrite      bool
	InputFileType  string
	OutputFileType string
	InputDirs      []string
	JPEGQuality    int
	PNGLevel       int
}

// Config contains program options
var Config Options

func init() {
	const (
		defaultOverwrite = false
		usageOverwrite   = "specify overwrite option, when there are same name file(s), then they will be overwritten."

		defaultInputType = "jpg"
		usageInputType   = "input filetype (jpg, png)"

		defaultOutputType = "png"
		usageOutputType   = "output filetype (jpg, png)"

		defaultJPEGQuality = jpeg.DefaultQuality
		usageJPEGQuality   = "JPEG quality. 1(smaller)-100(better quality)"

		defaultPNGLevel = int(png.DefaultCompression)
		usagePNGLevel   = "PNG compression level. 0: Default, 1: No compression, 2: Best speed, 3: Best compression"
	)

	flag.BoolVar(&Config.Overwrite, "overwrite", defaultOverwrite, usageOverwrite)
	flag.StringVar(&Config.InputFileType, "intype", defaultInputType, usageInputType)
	flag.StringVar(&Config.OutputFileType, "outtype", defaultOutputType, usageOutputType)
	flag.IntVar(&Config.JPEGQuality, "jpegq", defaultJPEGQuality, usageJPEGQuality)
	flag.IntVar(&Config.PNGLevel, "pnglv", defaultPNGLevel, usagePNGLevel)
}

// InitOptions initialize options and verifies them
func InitOptions() error {
	flag.Parse()
	Config.InputDirs = flag.Args()
	return Config.verify()
}

// Verify option values
func (o *Options) verify() error {

	switch o.InputFileType {
	case "png", "jpg":
	case "jpeg":
		o.InputFileType = "jpg"
	default:
		return errors.New("unsupported input file type")
	}

	switch o.OutputFileType {
	case "png":
		if o.PNGLevel < 0 || 3 < o.PNGLevel {
			return errors.New("PNG compression level must be 0 to 3")
		}
	case "jpeg":
		o.OutputFileType = "jpg"
		fallthrough
	case "jpg":
		if o.JPEGQuality < 1 || 100 < o.JPEGQuality {
			return errors.New("JPEG quality must be 1 to 100")
		}
	default:
		return errors.New("unsupported output file type")
	}

	if len(o.InputDirs) == 0 {
		return errors.New("input directories are required")
	}

	return nil
}
