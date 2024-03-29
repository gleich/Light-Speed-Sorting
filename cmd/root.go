package cmd

import (
	"time"

	"github.com/Matt-Gleich/Light-Speed-Sorting/ask"
	"github.com/Matt-Gleich/Light-Speed-Sorting/files"
	"github.com/Matt-Gleich/statuser/v2"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "light-speed-sorting",
	Short: "Blazing fast file sorting",
	Run: func(cmd *cobra.Command, args []string) {
		continuous, err := cmd.Flags().GetBool("continuous")
		if err != nil {
			statuser.Error("Failed to get 'continuous' flag", err, 1)
		}
		if continuous {
			for {
				run(cmd, args)
				time.Sleep(30 * time.Second)
			}
		} else {
			run(cmd, args)
		}
	},
}

// Execute ... Execute root command
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		statuser.Error("Failed to execute root command", err, 1)
	}
}

func init() {
	rootCmd.Flags().BoolP("continuous", "c", false, "Run the sorter every 30 seconds")
}

func run(cmd *cobra.Command, args []string) {
	ask.OSBirthTime(cmd, args)
	projectFiles := files.GetFiles()
	files.MoveFiles(projectFiles)
}
