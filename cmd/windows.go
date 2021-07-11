package cmd

import (
	"fmt"
	"os/exec"
	"strconv"

	"github.com/spf13/cobra"
)

// windowsCmd represents the windows command
var windowsCmd = &cobra.Command{
	Use:   "windows",
	Short: "Control Windows OS",
	Long:  `Options:
	Shutdown forcefully using timer (in minutes)`,
	PreRun: func(cmd *cobra.Command, args []string) {if len(args) == 0 {
		cmd.Help()
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("windows called")
		fmt.Println(args)
		
		// explanation on indexing of slices in golang 
		// https://www.golangprograms.com/how-to-get-first-or-last-element-of-slice-in-golang.html
		minutes := args[len(args)-1]
		cmd.Execute(shutdownTimer(minutes))
	}, 
	Aliases: []string{"win","w"},
}

func init() {
	rootCmd.AddCommand(windowsCmd)
	windowsCmd.PersistentFlags().String("shutdown", "", "A help for foo")
	windowsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func shutdownTimer(minutes string) {
	if minutes == "now" {
		exec.Command(`shutdown.exe -s -f -t 0`)
	}

	secondsConverted, err:= (strconv.Atoi(minutes))
	cobra.CheckErr(err)
	seconds := secondsConverted * 60
	exec.Command(`shutdown.exe -s -f -t `+strconv.Itoa(seconds))
}
