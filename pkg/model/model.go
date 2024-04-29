package model

type NextDate struct {
	Now    string `form:"now"`
	Date   string `form:"date"`
	Repeat string `form:"repeat"`
}

type Task struct {
	ID      string `json:"id" db:"id"`
	Date    string `json:"date" db:"date"`
	Title   string `json:"title" db:"title"`
	Comment string `json:"comment" db:"comment"`
	Repeat  string `json:"repeat" db:"repeat"`
}

type ListTasks struct {
	Tasks []Task `json:"tasks"`
}
