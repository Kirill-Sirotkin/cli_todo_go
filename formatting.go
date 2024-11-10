package main

// -----
// Functions to visually display the todos and commands
// -----

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
	"time"
)

// Mostly arbitrary constant values for formatting spaces and separators
const TITLE_LENGTH_LIMIT int = 20
const DESC_LENGTH_LIMIT int = 30
const COMPLETE_FLAG_LENGTH_LIMIT int = 10
const DATE_LENGTH_LIMIT int = 21
const SEPARATORS_COUNT int = 17
const ID = "id"
const TITLE = "title"
const DESC = "description"
const STATUS = "completed?"
const TIME_START = "time started"
const TIME_END = "time finished"

func displayAvailableCommands() {
	fmt.Printf("\"help\" or \"h\"       - see list of available commands\n")
	fmt.Printf("\"quit\" or \"q\"       - close the app\n")
	fmt.Printf("\"show\"               - show full todo list\n")
	fmt.Printf("\"add\"                - add a new todo\n")
	fmt.Printf("\"delete\"             - delete a todo with the given id\n")
	fmt.Printf("\"toggle\"             - toggle a todo completion status\n")
	fmt.Printf("\"edit\"               - edit the title or description of a todo\n")
	fmt.Printf("The list of todos is saved automatically after every executed command\n")
}

func displayTodos(tm TodosMap) {
	maxIdLength := max(getMaxIdLength(tm), len(ID))
	totalLength := maxIdLength +
		TITLE_LENGTH_LIMIT +
		DESC_LENGTH_LIMIT +
		COMPLETE_FLAG_LENGTH_LIMIT +
		DATE_LENGTH_LIMIT*2 +
		SEPARATORS_COUNT

	keys := make([]int, 0)
	for k := range tm.Todos {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	fmt.Println("")
	fmt.Printf("%v\n", strings.Repeat("-", max(0, totalLength)))

	if len(tm.Todos) == 0 {
		fmt.Println("No todos to display")
		fmt.Printf("%v\n", strings.Repeat("-", max(0, totalLength)))
		return
	}

	formatHeader(maxIdLength)
	fmt.Printf("%v\n", strings.Repeat("-", max(0, totalLength)))

	for _, k := range keys {
		fmt.Println(formatBlock(tm.Todos[k], maxIdLength, totalLength))
	}
}

func formatHeader(idLength int) {
	idString := ID + strings.Repeat(" ", max(0, idLength-len(ID)))
	completeFlagString := STATUS + strings.Repeat(" ", max(0, COMPLETE_FLAG_LENGTH_LIMIT-len(STATUS)))
	startedDateString := TIME_START + strings.Repeat(" ", max(0, DATE_LENGTH_LIMIT-len(TIME_START)))
	completedDateString := TIME_END + strings.Repeat(" ", max(0, DATE_LENGTH_LIMIT-len(TIME_END)))
	titleString := TITLE + strings.Repeat(" ", max(0, TITLE_LENGTH_LIMIT-len(TITLE)))
	descString := DESC + strings.Repeat(" ", max(0, DESC_LENGTH_LIMIT-len(DESC)))

	fmt.Printf("%v", getLineWithSeparators(idString, titleString, descString, completeFlagString, startedDateString, completedDateString))
}

func formatBlock(t Todo, idLength int, totalLength int) string {
	titleLineBreaksCount := len(t.Title) / TITLE_LENGTH_LIMIT
	descLineBreaksCount := len(t.Description) / DESC_LENGTH_LIMIT
	var resultLine string
	var tempBlock string

	// Sometimes a description or title can be longer than the character limit per line.
	// In this case several lines are needed to fit the data.
	// The block needs to be constructed from several rows, according to the maximum line breaks.
	for i := 0; i <= max(titleLineBreaksCount, descLineBreaksCount); i++ {
		var idString string
		var completeFlagString string
		var startedDateString string
		var completedDateString string
		var titleString string
		var descString string

		// Title and description do not have a fixed size.
		// Only a substring that fits into the character limit is taken from the whole string per line.
		if i <= titleLineBreaksCount {
			titleString = t.Title[i*TITLE_LENGTH_LIMIT : min((i+1)*TITLE_LENGTH_LIMIT, len(t.Title))]
		}
		if i <= descLineBreaksCount {
			descString = t.Description[i*DESC_LENGTH_LIMIT : min((i+1)*DESC_LENGTH_LIMIT, len(t.Description))]
		}

		// Id, Complete Flag, Start Date and Complete Date all have a fixed size.
		// They will only exist in the first row of the block.
		if i == 0 {
			idString = strconv.Itoa(t.Id)
			completeFlagString = strconv.FormatBool(t.Completed)
			startedDateString = t.StartedAt.Format(time.RFC822)
			if t.CompletedAt != nil {
				completedDateString = t.CompletedAt.Format(time.RFC822)
			} else {
				completedDateString = "in progress..."
			}
		}

		// All fields need to be filled with spaces up to the character limit per column.
		idString += strings.Repeat(" ", max(0, idLength-len(idString)))
		completeFlagString += strings.Repeat(" ", max(0, COMPLETE_FLAG_LENGTH_LIMIT-len(completeFlagString)))
		startedDateString += strings.Repeat(" ", max(0, DATE_LENGTH_LIMIT-len(startedDateString)))
		completedDateString += strings.Repeat(" ", max(0, DATE_LENGTH_LIMIT-len(completedDateString)))
		titleString += strings.Repeat(" ", max(0, TITLE_LENGTH_LIMIT-len(titleString)))
		descString += strings.Repeat(" ", max(0, DESC_LENGTH_LIMIT-len(descString)))

		tempLine := getLineWithSeparators(idString, titleString, descString, completeFlagString, startedDateString, completedDateString)

		tempBlock += tempLine
	}

	resultLine += tempBlock + strings.Repeat("-", max(0, totalLength))
	return resultLine
}

func getLineWithSeparators(id, title, desc, status, timeStart, timeEnd string) string {
	return fmt.Sprintf("%v | %v | %v | %v | %v | %v |\n",
		id,
		title,
		desc,
		status,
		timeStart,
		timeEnd)
}

func getMaxIdLength(tm TodosMap) int {
	var idLength int
	var maxIdLength int

	for k := range tm.Todos {
		idLength = len(strconv.Itoa(k))
		if idLength > maxIdLength {
			maxIdLength = idLength
		}
	}

	return maxIdLength
}
