package db

import (
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
	Files  []FileInfo
	Sorted []FileInfo
}

// function GetFileList
// lists files and size
//
// possible gui: Fyne, Gio, go-gtk, Wails, Unison
func GetFileList(ctrl *ctrl.Control_Object, obj *FileList) error {
	var files []FileInfo
	var sorted []FileInfo

	err := filepath.Walk(ctrl.Dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		// no dictory but file, add to files
		if !info.IsDir() && filepath.Ext(path) == ".db" {

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
			// fmt.Print(" ", file.Size, " ", file.FileName, " ", i, " \n")
			sorted = append(sorted, FileInfo{FileName: file.FileName, Size: file.Size})
		}

	}

	obj.Files = files
	obj.Sorted = sorted

	// // remove database if   delete    is activated in ctrl.View
	// if ctrl.View == "delete" && len(sorted) > 0 {
	// 	for i := 0; i < int(ctrl.Num) && i < len(sorted); i++ {
	// 		dbPath := filepath.Join(ctrl.Dir, sorted[i].FileName)
	// 		if err := os.Remove(dbPath); err != nil {
	// 			fmt.Printf("ERROR: %s - %v\n", dbPath, err)
	// 		} else {
	// 			fmt.Printf("REMOVED: %s\n", dbPath)
	// 		}
	// 	}
	// }

	return nil
}
