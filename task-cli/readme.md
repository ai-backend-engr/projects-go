# Task Tracker CLI

A command-line task tracker built in Go. This project is based on the [Task Tracker CLI project on roadmap.sh](https://roadmap.sh/projects/task-tracker).

Track what you need to do, what you are working on, and what you have done. Tasks are stored in a JSON file in the current directory.

## Features

- Add, update, and delete tasks
- Mark a task as in progress or done
- List all tasks
- List tasks by status (`todo`, `done`, `in-progress`)

## Task Properties

Each task has the following fields:

| Field         | Description                                      |
| ------------- | ------------------------------------------------ |
| `id`          | Unique identifier for the task                   |
| `description` | Short description of the task                      |
| `status`      | One of `todo`, `in-progress`, or `done`          |
| `createdAt`   | Date and time when the task was created          |
| `updatedAt`   | Date and time when the task was last updated     |

## Requirements

- Go 1.22 or later

## Build

From the `task-cli` directory:

```bash
go build -o task ./cmd
```

## Usage

Run the binary from the directory where you want `tasks.json` to be created. The file is created automatically when you add your first task.

### Adding a new task

```bash
./task -add "Buy groceries"
# Output: Task added successfully (ID: 1)
```

### Updating and deleting tasks

```bash
./task -update 1 "Buy groceries and cook dinner"
./task -delete 1
```

When updating, the new description must be passed as a positional argument after the flags.

### Marking a task as in progress or done

```bash
./task -mark-in-progress 1
./task -mark-done 1
```

### Listing tasks

```bash
# List all tasks
./task -list

# List tasks by status
./task -list todo
./task -list done
./task -list in-progress
```

## Data Storage

Tasks are persisted to `./tasks.json` in the current working directory. If the file does not exist, it is created when a task is saved.

## Project Structure

```
task-cli/
├── cmd/           # CLI entrypoint
├── pkg/           # Task model and persistence
└── internal/ui/   # Table rendering
```



