package serve

import (
	"github.com/spf13/cobra"
)

var startCMD = &cobra.Command{
	Use:   "start",
	Short: "start web serve",
	Run: func(cmd *cobra.Command, args []string) {
		start()
	},
}

func start() {
}
