package serve

import (
	"github.com/spf13/cobra"
)

var stopCMD = &cobra.Command{
	Use:   "stop",
	Short: "stop web serve",
	Run: func(cmd *cobra.Command, args []string) {
		stop()
	},
}

func stop() {
}
