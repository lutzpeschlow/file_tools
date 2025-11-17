package sizing

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"

	"github.com/lutzpeschlow/file_tools/ctrl"
)

type FileInfo struct {
	FileName string
	Size     int64
}

type FileList struct {
	Files []FileInfo
}

func GetFileList(ctrl *ctrl.Control_Object, obj *FileList) error {
	var files []FileInfo

	err := filepath.Walk(ctrl.Dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		// no dictory but file, add to files
		if !info.IsDir() {

			files = append(files, FileInfo{
				FileName: path,
				Size:     info.Size(),
			})
			// fmt.Print(info.Size(), " - ", info.Name(), "\n")
		}

		return nil
	})
	if err != nil {
		return err
	}

	sort.Slice(files, func(i, j int) bool {
		return files[i].Size > files[j].Size
	})
	for i, file := range files {
		if i <= ctrl.Num {
			fmt.Print(" ", file.Size, " ", file.FileName, " ", i, " \n")
		}

	}

	obj.Files = files
	return nil
}
