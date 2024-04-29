package api

import (
	"github/Rarik88/go_final_project/pkg/hub"
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

type Api struct {
	Task
}

func NewApi(hub hub.Task) *Api {
	return &Api{
		Task: NewApiTask(hub),
	}
}
