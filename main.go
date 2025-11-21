package main

import (
	"fmt"

	"os"
	"runtime"

	// separator - standard vs. own
	"github.com/lutzpeschlow/file_tools/ctrl"
	"github.com/lutzpeschlow/file_tools/scaledown"
	"github.com/lutzpeschlow/file_tools/sizing"
)

// ============================================================================

func main() {
	// instance of control object
	ctrl_obj := ctrl.Control_Object{}
	file_obj := sizing.FileList{}
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

	if ctrl_obj.Action == "SIZING" {
		err_siz := sizing.GetFileList(&ctrl_obj, &file_obj)
		fmt.Print("finalize: ", err_siz)
	}

	if ctrl_obj.Action == "SCALEDOWN" {
		err_scale := scaledown.ScaleDown(&ctrl_obj)
		fmt.Print("finalize: ", err_scale)

	}

	// error value

}

// =====================================================

// convert input.jpg -resize 2048x2048\> -strip -quality 70 output.jpg
