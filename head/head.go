package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

var (
	printLines int
	varbose    bool
	quiet      bool
)

func init() {
	const (
		defaultLines = 10
		usageLines   = "print the first NUM lines instead of the first 10. If specify negative number, print the last NUM lines"

		defaultVarbose = false
		usageVarbose   = "print header giving file name when specify single file"

		defaultQuiet = false
		usageQuiet   = "never print headers giving file names"
	)
	flag.IntVar(&printLines, "n", defaultLines, usageLines)
	flag.IntVar(&printLines, "lines", defaultLines, usageLines)

	flag.BoolVar(&varbose, "v", defaultVarbose, usageVarbose)
	flag.BoolVar(&varbose, "varbose", defaultVarbose, usageVarbose)

	flag.BoolVar(&quiet, "q", defaultQuiet, usageQuiet)
	flag.BoolVar(&quiet, "quiet", defaultQuiet, usageQuiet)
}

func main() {

	flag.Parse()
	files := flag.Args()

	for _, filename := range files {
		fd, err := os.Open(filename)
		if err != nil {
			errorHandler("Can't open file: " + filename)
			break
		}
		defer fd.Close()

		if !quiet && (varbose || len(files) > 1) {
			fmt.Printf("**** %s ****\n", filename)
		}

		var lines []string

		switch {
		case printLines > 0:
			lines, err = GetHead(fd, printLines)
		case printLines < 0:
			lines, err = GetTail(fd, -printLines)
		}

		if err != nil {
			errorHandler("Error occured while reading: " + err.Error())
			break
		}

		for _, line := range lines {
			fmt.Println(line)
		}
	}
}

func GetHead(fd *os.File, lines int) ([]string, error) {

	scanner := bufio.NewScanner(fd)
	var buffer []string

	for i := 0; (i < lines) && scanner.Scan(); i++ {
		buffer = append(buffer, scanner.Text())
	}

	return buffer, scanner.Err()
}

func GetTail(fd *os.File, lines int) ([]string, error) {
	scanner := bufio.NewScanner(fd)
	buffer := make([]string, lines)
	lineCount := 0

	for ; scanner.Scan(); lineCount++ {
		buffer[lineCount%len(buffer)] = scanner.Text()
	}
	if scanner.Err() != nil {
		fmt.Println(scanner.Err().Error())
		return nil, scanner.Err()
	}

	var tailLines []string
	if lineCount >= lines {
		tailLines = append(tailLines, buffer[lineCount%lines:]...)
	}

	tailLines = append(tailLines, buffer[0:lineCount%lines]...)

	return tailLines, nil
}

func errorHandler(err string) {
	os.Stderr.WriteString(err)
}
