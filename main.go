package main

import (
	"fmt"
	"log"
	"os"
	"runtime"

	"gioui.org/app"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/widget"
	"gioui.org/widget/material"

	// separator - standard vs. own
	"github.com/lutzpeschlow/file_tools/ctrl"
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

	// short test of file list
	err_siz := sizing.GetFileList(&ctrl_obj, &file_obj)

	if ctrl_obj.View == "text" {
		fmt.Print("text output")
	}

	if ctrl_obj.View == "window" {
		go func() {
			var w app.Window
			w.Option(app.Title("Gio Beispiel"))
			if err := loop(&w); err != nil {
				log.Fatal(err)
			}
			os.Exit(0)
		}()
		app.Main()
	}

	// error value
	fmt.Print("finalize: ", err_siz)
}

// =====================================================

// =====================================================

func loop(w *app.Window) error {
	th := material.NewTheme()
	// th.Shaper = app.NewShaper(gofont.Collection())

	var ops op.Ops
	var list layout.List
	list.Axis = layout.Vertical // wird standardmäßig gesetzt, explizit gut

	items := []string{"Item 1",
		"Item 2", "Item 3", "Item 4", "Item 5", "Item 6", "Item 7", "Item 8", "Item 9",
		"Item 2", "Item 3", "Item 4", "Item 5", "Item 6", "Item 7", "Item 8", "Item 9",
		"Item 2", "Item 3", "Item 4", "Item 5", "Item 6", "Item 7", "Item 8", "Item 9",
		"Item 2", "Item 3", "Item 4", "Item 5", "Item 6", "Item 7", "Item 8", "Item 9",
		"Item 2", "Item 3", "Item 4", "Item 5", "Item 6", "Item 7", "Item 8", "Item 9",
		"Item 2", "Item 3", "Item 4", "Item 5", "Item 6", "Item 7", "Item 8", "Item 9",
		"Item 2", "Item 3", "Item 4", "Item 5", "Item 6", "Item 7", "Item 8", "Item 9",
		"Item 10"}
	btns := make([]widget.Clickable, len(items))

	for {
		e := w.Event()
		switch e := e.(type) {
		case app.DestroyEvent:
			return e.Err
		case app.FrameEvent:
			gtx := app.NewContext(&ops, e)

			// Container mit fester Höhe und Breite, damit Scrollbar sichtbar wird
			layout.Flex{Axis: layout.Vertical}.Layout(gtx,
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					return layout.Dimensions{Size: gtx.Constraints.Min}
				}),
				layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
					// Listbox mit Scrollbar
					return list.Layout(gtx, len(items), func(gtx layout.Context, i int) layout.Dimensions {
						btn := &btns[i]
						if btn.Clicked(gtx) {
							log.Println("Clicked:", items[i])
						}
						return material.Button(th, btn, items[i]).Layout(gtx)
					})
				}),
			)

			e.Frame(gtx.Ops)
		}
	}
}
