package scaledown

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/lutzpeschlow/file_tools/ctrl"
)

type FileInfo struct {
	FileName string
	Size     int64
}

type FileList struct {
	Files  []FileInfo
	Sorted []FileInfo
}

// is file extension .jpg oder .png
func isJpgOrPng(filename string) bool {
	return strings.HasSuffix(strings.ToLower(filename), ".jpg") ||
		strings.HasSuffix(strings.ToLower(filename), ".png")
}

// output file is getting name with additional characters
func GetOutputFileName(inputFile string) (string, error) {
	dotIndex := strings.LastIndex(inputFile, ".")
	fmt.Print(dotIndex, "\n")
	outputFile := inputFile[:dotIndex] + "_rs" + inputFile[dotIndex:]
	// no dot in file name
	if dotIndex < 1 {
		return "", errors.New("ERROR: no extension found")
	}
	// return renamed output file
	return outputFile, nil
}

// convert input.jpg -resize 2048x2048\> -strip -quality 70 output.jpg
func CmdImageMagick(inputFile string, outputFile string) error {
	fmt.Print(" executing image magick ...", "\n")

	cmd := exec.Command("convert",
		inputFile,
		"-resize", "2048x2048>",
		"-strip",
		"-quality", "70",
		outputFile,
	)

	err := cmd.Run()
	if err != nil {
		fmt.Println("Error :", err)
		return err
	}
	return nil
}

func ScaleDown(ctrl *ctrl.Control_Object) error {
	// variables
	var files []FileInfo
	// loop over files
	err := filepath.Walk(ctrl.Dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		// no dictory but file, add to files
		if !info.IsDir() {
			// fmt.Print(path, "  -  ", info.Size(), "\n")
			files = append(files, FileInfo{
				FileName: path,
				Size:     info.Size(),
			})
			// reduce if size larger than LimitSize
			if info.Size() > int64(ctrl.LimitSize) {
				fmt.Print(path, " with size: ", info.Size(), " \n")
				// executing image magick
				err_cmd := CmdImageMagick(path, outputFile)
				if err_cmd != nil {
					fmt.Print(err_cmd, "\n")
					return err_cmd
				}

			}
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}
