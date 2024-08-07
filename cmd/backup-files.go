package cmd

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

func BackupFiles(cmd *cobra.Command, args []string) {

	if err := exec.Command(
		"bash",
		"-c",
		"sshpass -p "+args[1]+" ssh "+args[0]+" StrictHostKeyChecking=no 'echo 1; exit'",
	).Run(); err != nil {
		fmt.Printf("Error connecting to server: %s\n", err)
		os.Exit(1)
	}

	debugFlag, err := cmd.Flags().GetBool("debug")
	if err != nil {
		fmt.Printf("Error getting debug flag: %s\n", err)
		os.Exit(1)
	}
	verboseFlag, err := cmd.Flags().GetBool("verbose")
	if err != nil {
		fmt.Printf("Error getting verbose flag: %s\n", err)
		os.Exit(1)
	}
	isDebug := debugFlag || verboseFlag

	createBackupFiles(args, isDebug)
	fmt.Printf("\n---------------------------------------\n\n")
	deleteServerBackupFiles(args, isDebug)
	fmt.Printf("\n---------------------------------------\n\n")
	scpBackupFiles(args, isDebug)
	fmt.Printf("\n---------------------------------------\n\n")
	deleteLocalBackupFiles(args, isDebug)
	fmt.Printf("\n---------------------------------------\n\n")
	_, _ = fmt.Println("Backup completed successfully.")
}

func createBackupFiles(args []string, isDebug bool) {
	for i := 2; i < len(args); i++ {
		createCmdText := fmt.Sprintf(
			"7z a -m0=lzma2 -mx=9 %s-Backup.7z %s",
			args[i],
			args[i],
		)

		if isDebug {
			fmt.Printf("Command Running: %s\n", createCmdText)
		}

		err := runCommand(createCmdText)

		if err != nil {
			fmt.Printf("Error creating backup file (%s): %s\n", args[i], err)
			os.Exit(1)
		}

		fmt.Printf("Backup file created: (%s)\n", args[i])
	}

}

func scpBackupFiles(args []string, isDebug bool) {
	scpCmdText := ""
	for i := 2; i < len(args); i++ {
		scpCmdText = scpCmdText + fmt.Sprintf(
			"sshpass -p %s scp %s-Backup.7z %s:~/backups/%s-Backup.7z; ",
			args[1],
			args[i],
			args[0],
			args[i],
		)
	}

	if isDebug {
		fmt.Printf("Command Running: %s\n", scpCmdText)
	}

	err := runCommand(scpCmdText)
	if err != nil {
		fmt.Printf("Error copying backup files to server: %s\n", err)
		os.Exit(1)
	}

	_, _ = fmt.Println("Backup files copied to server.")
}

func deleteServerBackupFiles(args []string, isDebug bool) {
	deleteServerCmdText := ""
	for i := 2; i < len(args); i++ {
		deleteServerCmdText = deleteServerCmdText + fmt.Sprintf(
			"rm ~/backups/%s-Backup.7z -f; ",
			args[i],
		)
	}
	deleteServerCmdText = fmt.Sprintf("sshpass -p %s ssh %s StrictHostKeyChecking=no '%s exit'",
		args[1],
		args[0],
		deleteServerCmdText,
	)

	if isDebug {
		fmt.Printf("Command Running: %s\n", deleteServerCmdText)
	}

	err := runCommand(deleteServerCmdText)

	if err != nil {
		fmt.Printf("Error deleting old server backup files: %s\n", err)
		os.Exit(1)
	}

	_, _ = fmt.Println("Old server backup files deleted.")
}

func deleteLocalBackupFiles(args []string, isDebug bool) {
	deleteCmdText := ""
	for i := 2; i < len(args); i++ {
		deleteCmdText = deleteCmdText + fmt.Sprintf("rm %s-Backup.7z; ", args[i])
	}

	if isDebug {
		fmt.Printf("Command Running: %s\n", deleteCmdText)
	}

	err := runCommand(deleteCmdText)

	if err != nil {
		fmt.Printf("Error deleting local backup files: %s\n", err)
		os.Exit(1)
	}

	_, _ = fmt.Println("Local backup files deleted. ")
}

func runCommand(cmdText string) error {
	cmd := exec.Command("bash", "-c", cmdText)

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return err
	}

	stderr, err := cmd.StderrPipe()
	if err != nil {
		return err
	}

	err = cmd.Start()
	if err != nil {
		return err
	}

	scanner := bufio.NewScanner(io.MultiReader(stdout, stderr))
	// fmt.Println("Here")
	for scanner.Scan() {
		// fmt.Println("For Loop Running")
		m := scanner.Text()
		_, _ = fmt.Println(m)
	}
	err = cmd.Wait()
	return err
}
