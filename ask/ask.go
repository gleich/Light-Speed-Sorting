package ask

import (
	"bufio"
	"fmt"
	"os"
	"runtime"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// OSBirthTime ... If birthtime no avalabile ask user if it is ok to still run
func OSBirthTime(cmd *cobra.Command, args []string) {
	if runtime.GOOS != "darwin" && runtime.GOOS != "freebsd" && runtime.GOOS != "netbsd" && runtime.GOOS != "windows" {
		askUser, err := cmd.Flags().GetBool("continuous")
		if err != nil {
			logrus.Fatal(err)
		}
		if askUser {
			fmt.Println("Looks like your OS doesn't store the creation date for files. Do you still want to run this program using the last modififed date instead? (y or n)")
			reader := bufio.NewReader(os.Stdin)
			continueProgram, err := reader.ReadString('\n')
			if err != nil {
				logrus.Fatal(err)
			}
			if continueProgram != "y\n" {
				os.Exit(0)
			}
		} else {
			logrus.Warn("YOUR OS DOESN'T STORE CREATION DATES FOR FILES. STEAMROLLING AND USING LAST MODIFIED DATE")
		}
	}
}
