package main

// -----
// Necessary structs and their methods
// -----

import (
	"errors"
	"strconv"
	"time"
)

const INDEX_INVALID_ERROR = "error: todo with the given index does not exist. id: "

type Todo struct {
	Id          int
	Title       string
	Description string
	Completed   bool
	StartedAt   time.Time
	CompletedAt *time.Time
}

type TodosMap struct {
	IdCounter int
	Todos     map[int]Todo
}

func (tm *TodosMap) addTodo(tl string, desc string) {
	todo := Todo{
		Id:          tm.IdCounter,
		Title:       tl,
		Description: desc,
		StartedAt:   time.Now(),
	}
	tm.Todos[tm.IdCounter] = todo
	tm.IdCounter++
}

func (tm *TodosMap) validateIndex(id int) bool {
	_, ok := tm.Todos[id]
	return ok
}

func (tm *TodosMap) deleteTodo(id int) error {
	if !tm.validateIndex(id) {
		return errors.New(INDEX_INVALID_ERROR + strconv.Itoa(id))
	}
	delete(tm.Todos, id)
	return nil
}

func (tm *TodosMap) editTodoTitle(id int, tl string) error {
	if !tm.validateIndex(id) {
		return errors.New(INDEX_INVALID_ERROR + strconv.Itoa(id))
	}
	if tl == "" {
		return nil
	}
	todo := tm.Todos[id]
	todo.Title = tl
	tm.Todos[id] = todo
	return nil
}

func (tm *TodosMap) editTodoDescription(id int, desc string) error {
	if !tm.validateIndex(id) {
		return errors.New(INDEX_INVALID_ERROR + strconv.Itoa(id))
	}
	if desc == "" {
		return nil
	}
	todo := tm.Todos[id]
	todo.Description = desc
	tm.Todos[id] = todo
	return nil
}

func (tm *TodosMap) toggleTodo(id int) error {
	if !tm.validateIndex(id) {
		return errors.New(INDEX_INVALID_ERROR + strconv.Itoa(id))
	}
	todo := tm.Todos[id]
	todo.Completed = !todo.Completed

	if todo.Completed {
		completedTime := time.Now()
		todo.CompletedAt = &completedTime
	} else {
		todo.CompletedAt = nil
	}

	tm.Todos[id] = todo
	return nil
}
