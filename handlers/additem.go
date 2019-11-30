package handlers

import (
	"fmt"
	"os"
	"encoding/json"
	"io/ioutil"
	"./models"
)

func AddItem(item []models.TodoItem,newItem models.TodoItem){
	var list models.TodoList

	item = append(item,newItem)
	
	list.Todo = item

	jsonString,_ := json.Marshal(list)

	ioutil.WriteFile("data.json",jsonString,os.ModePerm)

	fmt.Println("added item")

}