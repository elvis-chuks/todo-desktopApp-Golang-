package handlers

import (
	"fmt"
	"os"
	"encoding/json"
	"io/ioutil"
	"./models"
)

func RemoveItem(item []models.TodoItem,id string) {
	for index,i := range item {
		if i.Id == id {
			item = append(item[:index], item[index+1:]...)
		}
	}

	var list models.TodoList

	list.Todo = item

	jsonString,_ := json.Marshal(list)

	ioutil.WriteFile("data.json",jsonString,os.ModePerm)

	fmt.Println("removed item")
}