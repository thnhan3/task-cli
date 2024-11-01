package cmd

import (
	"fmt"
	"strconv"
	"taskcli/cmd/constant"

	"github.com/spf13/cobra"
)

var markInProcessCommand = &cobra.Command{
	Use:   "mark-in-process",
	Short: "Mark in process done by it id. Example: `task-cli mark-inprocess 1`",
	Run: func(cmd *cobra.Command, args []string) {
		taskId, err := strconv.Atoi(args[0])
		if err != nil {
			panic(err)
		}
		fmt.Println(taskId)
		markInProcessById(taskId)
	},
}

func markInProcessById(taskId int) {
	isDone := false
	tasks, _, err := readJsonFile()
	if err != nil {
		panic(err)
	}
	for i, v := range tasks {
		if v.ID == taskId {
			isDone = true
			tasks[i].Status = ETaskStatus("IN_PROCESS")
			fmt.Println("ok")
			break
		}
	}

	writeJsonFile(tasks)
	if isDone == false {
		fmt.Println(constant.ColorRed + "task id not found" + constant.ColorReset)
	}

}
