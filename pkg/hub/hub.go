package hub

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"github/Rarik88/go_final_project/pkg/model"
)

type Task interface {
	AddTask(task model.Task) (int64, error)
	TaskByID(id string) (model.Task, error)
	Tasks() (model.ListTasks, error)
	UpdateTask(task model.Task) error
	TaskDone(id string) error
	TaskDelete(id string) error
}

type Hub struct {
	Task
}

func NewHub(db *sqlx.DB) *Hub {
	return &Hub{
		Task: NewDB(db),
	}
}
