package cmd

import (
	"fmt"

	"github.com/Matt-Gleich/Light-Speed-Sorting/ask"
	"github.com/Matt-Gleich/Light-Speed-Sorting/files"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "light-speed-sorting",
	Short: "Blazing fast file sorting",
	Run: func(cmd *cobra.Command, args []string) {
		ask.OSBirthTime(cmd, args)
		fmt.Println(files.GetFiles())
	},
}

// Execute ... Execute root command
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		logrus.Fatal(err)
	}
}

func init() {
	rootCmd.Flags().Bool("no-out", true, "Show no stdout except for errors")
	rootCmd.Flags().BoolP("continuous", "c", true, "Run the sorter every 30 seconds")
}
