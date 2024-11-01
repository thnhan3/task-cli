package cmd

import (
	"fmt"
	"strconv"
	"strings"
	"taskcli/cmd/constant"

	"github.com/spf13/cobra"
)

var updateCommand = &cobra.Command{
	Use:   "update",
	Short: "Update a task",
	Run: func(cmd *cobra.Command, args []string) {
		var newContent = strings.Join(args[1:], " ")
		taskId, _ := strconv.Atoi(args[0])

		updateTask(taskId, newContent)
	},
}

func updateTask(taskId int, newContent string) {
	tasks, _, err := readJsonFile()
	isUpdated := false
	if err != nil {
		return
	}

	for i, task := range tasks {
		if task.ID == taskId {
			tasks[i].Content = newContent
			writeJsonFile(tasks)
			fmt.Printf("%sTask %d updated%s", constant.ColorBlue, taskId, constant.ColorReset)
			return
		}
	}
	if isUpdated == false {
		fmt.Println("Invalid Id", taskId)
	}
}
