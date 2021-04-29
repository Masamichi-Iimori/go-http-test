package infrastructure

import (
	"fmt"

	"github.com/masamichhhhi/go-clean/domain/model"
	"github.com/masamichhhhi/go-clean/domain/repository"
)

// type TodoRepository struct {
// 	SqlHandler
// }

type TodoRepository struct {
	GormHandler
}

func NewTodoRepository(sqlHandler GormHandler) repository.TodoRepository {
	todoRepository := TodoRepository{sqlHandler}
	return &todoRepository
}

// func (todoRepo *TodoRepository) FindAll() (todos []*model.Todo, err error) {
// 	rows, err := todoRepo.GormHandler.Conn.Query("SELECT id, task, limitDate, status FROM go_sample.todos")
// 	defer rows.Close()
// 	if err != nil {
// 		fmt.Print(err)
// 		return
// 	}
// 	for rows.Next() {
// 		todo := model.Todo{}
// 		rows.Scan(&todo.ID, &todo.Task, &todo.LimitDate, &todo.Status)

// 		todos = append(todos, &todo)
// 	}
// 	return
// }
func (todoRepo *TodoRepository) FindAll() (todos []*model.Todo, err error) {
	err = todoRepo.GormHandler.Conn.Find(&todos).Error

	if err != nil {
		fmt.Print(err)
		return
	}

	return
}

// func (todoRepo *TodoRepository) Find(word string) (todos []*model.Todo, err error) {
// 	rows, err := todoRepo.SqlHandler.Conn.Query("SELECT * FROM todos WHERE task LIKE ?", "%"+word+"%")
// 	defer rows.Close()
// 	if err != nil {
// 		fmt.Print(err)
// 		return
// 	}
// 	for rows.Next() {
// 		todo := model.Todo{}

// 		rows.Scan(&todo.ID, &todo.Task, &todo.LimitDate, &todo.Status)

// 		todos = append(todos, &todo)
// 	}
// 	return
// }

func (todoRepo *TodoRepository) Find(word string) (todos []*model.Todo, err error) {
	err = todoRepo.GormHandler.Conn.Raw("SELECT * FROM todos WHERE task LIKE ?", "%"+word+"%").Scan(todos).Error
	if err != nil {
		fmt.Print(err)
		return
	}

	return
}

// func (todoRepo *TodoRepository) Create(todo *model.Todo) (*model.Todo, error) {
// 	_, err := todoRepo.SqlHandler.Conn.Exec("INSERT INTO todos (task,limitDate,status) VALUES (?, ?, ?) ", todo.Task, todo.LimitDate, todo.Status)
// 	return todo, err
// }

func (todoRepo *TodoRepository) Create(todo *model.Todo) (*model.Todo, error) {
	err := todoRepo.GormHandler.Conn.Create(&todo).Error
	return todo, err
}

// func (todoRepo *TodoRepository) Update(todo *model.Todo) (*model.Todo, error) {
// 	_, err := todoRepo.SqlHandler.Conn.Exec("UPDATE todos SET task = ?,limitDate = ? ,status = ? WHERE id = ?", todo.Task, todo.LimitDate, todo.Status, todo.ID)
// 	return todo, err
// }

func (todoRepo *TodoRepository) Update(todo *model.Todo) (*model.Todo, error) {
	var targetTodo *model.Todo
	todoRepo.GormHandler.Conn.Where("id = ?", todo.ID).First(&targetTodo)
	targetTodo.LimitDate = todo.LimitDate
	targetTodo.Status = todo.Status
	targetTodo.Task = todo.Task

	err := todoRepo.GormHandler.Conn.Save(&targetTodo).Error
	return targetTodo, err
}
