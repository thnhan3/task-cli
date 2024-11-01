package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

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
	if err != nil {
		fmt.Println("Error encoding task:", err)
		return
	}

	fmt.Println("create task.json")

}
func readJsonFile() (tasks []Task, taskCount int, err error) {
	file, err := os.Open("task.json")
	if err != nil {
		if os.IsNotExist(err) {
			return []Task{}, 0, nil
		}
		return nil, 0, err
	}

	defer file.Close()

	// Read the file contents into a byte slice
	data, err := io.ReadAll(file)
	if err != nil {
		return nil, 0, err
	}

	if len(data) == 0 {
		return []Task{}, 0, nil
	}

	// Unmarshal the JSON data from the byte slice
	err = json.Unmarshal(data, &tasks)
	if err != nil {
		return nil, 0, err
	}

	return tasks, len(tasks), nil
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
