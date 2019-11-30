package handlers

import (
	"fmt"
	"os"
	"encoding/json"
	"io/ioutil"
	"./models"
)

func GetItems() models.TodoList{
	jsonFile,err := os.Open("data.json")

	if err != nil{

		fmt.Println(err)

	}

	defer jsonFile.Close()

	byteVal, _ := ioutil.ReadAll(jsonFile)

	var todo models.TodoList

	json.Unmarshal(byteVal,&todo)

	return todo
}