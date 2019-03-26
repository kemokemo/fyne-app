package main

import (
	"io/ioutil"
	"log"
	"os"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
	"github.com/sqweek/dialog"
)

func main() {
	a := app.New()

	w := a.NewWindow("text pad")
	w.Resize(fyne.NewSize(640, 480))
	le := widget.NewMultiLineEntry()
	le.SetPlaceHolder("Please input your words!")
	bt := widget.NewButton("Save as...", func() {
		filename, err := dialog.File().Filter("Text files", "txt").Title("Save as a text file").Save()
		if err != nil {
			log.Println(err)
			return
		}
		err = ioutil.WriteFile(filename, []byte(le.Text), os.ModePerm)
		if err != nil {
			log.Println(err)
			return
		}
	})

	box := fyne.NewContainerWithLayout(
		layout.NewBorderLayout(nil, bt, nil, nil),
		le, bt,
	)
	w.SetContent(box)

	w.ShowAndRun()
}
