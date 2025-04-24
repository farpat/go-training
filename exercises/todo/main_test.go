package todo

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestItAddsTodo(t *testing.T) {
	// ARRANGE
	input := Todo{
		Title:       "Test Todo",
		Description: "Test Description",
	}

	var todos Todos = make(map[string]Todo)

	// ACT
	err := todos.AddTodo(input)
	fmt.Printf("%v", err)

	// ASSERT
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if len(todos) != 1 {
		t.Errorf("Expected 1 todo, got %d", len(todos))
	}

	firstKey := firstKey(todos)
	if !isEqual(todos[firstKey], input) {
		t.Errorf("Expected todo %v, got %v", input, todos[firstKey])
	}
}

func TestItRemovesTodoById(t *testing.T) {
	// ARRANGE
	input := Todo{
		Id:          "any id",
		Title:       "Test Todo",
		Description: "Test Description",
		DoneAt:      nil,
	}
	var todos Todos = make(map[string]Todo)
	todos["any id"] = input

	// ACT
	todos.RemoveTodoById(input.Id)

	// ASSERT
	if len(todos) != 0 {
		t.Errorf("Expected 0 todos, got %d", len(todos))
	}
}

func TestItFindsTodosByTitle(t *testing.T) {
	// ARRANGE
	input1 := Todo{
		Id:          "any id",
		Title:       "Test Todo 1",
		Description: "Test Description 1",
		DoneAt:      nil,
	}
	input2 := Todo{
		Id:          "any other id",
		Title:       "Test Todo 2",
		Description: "Test Description 2",
		DoneAt:      nil,
	}
	var todos Todos = make(map[string]Todo)
	todos[input1.Id] = input1
	todos[input2.Id] = input2

	// ACT
	result := todos.FindTodosByTitle("Test Todo")

	// ASSERT
	if len(result) != 2 {
		t.Errorf("Expected 2 todos, got %d", len(result))
	}
}

func TestItSortsTodosByTitle(t *testing.T) {
	// ARRANGE
	input1 := Todo{
		Id:          "any id",
		Title:       "Test Todo 1",
		Description: "Test Description 1",
		DoneAt:      nil,
	}
	input2 := Todo{
		Id:          "any other id",
		Title:       "Test Todo 2",
		Description: "Test Description 2",
		DoneAt:      nil,
	}
	var todos Todos = make(map[string]Todo)
	todos[input1.Id] = input1
	todos[input2.Id] = input2

	// ACT
	sortedTitles := todos.getSortedTitles()

	// ASSERT
	if len(sortedTitles) != 2 {
		t.Errorf("Expected 2 todos, got %d", len(sortedTitles))
	}
	var expectedTitles []string = []string{"Test Todo 1", "Test Todo 2"}
	if diff := cmp.Diff(sortedTitles, expectedTitles); diff != "" {
		t.Errorf("getSortedTitles() = %v diff %v", sortedTitles, diff)
	}
}

func TestItMarksTodoAsDone(t *testing.T) {
	// ARRANGE
	input := Todo{
		Id:          "any id",
		Title:       "Test Todo",
		Description: "Test Description",
		DoneAt:      nil,
	}
	var todos Todos = make(map[string]Todo)
	todos[input.Id] = input

	// ACT
	todos.MarkAsDone(input.Id)

	// ASSERT
	if todos[input.Id].DoneAt == nil {
		t.Errorf("Expected todo to be marked as done, got %v", todos[input.Id])
	}
}

func firstKey(todos Todos) string {
	firstKey := ""
	for key := range todos {
		firstKey = key
		break
	}
	return firstKey
}

func isEqual(todo1 Todo, todo2 Todo) bool {
	return todo1.Title == todo2.Title && todo1.Description == todo2.Description
}
