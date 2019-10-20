package screens

import (
	"fmt"

	"fyne.io/fyne"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
)

func makeFormTabRAT() fyne.Widget {
	ip := widget.NewEntry()
	ip.SetPlaceHolder("192.168.1.20")
	port := widget.NewEntry()
	port.SetPlaceHolder("8080")

	form := &widget.Form{
		OnSubmit: func() {
			fmt.Println("Form submitted")
			fmt.Println("IP:", ip.Text)
			fmt.Println("Port:", port.Text)
		},
	}
	form.Append("IP", ip)
	form.Append("Port", port)

	return form
}

func Rat() fyne.CanvasObject {
	return widget.NewVBox(
		widget.NewLabelWithStyle("Rat Lists:", fyne.TextAlignCenter, fyne.TextStyle{Bold: true}),
		makeFormTabRAT(),
		layout.NewSpacer(),
	)
}
