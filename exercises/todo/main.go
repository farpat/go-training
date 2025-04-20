package todo

import (
	"sort"
	"strings"
	"time"

	"github.com/google/uuid"
)

type Todo struct {
	Id          string
	Title       string
	Description string
	DoneAt      *time.Time
}

type Todos map[string]Todo

type TodoManager interface {
	AddTodo(todo Todo) error
	RemoveTodoById(id string)
	FindTodosByTitle(title string) Todos
	SortByTitle() Todos
	MarkAsDone(id string)
}

func (todos *Todos) AddTodo(newTodo Todo) error {
	newTodo.Id = uuid.New().String()
	newTodo.DoneAt = nil
	(*todos)[newTodo.Id] = newTodo

	return nil
}

func (todos *Todos) RemoveTodoById(id string) {
	delete(*todos, id)
}

func (todos *Todos) FindTodosByTitle(title string) Todos {
	var foundedTodos Todos = make(map[string]Todo)

	for _, todo := range *todos {
		if strings.Contains(todo.Title, title) {
			foundedTodos[todo.Id] = todo
		}
	}

	return foundedTodos
}

func (todos *Todos) getSortedTitles() []string {
	var sortedTitles []string

	for _, todo := range *todos {
		sortedTitles = append(sortedTitles, todo.Title)
	}

	sort.Slice(sortedTitles, func(i, j int) bool {
		return sortedTitles[i] < sortedTitles[j]
	})

	return sortedTitles
}

func (todos *Todos) MarkAsDone(id string) {
	todo := (*todos)[id]
	now := time.Now()
	todo.DoneAt = &now
	(*todos)[id] = todo
}
