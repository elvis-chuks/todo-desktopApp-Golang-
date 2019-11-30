package ui

import (
	// "fyne.io/fyne/app"
	"fyne.io/fyne/widget"
	"fyne.io/fyne"
	"fyne.io/fyne/theme"
)

func Settings() fyne.widget {
	theme := []string{"Light Mode"}
	radio := widget.NewRadio(
		theme,
		func(s string){
			if s == "Light Mode" {
				fyne.CurrentApp().Settings().SetTheme(theme.LightTheme())
			}else{
				fyne.CurrentApp().Settings().SetTheme(theme.DarkTheme())
			}
		}
	)