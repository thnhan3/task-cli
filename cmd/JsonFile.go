package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/spf13/cobra"
)

var createJsonFile = &cobra.Command{
	Use:   "hello",
	Short: "Create task.json",
	Run: func(cmd *cobra.Command, args []string) {
		createFile()
	}}

var readjsonFile = &cobra.Command{
	Use:   "read",
	Short: "Read task.json",
	Run: func(cmd *cobra.Command, args []string) {
		readJsonFile()
	},
}

func createFile() {
	file, err := os.Create("task.json")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()
	task1 := Task{1, "Task 1", TODO, time.Now(), time.Now()}

	err = json.NewEncoder(file).Encode(task1)
	if err != nil {
		fmt.Println("Error encoding task:", err)
		return
	}

	fmt.Println("create ok")

}

func readJsonFile() (tasks []Task, taskCount int, err error) {
	file, err := os.Open("task.json")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	var task []Task
	err = json.NewDecoder(file).Decode(&task)
	if err != nil {
		fmt.Println("Error decoding task:", err)
		return
	}

	return task, len(task), err
}

func writeJsonFile(tasks []Task) {
	file, err := os.Create("task.json")
	if err != nil {
		return
	}

	for i := 0; i < len(tasks); i++ {
		tasks[i].ID = i + 1
	}

	data, err := json.Marshal(tasks)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	_, err = file.Write(data)
	if err != nil {
		panic(err)
	}
}
