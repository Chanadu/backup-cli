/*
Copyright © 2024 Chandu Peddada chandu.peddada@gmail.com
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:     "Backup-Cli server password [files]",
	Aliases: []string{"backup-cli", "backupCli"},
	Short:   "An easy tool to backup files to a remote server.",
	Long:    `This Cli tool will allow you to connect to a remote server and backup files with ease.`,

	Args: cobra.MinimumNArgs(3),

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Backup-Cli is a tool to backup files to a remote server.")
		fmt.Println("")
		fmt.Println("--------------------")
		fmt.Println("")

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
	rootCmd.PersistentFlags().BoolP("debug", "d", false, "debug mode(includes messages of ran commands)")

	// cobra.OnInitialize(InitConfig)

}
