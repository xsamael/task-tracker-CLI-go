
# Task Tracker CLI: Manage Your Tasks with Ease  

Welcome to **Task Tracker CLI**, a command-line application to manage your tasks efficiently! Whether you're juggling multiple projects or just trying to stay organized, this simple and intuitive tool has got you covered.

## Features  
- **Add tasks**: Keep track of what needs to be done.  
- **Update tasks**: Modify task descriptions easily.  
- **Delete tasks**: Remove tasks that are no longer needed.  
- **Change task status**: Mark tasks as `todo`, `in-progress`, or `done`.  
- **List tasks**: View all tasks or filter them by status.  

## Getting Started  

### Prerequisites  
- A programming language runtime installed on your system (e.g., Python or Node.js).  
- Basic familiarity with using the terminal or command line.  

### Installation  
1. Clone this repository:  
   ```bash
   git clone https://github.com/yourusername/task-tracker-cli.git
   cd task-tracker-cli
   ```  
2. (Optional) If required, set execution permissions on the script:  
   ```bash
   chmod +x task-cli
   ```

### Usage  
The CLI accepts commands and arguments directly from the terminal. Below are the available commands:  

#### **Add a new task**  
```bash
task-cli add "Your task description"
```  
Example:  
```bash
task-cli add "Buy groceries"
# Output: Task added successfully (ID: 1)
```  

#### **Update an existing task**  
```bash
task-cli update <task-id> "Updated task description"
```  
Example:  
```bash
task-cli update 1 "Buy groceries and cook dinner"
# Output: Task updated successfully
```  

#### **Delete a task**  
```bash
task-cli delete <task-id>
```  
Example:  
```bash
task-cli delete 1
# Output: Task deleted successfully
```  

#### **Mark a task as in-progress or done**  
```bash
task-cli mark-in-progress <task-id>
task-cli mark-done <task-id>
```  
Example:  
```bash
task-cli mark-in-progress 1
# Output: Task marked as in progress
task-cli mark-done 1
# Output: Task marked as done
```  

#### **List tasks**  
- To list all tasks:  
  ```bash
  task-cli list
  ```  
- To filter tasks by status (`done`, `todo`, or `in-progress`):  
  ```bash
  task-cli list <status>
  ```  
Example:  
```bash
task-cli list
# Output: 
# 1. Buy groceries (todo) - Created at: YYYY-MM-DD HH:MM:SS

task-cli list done
# Output: No tasks found with status "done"
```  

## Task Properties  
Each task in Task Tracker CLI has the following properties:  
- **ID**: A unique identifier assigned automatically.  
- **Description**: A short description of the task.  
- **Status**: The current status of the task (`todo`, `in-progress`, or `done`).  
- **Created At**: The date and time when the task was created.  
- **Updated At**: The date and time when the task was last updated.  

## File Storage  
Tasks are stored locally in a JSON file named `tasks.json`, located in the project directory. This file is automatically created if it doesn't exist.  

## Error Handling  
The CLI gracefully handles errors, such as:  
- Invalid task IDs (e.g., trying to update or delete a non-existent task).  
- Missing or incorrect arguments.  

If an error occurs, the CLI will provide a helpful error message.  

## Contributing  
Want to improve Task Tracker CLI? Feel free to fork the repository and submit a pull request. Contributions are always welcome!  

## License  
This project is licensed under the MIT License.  

## Additional Resources  
For more project ideas and inspiration, visit the [Task Tracker project page](https://roadmap.sh/projects/task-tracker).  

---  
Start managing your tasks effectively with Task Tracker CLI today! 🚀  
