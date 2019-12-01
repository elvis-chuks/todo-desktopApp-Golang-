package ui

import (
	"../"
	"fyne.io/fyne/widget"
	"fyne.io/fyne"
	// "fyne.io/fyne/layout"
	"fmt"
)

func Tasks() fyne.Widget {

	mainList := handlers.GetItems()
	subList := mainList.Todo
	list := widget.NewVBox()
	// write function to open a dialog for setting tasks status
	for _,i := range subList {
		btn := widget.NewButton(fmt.Sprintf("ID : %v || Name : %v  ||  Description : %v",i.Id,i.Name,i.Description),nil)

		btn.Resize(fyne.NewSize(200,300))

		list.Append(btn)
	}
	scroll := widget.NewScrollContainer(list)

	scroll.Resize(fyne.NewSize(400, 400))

	container := widget.NewVBox(
		widget.NewLabel("tasks"),
		
		fyne.NewContainer(
			
			scroll,
		),
	)
	return container
}
