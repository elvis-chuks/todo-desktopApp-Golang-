package ui

import (
	"../"
	"fyne.io/fyne/widget"
	"fyne.io/fyne"
	"fyne.io/fyne/dialog"
	"fmt"
	"strconv"
	"../models"
)



func NewTask(w fyne.Window) fyne.Widget{
	taskName := widget.NewEntry()
	taskName.SetPlaceHolder("Task Name")
	description := widget.NewMultiLineEntry()
	description.SetPlaceHolder("task description")
	mainList := handlers.GetItems()
	subList := mainList.Todo
	var listy []int

	for _,i := range subList {
		num,_ := strconv.Atoi(i.Id)

		listy = append(listy,num)
	}
	id := handlers.Biggest(listy)
	// fmt.Println(id)
	

	btn := widget.NewButton("Add Task",func(){
		// get item []models.TodoItem
		// create newItem models.TodoItem
		// pass them to handlers.AddItem
	
		// create variable with type models.TodoItem

		var newItem models.TodoItem

		newItem.Id = id
		newItem.Name = taskName.Text
		newItem.Description = description.Text

		fmt.Println(newItem)
		handlers.AddItem(subList,newItem)
		dialog.ShowInformation("Success","Task added succesfully",w)
		// Task()
	},)

	// form := fyne.NewContainerWithLayout(
	// 	layout.NewGridLayout(1),
	// 	taskName,
	// 	description,
	// 	btn,
	// )
	container := widget.NewVBox(
		// form,
		taskName,
		description,
		btn,

	)
	return container
}