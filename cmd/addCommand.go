package cmd

import (
	"fmt"
	"strings"
	"taskcli/cmd/constant"
	"time"

	"github.com/spf13/cobra"
)

var addTaskCommand = &cobra.Command{
	Use:   `add "task content"`,
	Short: "Add a task",
	Run: func(cmd *cobra.Command, args []string) {
		taskContent := strings.Join(args, " ")
		addTask(taskContent)
	},
}

func addTask(taskContent string) {
	tasks, taskCount, err := readJsonFile()
	if err != nil {
		return
	}

	task := Task{taskCount + 1, taskContent, TODO, time.Now(), time.Now()}
	tasks = append(tasks, task)
	writeJsonFile(tasks)
	fmt.Printf(constant.ColorGreen+"Task added with id: %d"+constant.ColorReset, taskCount+1)
}
