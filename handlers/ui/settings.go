package ui

import (
	// "fyne.io/fyne/app"
	"fyne.io/fyne/widget"
	"fyne.io/fyne"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/layout"
)

func Settings() fyne.Widget {
	themeSettings := []string{"Light Mode"}
	radio := widget.NewRadio(
		themeSettings,
		func(s string){
			if s == "Light Mode" {
				fyne.CurrentApp().Settings().SetTheme(theme.LightTheme())
			}else{
				fyne.CurrentApp().Settings().SetTheme(theme.DarkTheme())
			}
		},
	)

	container := widget.NewVBox(
		fyne.NewContainerWithLayout(
			layout.NewGridLayout(2),
			radio,
		),
	)
	return container
}