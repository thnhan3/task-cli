package cmd

import (
	"fmt"
	"strconv"
	"taskcli/cmd/constant"

	"github.com/spf13/cobra"
)

func deleteTaskById(taskId int) {
	tasks, _, err := readJsonFile()
	if err != nil {
		panic(err)
	}

	index := -1
	for i, task := range tasks {
		if task.ID == taskId {
			index = i
			break
		}
	}

	if index == -1 {
		fmt.Println(constant.ColorRed + "Task ID not found" + constant.ColorReset)
		return
	}

	// Remove the task from the slice
	tasks = append(tasks[:index], tasks[index+1:]...)

	writeJsonFile(tasks)
	fmt.Println(constant.ColorGreen + "Task deleted successfully" + constant.ColorReset)
}

func clearAll() {
	writeJsonFile(nil)
}

var clearTaskCommand = &cobra.Command{
	Use:   "clear",
	Short: "Clear tasks",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			clearAll()
		}
		if len(args) == 1 {
			taskId, err := strconv.Atoi(args[0])
			if err != nil {
				panic(err)
			}

			deleteTaskById(taskId)
		} else if len(args) > 1 {
			tasks, _, err := readJsonFile()
			if err != nil {
				panic(err)
			}

			for _, arg := range args {
				taskId, err := strconv.Atoi(arg)
				if err == nil {
					index := -1
					for i, task := range tasks {
						if task.ID == taskId {
							index = i
							break
						}
					}
					if index != -1 {
						tasks = append(tasks[:index], tasks[index+1:]...)
					}
				}
			}
			writeJsonFile(tasks)
			fmt.Println(constant.ColorGreen + "Tasks deleted successfully" + constant.ColorReset)
		}
	},
}
