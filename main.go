package main

import (
	"fmt"
	"os"
	"encoding/json"
	"io/ioutil"
	"strconv"
	"./handlers"
	"./handlers/models"
	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"
	"fyne.io/fyne"
	// "fyne.io/fyne/theme"
)



func main(){
	// call the gui
	a := app.New()
	w := a.NewWindow("Todo App")
	w.Resize(fyne.NewSize(600,600))


	w.SetContent(widget.NewVBox(
		widget.NewLabel("todo app"),
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