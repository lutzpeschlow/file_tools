package sizing

import (
	"fmt"
	"os"
	"path/filepath"
)

type FileInfo struct {
	FileName string
	Size     int64
}

type FileList struct {
	Files []FileInfo
}

func GetFileList(path string, obj *FileList) error {
	var files []FileInfo

	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		// no dictory but file, add to files
		if !info.IsDir() {
			fmt.Println(path)
			files = append(files, FileInfo{
				FileName: path,
				Size:     info.Size(),
			})
		}

		return nil
	})
	if err != nil {
		return err
	}
	return nil
}
