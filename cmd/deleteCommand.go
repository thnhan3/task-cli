package cmd

import (
	"fmt"
	"strconv"
	"taskcli/cmd/constant"

	"github.com/spf13/cobra"
)

var deleteCommand = &cobra.Command{
	Use:   "delete",
	Short: "Delete a task",
	Run: func(cmd *cobra.Command, args []string) {
		taskId, err := strconv.Atoi(args[0])
		if err != nil {
			panic(err)
		}

		deleteTask(taskId)
		fmt.Printf("%sTask %d deleted%s", constant.ColorRed, taskId, constant.ColorReset)
	},
}

func deleteTask(taskId int) {
	tasks, _, err := readJsonFile()
	isDeleted := false
	if err != nil {
		panic(err)
	}

	for i, task := range tasks {
		if task.ID == taskId {
			tasks = append(tasks[:i], tasks[i+1:]...)
			writeJsonFile(tasks)
			return
		}
	}
	if isDeleted == false {
		fmt.Println("Invalid Id", taskId)
	}
}
