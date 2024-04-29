package hub

import (
	"fmt"
	"github.com/sirupsen/logrus"
	nextDate "github/Rarik88/go_final_project/pkg/data"
	"github/Rarik88/go_final_project/pkg/model"
	"regexp"
	"time"
)

const TaskTable string = "scheduler"

func (t *TaskSQLite) AddTask(task model.Task) (int64, error) {
	err := t.checkTask(&task)
	if err != nil {
		return 0, err
	}

	query := fmt.Sprintf("INSERT INTO %s (title, comment, date, repeat) VALUES ($1, $2, $3, $4) RETURNING id", "scheduler")
	row := t.db.QueryRow(query, task.Title, task.Comment, task.Date, task.Repeat)

	var id int64
	if err = row.Scan(&id); err != nil {
		return 0, err
	}

	logrus.Println("task added, id = ", id)

	return id, nil
}

func (t *TaskSQLite) TaskByID(id string) (model.Task, error) {
	var task model.Task

	stmt, err := t.db.Prepare("SELECT * FROM scheduler WHERE id = ?")
	if err != nil {
		return model.Task{}, err
	}
	err = stmt.QueryRow(id).Scan(&task.ID, &task.Date, &task.Title, &task.Comment, &task.Repeat)
	if err != nil {
		return model.Task{}, err
	}
	return task, nil
}

// TasksHandler обрабатывает GET-запросы для получения списка ближайших задач.
func (t *TaskSQLite) Tasks() (model.ListTasks, error) {
	var tasks []model.Task

	query := fmt.Sprintf("SELECT * FROM %s ORDER BY date LIMIT ?", TaskTable)
	err := t.db.Select(&tasks, query, 10)
	if err != nil {
		return model.ListTasks{}, err
	}

	if len(tasks) == 0 {
		return model.ListTasks{Tasks: []model.Task{}}, nil
	}
	return model.ListTasks{Tasks: tasks}, nil
}

// UpdateTaskHandler обрабатывает PUT-запросы для обновления параметров задачи по её идентификатору.
func (t *TaskSQLite) UpdateTask(task model.Task) error {

	err := t.checkTask(&task)
	if err != nil {
		return err
	}

	stmt, err := t.db.Prepare("UPDATE scheduler SET date = ?, title = ?, comment = ?, repeat = ? WHERE id = ?")
	_, err = stmt.Exec(task.Date, task.Title, task.Comment, task.Repeat, task.ID)
	if err != nil {
		return err
	}
	return nil
}

// TaskDoneHandler обрабатывает POST-запросы для отметки задачи как выполненной.
func (t *TaskSQLite) TaskDone(id string) error {
	task, err := t.TaskByID(id)
	if err != nil {
		return err
	}

	if task.Repeat == "" {
		stmt, err := t.db.Prepare("DELETE FROM scheduler WHERE id = ?")
		if err != nil {
			return err
		}
		_, err = stmt.Exec(id)
		if err != nil {
			return err
		}
		return nil
	}

	nd := model.NextDate{
		Date:   task.Date,
		Now:    time.Now().Format(`20060102`),
		Repeat: task.Repeat,
	}

	newDate, err := nextDate.NextDate(nd)
	if err != nil {
		return err
	}

	task.Date = newDate
	stmt, err := t.db.Prepare("UPDATE scheduler SET date = ? WHERE id = ?")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(task.Date, id)
	if err != nil {
		return err
	}
	return nil
}

// TaskDeleteHandler обрабатывает DELETE-запросы для удаления задачи.
func (t *TaskSQLite) TaskDelete(id string) error {
	_, err := t.TaskByID(id)
	if err != nil {
		return err
	}
	stmt, err := t.db.Prepare("DELETE FROM scheduler WHERE id = ?")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}
	return nil
}
func (t *TaskSQLite) checkTask(task *model.Task) error {
	const (
		Format_yyyymmdd = `20060102`
	)

	if task.Title == "" {
		return fmt.Errorf("отсутствует заголовок задачи")
	}

	if !regexp.MustCompile(`^(d\s.*|y)?$`).MatchString(task.Repeat) {
		return fmt.Errorf("некорректное повторение задачи - %s", task.Repeat)
	}

	now := time.Now().Format(Format_yyyymmdd)

	if task.Date == "" {
		task.Date = now
	}

	_, err := time.Parse(Format_yyyymmdd, task.Date)
	if err != nil {
		return fmt.Errorf("неверный формат даты - %s", task.Date)
	}

	if task.Date < now {
		if task.Repeat == "" {
			task.Date = now
		}
		if task.Repeat != "" {
			nd := model.NextDate{
				Date:   task.Date,
				Now:    now,
				Repeat: task.Repeat,
			}
			task.Date, err = nextDate.NextDate(nd)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
