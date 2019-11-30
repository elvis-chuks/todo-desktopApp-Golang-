package models

type TodoList struct {
	Todo []TodoItem `json:"todo"`
}

type TodoItem struct {
	Id 			string `json:"id"`
	Name 		string  `json:"name"`
	Description string `json:"description"`
}