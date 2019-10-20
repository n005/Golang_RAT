package screens

import (
	"fmt"

	"fyne.io/fyne"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
)

func makeFormTab() fyne.Widget {
	ip := widget.NewEntry()
	ip.SetPlaceHolder("192.168.1.20")
	port := widget.NewEntry()
	port.SetPlaceHolder("8080")
	os := widget.NewSelect([]string{"Linux", "Windows"}, func(s string) { fmt.Println("selected", s) })

	form := &widget.Form{
		OnSubmit: func() {
			fmt.Println("Form submitted")
			fmt.Println("IP:", ip.Text)
			fmt.Println("Port:", port.Text)
		},
	}
	form.Append("IP", ip)
	form.Append("Port", port)
	form.Append("OS", os)

	return form
}

func Build() fyne.CanvasObject {
	return widget.NewVBox(
		widget.NewLabelWithStyle("Build Server", fyne.TextAlignCenter, fyne.TextStyle{Bold: true}),
		makeFormTab(),
		layout.NewSpacer(),
	)
}
