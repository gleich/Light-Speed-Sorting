package cmd

import (
	"fmt"

	"github.com/Matt-Gleich/Light-Speed-Sorting/files"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "light-speed-sorting",
	Short: "Blazing fast file sorting",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(files.GetFiles())
	},
}

// Execute ... Execute main command
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		logrus.Fatal(err)
	}
}

func init() {
	rootCmd.Flags().Bool("no-out", false, "Show no stdout except for errors")
}
