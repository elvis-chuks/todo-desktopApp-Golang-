package main

import (
	"fmt"
	"encoding/json"
	"io/ioutil"
	"os"
	"reflect"
)

type TodoList struct {
	Todo []TodoItem `json:"todo"`
}

type TodoItem struct {
	Id 			string `json:"id"`
	Name 		string  `json:"name"`
	Description string `json:"description"`
}

func additem(id string,name string,description string,items []TodoItem){
	var item TodoItem
	item.Id = id
	item.Name = name
	item.Description = description

	jsonFile,err := os.Open("data.json")

	if err != nil{

		fmt.Println(err)

	}

	defer jsonFile.Close()

	// get the already existing data and 
	items = append(items,item)
	var list TodoList
	list.Todo = items
	jsonString,_ := json.Marshal(list.Todo)
	ioutil.WriteFile("data.json",jsonString,os.ModePerm)
}

func removeitem(item []TodoItem) []TodoItem{
	fmt.Println(len(item),item)
	for index,i := range item {
		// i = i.()
		if i.Id == "3"{
			fmt.Println(i,index)
			item = append(item[:index], item[index+1:]...)
		}
		
	}
	return item

}

type Stuff map[string]string

func main(){
	fmt.Println("Hello world")

	jsonFile,err := os.Open("data.json")

	if err != nil{

		fmt.Println(err)

	}

	defer jsonFile.Close()

	byteVal, _ := ioutil.ReadAll(jsonFile)

	

	
	newL := TodoItem{Id:"4",Name:"todo 4",Description:"the forth"}
	// neww := TodoItem{Id:"4"}
	var todo TodoList
	// var stuffy Stuff
	// json.Unmarshal(byteVal,&stuffy)
	json.Unmarshal(byteVal, &todo)

	// fmt.Println(newL)
	fmt.Println(reflect.TypeOf(todo.Todo))
	todo.Todo = append(todo.Todo,newL)
	yo:=removeitem(todo.Todo)
	fmt.Println(yo)
	// for index,i := range todo.Todo {
	// 	// i = i.()
	// 	if i.Id == "1"{
	// 		fmt.Println(i,index)
	// 	}
	// }
	// to := stuffy["todo"]
	// fmt.Println(to)
	


	// jsonString,_ := json.Marshal(todo)
	// ioutil.WriteFile("data.json",jsonString,os.ModePerm)
	// write file
}