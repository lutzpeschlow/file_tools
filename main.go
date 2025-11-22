package main

import (
	// standard libraries
	"fmt"
	"os"
	"runtime"

	// --- separator - standard vs. own ---
	"github.com/lutzpeschlow/file_tools/ctrl"
	"github.com/lutzpeschlow/file_tools/scaledown"
	"github.com/lutzpeschlow/file_tools/sizing"
)

// ============================================================================
//
// main function
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
	ctrl.DebugPrintoutCtrlObj(&ctrl_obj)

	// (1) ACTION
	//
	// (1.1) SIZING
	if ctrl_obj.Action == "SIZING" {
		fmt.Print(("SIZING ... \n"))
		err_siz := sizing.GetFileList(&ctrl_obj, &file_obj)
		if err_siz != nil {
			fmt.Print(err_siz, "\n")
		}
		for i, file := range file_obj.Files {
			if i <= ctrl_obj.Num {
				fmt.Print(" ", file.Size, " ", file.FileName, " ", i, " \n")
			}
		}
	}
	// (1.2) SCALEDOWN
	if ctrl_obj.Action == "SCALEDOWN" {
		fmt.Print(("SIZING ... \n"))
		err_scale := scaledown.ScaleDown(&ctrl_obj)
		if err_scale != nil {
			fmt.Print("err_scale: ", err_scale, "\n")
		}

	}
	// // Fyne-Aufruf nur hier innerhalb der if-Bedingung
	// myApp := app.New()
	// fmt.Print(myApp)
	// myWindow := myApp.NewWindow("Sizing Results")
	// //
	// // data := []string{}
	// // for i, file := range file_obj.Files {
	// // 	if i <= ctrl_obj.Num {
	// // 		data = append(data, fmt.Sprintf("%d: %s (%s)", i, file.FileName, file.Size))
	// // 	}
	// // }
	// //
	// // // list := widget.NewList(
	// // // 	func() int {
	// // // 		return len(data)
	// // // 	},
	// // // 	func() fyne.CanvasObject {
	// // // 		return widget.NewLabel("template")
	// // // 	},
	// // // 	func(i widget.ListItemID, o fyne.CanvasObject) {
	// // // 		o.(*widget.Label).SetText(data[i])
	// // // 	},
	// // // )
	// //
	// // myWindow.SetContent(container.NewBorder(nil, nil, nil, nil, list))
	// myWindow.Resize(fyne.NewSize(400, 400))
	// myWindow.ShowAndRun()
} //

// =====================================================
