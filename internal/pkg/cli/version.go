package cli

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Prints program version information",
	Long:  `Prints program version information`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(viper.GetString("version"))
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
