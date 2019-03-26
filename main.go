package main

import (
	"log"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
)

func main() {
	a := app.New()

	w := a.NewWindow("text pad")
	w.Resize(fyne.NewSize(640, 480))
	le := widget.NewMultiLineEntry()
	le.SetPlaceHolder("Please input your words!")
	bt := widget.NewButton("Save", func() {
		log.Println("The save button was pressed!")
	})

	box := fyne.NewContainerWithLayout(
		layout.NewBorderLayout(nil, bt, nil, nil),
		le, bt,
	)
	w.SetContent(box)

	w.ShowAndRun()
}
