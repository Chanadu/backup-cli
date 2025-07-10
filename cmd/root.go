/*
Copyright © 2025 Chandu Peddada chandu.peddada@gmail.com
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:     "backup-cli server password [files]",
	Aliases: []string{"Backup-Cli", "backupCli"},
	Short:   "An easy tool to backup files to a remote server.",
	Long:    `This Cli tool will allow you to connect to a remote server and backup files with ease.`,

	Args: cobra.MinimumNArgs(3),

	Run: func(cmd *cobra.Command, args []string) {
		_, _ = fmt.Println("backup-cli is a tool to backup files to a remote server. (1.1)")
		_, _ = fmt.Println("")
		_, _ = fmt.Println("--------------------")
		_, _ = fmt.Println("")
		workingDir, err := os.Getwd()
		if err != nil {
			printErrorf("Error getting current working directory: %s\n", err)
			os.Exit(1)
		}

		for i := 2; i < len(args); i++ {
			if _, err := os.Stat(workingDir + "/" + args[i]); err != nil {
				printErrorf("Error: File %s does not exist in the current working directory\n", args[i])
				printErrorf("File is not at: %s\n", workingDir+args[i])
				os.Exit(1)
			}
		}
		BackupFiles(cmd, args)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().BoolP("debug", "d", false, "DEPRECIATED: USE -v or --verbose, debug mode(includes messages of ran commands)")

	err := rootCmd.PersistentFlags().MarkDeprecated("debug", "Use -v or --verbose instead")
	if err != nil {
		printErrorf("Error marking debug flag as depreciated: %s\n", err)
		os.Exit(1)
	}

	rootCmd.PersistentFlags().BoolP("verbose", "v", false, "verbose mode(includes messages of ran commands)")
	// cobra.OnInitialize(InitConfig)
}
