## Todo cli app
Made in Go.

**Features:**
- Simple UI
- Todo CRUD
- Visual display of the Todo table
- Saving and loading Todo data locally using .json

**To get started:**
- Type "go build -o bin/todo.exe" into the root folder terminal

This will create the "bin" folder and build an executable, which is the cli app. 

On first launch of the app a .json file will also be created in the "bin" folder. It serves as a local save file.

#### Documentation

The app supports several simple commands. They only contain 1 word and don't need any additional flags. Most of the commands are interactive, meaning they will ask the user to enter needed data as the command is being executed.

The full list of commands and their functionality:
| Command | Functionality |
|---------|---------------|
|help|display available commands|
|quit|close the app|
|show|display the table of todos|
|add|add a new todo|
|delete|delete an existing todo by id|
|toggle|toggle whether a todo is completed or not|
|edit|edit a title, a description or both for the todo by id|

The "help" and "quit" commands also support a short version: "h" and "q" respectively.

##### Data structure

The todo table is structured in the following way:
| id | title | description | completion status | time started | time finished |
|--|--|--|--|--|--|
| int | string | string | bool | RFC822 date & time format* | RFC822 date & time format* |

<sup>\* Format is provided by Go's standard "time" library. An example of the formatting: 
"23 Dec 98 13:45 EST"</sup>

##### JSON

An example .json save file with 2 todos looks like this:

```json
{
    "IdCounter": 2,
    "Todos": {
        "0": {
            "Id": 0,
            "Title": "buy milk",
            "Description": "needed to make pancakes",
            "Completed": true,
            "StartedAt": "2024-10-14T18:37:52.5178479+03:00",
            "CompletedAt": "2024-10-14T18:45:23.5083906+03:00"
        },
        "1": {
            "Id": 1,
            "Title": "program in Go",
            "Description": "an awesome Todo CLI",
            "Completed": false,
            "StartedAt": "2024-10-14T18:38:34.1289959+03:00",
            "CompletedAt": null
        }
    }
}
```