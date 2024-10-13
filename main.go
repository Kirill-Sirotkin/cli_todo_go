package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const ERROR_MSG = "wrong command usage; to see available command usage, type \"help\""
const DATA_FILE_PATH = "todos.json"

var todos = TodosMap{
	Todos: make(map[int]Todo),
}

func main() {
	fmt.Println("------ TODO CLI APP ------")
	fmt.Println("type the \"help\" command to see available commands")

	var err error
	err = LoadData(&todos, DATA_FILE_PATH)
	if err != nil {
		fmt.Printf("%v\ncreating a new save file\n", err)
		SaveData(todos, DATA_FILE_PATH)
	} else {
		fmt.Println("previous data loaded successfully")
	}
	displayTodos(todos)

	var arg string
	for {
		_, err = fmt.Scanln(&arg)
		if err != nil {
			fmt.Println(err)
		}

		quit, err := handleCommandInput(&arg)

		if err != nil {
			fmt.Println(err)
		}

		if quit {
			break
		}

		arg = ""
	}
}

func handleCommandInput(arg *string) (bool, error) {
	if arg == nil {
		return false, errors.New(ERROR_MSG)
	}

	argVal := *arg
	if argVal == "" {
		return false, errors.New(ERROR_MSG)
	}

	switch argVal {
	case "quit":
		return true, nil
	case "q":
		return true, nil
	case "help":
		displayAvailableCommands()
	case "h":
		displayAvailableCommands()
	case "add":
		err := handleAddTodo()
		if err != nil {
			return false, err
		}
	case "edit":
		err := handleEditTodo()
		if err != nil {
			return false, err
		}
	case "toggle":
		err := handleToggleTodo()
		if err != nil {
			return false, err
		}
	case "delete":
		err := handleDeleteTodo()
		if err != nil {
			return false, err
		}
	case "show":
		displayTodos(todos)
	default:
		return false, errors.New(ERROR_MSG)
	}
	err := SaveData(todos, DATA_FILE_PATH)
	if err != nil {
		return false, err
	}

	return false, nil
}

func handleAddTodo() error {
	scanner := bufio.NewReader(os.Stdin)

	fmt.Printf("enter the title of the todo: ")
	title, err := scanner.ReadString('\n')
	if err != nil {
		return err
	}
	title = strings.TrimSpace(title)

	fmt.Printf("enter the description of the todo: ")
	desc, err := scanner.ReadString('\n')
	if err != nil {
		return err
	}
	desc = strings.TrimSpace(desc)

	todos.addTodo(title, desc)
	return nil
}

func handleDeleteTodo() error {
	scanner := bufio.NewReader(os.Stdin)

	fmt.Printf("enter the ID of the todo to delete: ")
	id, err := scanner.ReadString('\n')
	if err != nil {
		return err
	}
	id = strings.TrimSpace(id)

	idInt, err := strconv.Atoi(id)
	if err != nil {
		return err
	}
	err = todos.deleteTodo(idInt)
	if err != nil {
		return err
	}

	return nil
}

func handleEditTodo() error {
	scanner := bufio.NewReader(os.Stdin)

	fmt.Printf("enter the ID of the todo to delete: ")
	id, err := scanner.ReadString('\n')
	if err != nil {
		return err
	}
	id = strings.TrimSpace(id)

	idInt, err := strconv.Atoi(id)
	if err != nil {
		return err
	}

	fmt.Printf("enter the new title (empty field will leave the title unchanged): ")
	input, err := scanner.ReadString('\n')
	if err != nil {
		return err
	}
	input = strings.TrimSpace(input)
	err = todos.editTodoTitle(idInt, input)
	if err != nil {
		return err
	}

	fmt.Printf("enter the new description (empty field will leave the description unchanged): ")
	input, err = scanner.ReadString('\n')
	if err != nil {
		return err
	}
	input = strings.TrimSpace(input)
	err = todos.editTodoDescription(idInt, input)
	if err != nil {
		return err
	}

	return nil
}

func handleToggleTodo() error {
	scanner := bufio.NewReader(os.Stdin)

	fmt.Printf("enter the ID of the todo to delete: ")
	id, err := scanner.ReadString('\n')
	if err != nil {
		return err
	}
	id = strings.TrimSpace(id)

	idInt, err := strconv.Atoi(id)
	if err != nil {
		return err
	}
	err = todos.toggleTodo(idInt)
	if err != nil {
		return err
	}

	return nil
}
