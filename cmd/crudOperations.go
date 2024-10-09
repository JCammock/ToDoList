package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"text/tabwriter"
	"time"
)

type Tasks struct {
	Tasks []Task `json:"tasks"`
}

type Task struct {
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Status    string `json:"status"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

func getFileData() (Tasks, error) {
	jsonFile, err := os.Open("tasks.json")

	if err != nil {
		return Tasks{}, err
	}

	defer jsonFile.Close()

	jsonData, _ := io.ReadAll(jsonFile)

	var tasks Tasks

	err = json.Unmarshal(jsonData, &tasks)

	if err != nil {
		errorMessage := fmt.Sprint("Error: ", err)
		panic(errorMessage)
	}

	return tasks, nil
}

func getCurrentMaxInt() int {
	tasks, err := getFileData()

	if err != nil {
		return 0
	}

	var ids []int

	for task := 0; task < len(tasks.Tasks); task++ {
		ids = append(ids, tasks.Tasks[task].Id)
	}

	if len(ids) == 0 {
		return 0
	}
	largest := ids[0]
	for id := 0; id < len(ids); id++ {
		if ids[id] > largest {
			largest = ids[id]
		}
	}

	return largest
}

func addTask(taskTitle string) {
	tasks, _ := getFileData()

	task := Task{
		Id:        getCurrentMaxInt() + 1,
		Title:     taskTitle,
		Status:    "todo",
		CreatedAt: time.Now().Format("2006-01-02 15:04:05"),
		UpdatedAt: time.Now().Format("2006-01-02 15:04:05"),
	}

	tasks.Tasks = append(tasks.Tasks, task)

	jsonData, err := json.Marshal(tasks)

	if err != nil {
		errorMessage := fmt.Sprint("Error: ", err)
		panic(errorMessage)
	}
	jsonFile, err := os.Create("tasks.json")

	if err != nil {
		errorMessage := fmt.Sprint("Error: ", err)
		panic(errorMessage)
	}

	jsonFile.Write(jsonData)

	listTasks()
}

func deleteTask(taskID int) {
	tasks, _ := getFileData()

	if taskID >= len(tasks.Tasks) || taskID < 1 {
		fmt.Println("Task does not exist")
		return
	}

	for i := taskID - 1; i < len(tasks.Tasks); i++ {
		tasks.Tasks[i].Id -= 1
	}

	tasks.Tasks = append(tasks.Tasks[:taskID-1], tasks.Tasks[taskID:]...)

	jsonData, err := json.Marshal(tasks)

	if err != nil {
		errorMessage := fmt.Sprint("Error: ", err)
		panic(errorMessage)
	}

	jsonFile, err := os.Create("tasks.json")

	if err != nil {
		errorMessage := fmt.Sprint("Error: ", err)
		panic(errorMessage)
	}

	jsonFile.Write(jsonData)

	listTasks()
}

func listTasks(status ...string) {
	fmt.Println(status)
	tasks, err := getFileData()

	if err != nil {
		fmt.Println("No tasks in list")
	}

	tabFormatting := tabwriter.NewWriter(os.Stdout, 0, 10, 1, ' ', 0)

	fmt.Fprintln(tabFormatting, "|\tID\t|\tTASK\t|\tSTATUS\t|\tCREATED AT\t|\tUPDATED AT\t|")

	for i := 0; i < len(tasks.Tasks); i++ {
		if tasks.Tasks[i].Status != status[0] {
			continue
		}
		fmt.Fprintf(
			tabFormatting, "|\t%d\t|\t%s\t|\t%s\t|\t%s\t|\t%s\t|\n",
			tasks.Tasks[i].Id,
			tasks.Tasks[i].Title,
			tasks.Tasks[i].Status,
			tasks.Tasks[i].CreatedAt,
			tasks.Tasks[i].UpdatedAt,
		)
	}

	tabFormatting.Flush()
}

func startTask(taskID int) {
	tasks, err := getFileData()

	if err != nil {
		errorMessage := fmt.Sprint("Error: ", err)
		panic(errorMessage)
	}

	if taskID >= len(tasks.Tasks) || taskID < 0 {
		fmt.Println("Task does not exist")
		return
	}

	tasks.Tasks[taskID].Status = "in-progress"
	tasks.Tasks[taskID].UpdatedAt = time.Now().Format("2006-01-02 15:04:05")

	jsonData, err := json.Marshal(tasks)

	if err != nil {
		errorMessage := fmt.Sprint("Error: ", err)
		panic(errorMessage)
	}

	jsonFile, err := os.Create("tasks.json")

	if err != nil {
		errorMessage := fmt.Sprint("Error: ", err)
		panic(errorMessage)
	}

	jsonFile.Write(jsonData)

	listTasks()
}

func updateTask(taskID int, taskTitle string) {
	tasks, err := getFileData()

	if err != nil {
		errorMessage := fmt.Sprint("Error: ", err)
		panic(errorMessage)
	}

	if taskID >= len(tasks.Tasks) || taskID < 0 {
		fmt.Println("Task does not exist")
		return
	}

	tasks.Tasks[taskID].Title = taskTitle
	tasks.Tasks[taskID].UpdatedAt = time.Now().Format("2006-01-02 15:04:05")

	jsonData, err := json.Marshal(tasks)

	if err != nil {
		errorMessage := fmt.Sprint("Error: ", err)
		panic(errorMessage)
	}

	jsonFile, err := os.Create("tasks.json")

	if err != nil {
		errorMessage := fmt.Sprint("Error: ", err)
		panic(errorMessage)
	}

	jsonFile.Write(jsonData)

	listTasks()
}

func finishTask(taskID int) {
	tasks, err := getFileData()

	if err != nil {
		errorMessage := fmt.Sprint("Error: ", err)
		panic(errorMessage)
	}

	if taskID >= len(tasks.Tasks) || taskID < 0 {
		fmt.Println("Task does not exist")
		return
	}

	tasks.Tasks[taskID].Status = "done"
	tasks.Tasks[taskID].UpdatedAt = time.Now().Format("2006-01-02 15:04:05")

	jsonData, err := json.Marshal(tasks)

	if err != nil {
		errorMessage := fmt.Sprint("Error: ", err)
		panic(errorMessage)
	}

	jsonFile, err := os.Create("tasks.json")

	if err != nil {
		errorMessage := fmt.Sprint("Error: ", err)
		panic(errorMessage)
	}

	jsonFile.Write(jsonData)

	listTasks()
}
