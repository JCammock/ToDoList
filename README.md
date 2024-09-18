# To Do List

This to do list application allows users to track a custom list of tasks, the application currently has the following commands:

### ToDoList add

The `ToDoList add [task]` command will add the specified task to the list and output the updated list:
```
$ ToDoList add test
| ID | TASK     | STATUS      |
| 1  | test     | Not Started |
| 2  | test2    | Not Started |
| 3  | test2233 | Not Started |
| 4  | test     | Not Started |
```

### ToDoList delete

The `ToDoList delete [taskId]` command will delete the task matching the ID:
```
$ ToDoList delete 2
| ID | TASK     | STATUS      |
| 1  | test     | Not Started |
| 2  | test2233 | Not Started |
| 3  | test     | Not Started |
```

### ToDoList list

The `ToDoList list` command will list all current tasks in the list:
```
$ ToDoList list
| ID | TASK     | STATUS      |
| 1  | test     | Not Started |
| 2  | test2    | Not Started |
| 3  | test2233 | Not Started |
| 4  | test     | Not Started |
```

### ToDoList start

The `ToDoList start [taskId]` command will change the status of the specified task to "In Progress":
```
$ ToDoList start 3
| ID | TASK     | STATUS      |
| 1  | test     | Not Started |
| 2  | test2    | Not Started |
| 3  | test2233 | In Progress |
| 4  | test     | Not Started |
```

### ToDoList complete

The `ToDoList complete [taskId]` command will change the status of the specified task to "Completed":
```
$ ToDoList complete 3
| ID | TASK     | STATUS      |
| 1  | test     | Not Started |
| 2  | test2    | Not Started |
| 3  | test2233 | Completed   |
| 4  | test     | Not Started |
```