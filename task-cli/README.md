// Track what you are going to do
// Track what you have done
// Track what you are doing

// Must work from the CLI
// accept users actions and inputs as arguments
// and store the file in a JSON file

// user should be able to
// add, update and delete task
// make a task as done or in progress
// list all tasks
// list all done tasks
// list all ongoing tasks
// list tasks that are in progress.

// store the data in JSON file in the current directory

// sample response
/*
	# Adding a new task
	task-cli add "Buy groceries"
	# Output: Task added successfully (ID: 1)

	# Updating and deleting tasks
	task-cli update 1 "Buy groceries and cook dinner"
	task-cli delete 1

	# Marking a task as in progress or done
	task-cli mark-in-progress 1
	task-cli mark-done 1

	# Listing all tasks
	task-cli list

	# Listing tasks by status
	task-cli list done
	task-cli list todo
	task-cli list in-progress
*/