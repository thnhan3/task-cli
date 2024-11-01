package cmd

import (
	"fmt"
	"os"
	"strings"
	"taskcli/cmd/constant"

	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

var listCommand = &cobra.Command{
	Use:   "list",
	Short: "show list of tasks",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			getAllTasks()
		}
		if len(args) > 0 {
			action := args[0]

			if strings.EqualFold(action, "todo") ||
				strings.EqualFold(action, "in_process") ||
				strings.EqualFold(action, "done") {
				getTasksByStatus(action)
			} else {
				fmt.Println(constant.ColorRed + "Invalid action" + constant.ColorReset)
			}
		}
	},
}

func getTasksByStatus(action string) {
	tasks, _, err := readJsonFile()
	if err != nil {
		panic(err)
	}

	var filteredTasks []Task
	for _, task := range tasks {
		if strings.EqualFold(string(task.Status), action) {
			filteredTasks = append(filteredTasks, task)
		}
	}

	if len(filteredTasks) == 0 {
		fmt.Println(constant.ColorRed + "No tasks found with status: " + action + constant.ColorReset)
		return
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"ID", "Content", "Status", "Created At", "Updated At"})

	for _, task := range filteredTasks {
		row := []string{
			fmt.Sprintf("%d", task.ID),
			task.Content,
			getStatusColor(string(task.Status)),
			task.CreatedAt.Local().Format("02/01/2006 15:04"),
			task.UpdatedAt.Local().Format("02/01/2006 15:04"),
		}
		table.Append(row)
	}

	table.Render()
}

func getAllTasks() {
	tasks, taskSize, err := readJsonFile()
	if err != nil {
		panic(err)
	}

	if taskSize == 0 {
		fmt.Println("No task in data.")
		return
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"ID", "Content", "Status", "Created At", "Updated At"})

	for _, task := range tasks {
		row := []string{fmt.Sprintf("%d", task.ID),
			task.Content,
			getStatusColor(string(task.Status)),
			task.CreatedAt.Format("02/01/2006 15:04"),
			task.UpdatedAt.Format("02/01/2006 15:04"),
		}
		table.Append(row)

	}
	table.SetAlignment(tablewriter.ALIGN_CENTER)
	table.Render()
}

func getStatusColor(st string) string {
	if strings.EqualFold(st, string(TODO)) {
		return constant.ColorBlue + st + constant.ColorReset
	}
	if strings.EqualFold(st, string(DONE)) {
		return constant.ColorGreen + st + constant.ColorReset

	}
	if strings.EqualFold(st, string(IN_PROCESS)) {
		return constant.ColorPurple + st + constant.ColorReset
	}

	return st
}
