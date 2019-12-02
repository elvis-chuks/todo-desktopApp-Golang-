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
		// instead of button, create a widget with a label and button
		btn := widget.NewButton("Action",nil)
		taskWidget := widget.NewVBox(
			widget.NewLabel(fmt.Sprintf("Task : %v",i.Id)),
			widget.NewLabel(fmt.Sprintf("Name : %v",i.Name)),
			widget.NewLabel(fmt.Sprintf("Description : %v",i.Description)),
			btn,
		)
		

		// btn.Resize(fyne.NewSize(200,300))

		list.Append(taskWidget)
	}
	scroll := widget.NewScrollContainer(list)

	scroll.Resize(fyne.NewSize(600, 400))
	TaskWid := fyne.NewContainer(
		scroll,
	)
	container := widget.NewVBox(
		widget.NewLabel("tasks"),
		
		TaskWid,
	)
	return container
}
