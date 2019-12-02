package main

import (
	"fmt"
	"os"
	"encoding/json"
	"io/ioutil"
	"strconv"
	"./handlers"
	"./handlers/models"
	"./handlers/ui"
	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"
	"fyne.io/fyne"
	"fyne.io/fyne/theme"
	// "fyne.io/fyne/canvas"
	"fyne.io/fyne/layout"
	// "fyne.io/fyne"
	// "image/color"
)

func test() fyne.Widget{
	return widget.NewVBox(
		widget.NewLabel("todo"),
	)
}

func main(){
	// call the gui
	a := app.New()
	w := a.NewWindow("Todo App")
	w.Resize(fyne.NewSize(600,600))
	
	tabs := widget.NewTabContainer(
		widget.NewTabItemWithIcon("Tasks",theme.ContentPasteIcon(),ui.Tasks()),
		widget.NewTabItemWithIcon("New Task",theme.FolderNewIcon(),ui.NewTask(w)),
		widget.NewTabItemWithIcon("Settings",theme.SettingsIcon(),ui.Settings()),	
	)

	w.SetContent(widget.NewVBox(
		tabs,
		layout.NewSpacer(),
		// fyne.NewContainerWithLayout(
		// 	layout.NewFixedGridLayout(fyne.NewSize(600,90)),
			
		// 	&canvas.Rectangle{
		// 		FillColor: color.RGBA{0x80, 167, 0, 0xff},
		// 		StrokeColor: color.RGBA{0xff, 0xff, 0xff, 0xff},
		// 		StrokeWidth: 1,
		// 	},
		// ),
		
	))

	jsonFile,err := os.Open("data.json")

	if err != nil{

		fmt.Println(err)

	}

	defer jsonFile.Close()

	byteVal, _ := ioutil.ReadAll(jsonFile)

	var todo models.TodoList

	json.Unmarshal(byteVal,&todo)


	var listy []int

	for _,i := range todo.Todo {
		num,_ := strconv.Atoi(i.Id)

		listy = append(listy,num)
	}

	egg := handlers.Biggest(listy)

	fmt.Println(egg)

	w.CenterOnScreen()
	w.ShowAndRun()
}