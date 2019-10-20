package screens

import (
	"fmt"

	"fyne.io/fyne"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
)

var pport = string("")

func portchoose() fyne.Widget {
	port := widget.NewEntry()
	port.SetPlaceHolder("8080")
	botton := widget.NewButton("Set Port", func() {
		fmt.Println("Port:", port.Text)
		pport = ("Port Set " + port.Text)
	})

	form := widget.NewForm()

	form.Append("Port", port)
	form.Append("", botton)

	return form
}

func Listener() fyne.CanvasObject {
	return widget.NewVBox(
		widget.NewLabelWithStyle("Port Setting", fyne.TextAlignCenter, fyne.TextStyle{Bold: true}),
		portchoose(),
		layout.NewSpacer(),
		widget.NewLabelWithStyle(pport, fyne.TextAlignCenter, fyne.TextStyle{Bold: true}),
		layout.NewSpacer(),
	)
}
