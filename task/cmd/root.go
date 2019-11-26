package cmd

import (
	"github.com/spf13/cobra"
)

// RootCmd is the root command
var RootCmd = &cobra.Command{
	Use:   "task",
	Short: "Do a task",
	Long:  "Do a task",
}
