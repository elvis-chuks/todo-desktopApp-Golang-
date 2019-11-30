package ui

import (
	"../"
	// "fyne.io/fyne/app"
	"fyne.io/fyne/widget"
	"fyne.io/fyne"
	// "fyne.io/fyne/theme"
	"fmt"
)

func Tasks() fyne.Widget {

	mainList := handlers.GetItems()
	subList := mainList.Todo
	list := widget.NewVBox()

	for _,i := range subList {
		list.Append(widget.NewButton(fmt.Sprintf("%v",i.Id),nil))
	}
	scroll := widget.NewScrollContainer(list)
	scroll.Resize(fyne.NewSize(200, 200))
	container := widget.NewVBox(
		widget.NewLabel("tasks"),
		fyne.NewContainer(
			scroll,
		),
	)
	return container
}
