package main

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"

	// "github.com/lutzpeschlow/file_tools/ctrl"
	"github.com/lutzpeschlow/file_tools/ctrl"
)

// ============================================================================

func main() {
	// instance of control object
	ctrl_obj := ctrl.Control_Object{}
	// check operating system
	osName := runtime.GOOS
	fmt.Print("operating system: ", osName, "\n")

	// (0) CONTROL
	// set settings via control file
	err_ctrl := ctrl.ReadControlFile("control.txt", &ctrl_obj, osName)
	if err_ctrl != nil {
		fmt.Printf(" %v\n", err_ctrl)
		os.Exit(1)
	}
	// content of control object
	fmt.Print("Settings: ", ctrl_obj.Action, " ", ctrl_obj.Dir, " ",
		ctrl_obj.Num, " ", ctrl_obj.View, "\n")

	// short test of file path walk
	root := "."
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			fmt.Println(path)
		}
		return nil
	})
	// error value
	fmt.Print("finalize: ", err)
}
