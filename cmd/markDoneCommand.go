package cmd

import (
	"fmt"
	"strconv"
	"taskcli/cmd/constant"

	"github.com/spf13/cobra"
)

var markDoneCommand = &cobra.Command{
	Use:   "mark-done",
	Short: "Mark task done by id. Example: `task-cli mark-done 1`",
	Run: func(cmd *cobra.Command, args []string) {
		taskId, err := strconv.Atoi(args[0])
		if err != nil {
			panic(err)
		}
		makeDoneById(taskId)
	},
}

func makeDoneById(taskId int) {
	isDone := false
	tasks, _, err := readJsonFile()
	if err != nil {
		panic(err)
	}
	for i, v := range tasks {
		if v.ID == taskId {
			isDone = true
			tasks[i].Status = ETaskStatus("DONE")
			break
		}
	}

	writeJsonFile(tasks)
	if isDone == false {
		fmt.Println(constant.ColorRed + "task id not found" + constant.ColorReset)
	}

}
