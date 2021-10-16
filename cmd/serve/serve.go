package serve

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var ServeCMD = &cobra.Command{
	Use:   "serve",
	Short: "manage serve",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Fprintf(os.Stderr, " you can try adding a '--help or -h' flag \n")
	},
}

func init() {
	ServeCMD.AddCommand(startCMD)
}
