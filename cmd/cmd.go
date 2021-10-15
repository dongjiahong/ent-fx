// Package cmd 命令包
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"web/cmd/serve"
)

var Cmd = &cobra.Command{
	Use:   "web",
	Short: "web template",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Fprintf(os.Stderr, "you can try adding a '--help or -h' flag \n")
	},
}

func init() {
	Cmd.AddCommand(serve.ServeCMD)
}
