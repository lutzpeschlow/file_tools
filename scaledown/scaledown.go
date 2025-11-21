package scaledown

import (
	"fmt"
	"os"
	"path/filepath"

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

func ScaleDown(ctrl *ctrl.Control_Object) error {
	var files []FileInfo
	// var sorted []FileInfo

	err := filepath.Walk(ctrl.Dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		// no dictory but file, add to files
		if !info.IsDir() {
			fmt.Print(path, "  -  ", info.Size(), "\n")
			files = append(files, FileInfo{
				FileName: path,
				Size:     info.Size(),
			})
			// fmt.Print(info.Size(), " - ", info.Name(), "\n")
			if info.Size() > 4000000 {
				fmt.Print(" !!! ", path, " \n")
			}
		}

		return nil
	})
	if err != nil {
		return err
	}

	// sort.Slice(files, func(i, j int) bool {
	// 	return files[i].Size > files[j].Size
	//})
	// for i, file := range files {
	//	if i <= ctrl.Num {
	//		fmt.Print(" ", file.Size, " ", file.FileName, " ", i, " \n")
	//		sorted = append(sorted, FileInfo{FileName: file.FileName, Size: file.Size})
	//	}
	// }

	return nil
}
