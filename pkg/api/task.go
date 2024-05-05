package api

import (
	"github/Rarik88/go_final_project/pkg/hub"
	"github/Rarik88/go_final_project/pkg/model"
)

type ApiTask struct {
	hub hub.Task
}

func NewApiTask(hub hub.Task) *ApiTask {
	return &ApiTask{
		hub: hub,
	}
}

func (a *ApiTask) AddTask(task model.Task) (int64, error) {
	return a.hub.AddTask(task)
}

func (a *ApiTask) TaskByID(id string) (model.Task, error) {
	return a.hub.TaskByID(id)
}

// TasksHandler обрабатывает GET-запросы для получения списка ближайших задач.
func (a *ApiTask) Tasks() (model.ListTasks, error) {
	return a.hub.Tasks()
}

// UpdateTaskHandler обрабатывает PUT-запросы для обновления параметров задачи по её идентификатору.
func (a *ApiTask) UpdateTask(task model.Task) error {
	return a.hub.UpdateTask(task)
}

// TaskDoneHandler обрабатывает POST-запросы для отметки задачи как выполненной.
func (a *ApiTask) TaskDone(id string) error {
	return a.hub.TaskDone(id)
}

// TaskDeleteHandler обрабатывает DELETE-запросы для удаления задачи.
func (a *ApiTask) TaskDelete(id string) error {
	return a.hub.TaskDelete(id)
}
