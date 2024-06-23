package cmd

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
)

func BackupFiles(args []string) {

	if err := exec.Command(
		"bash",
		"-c",
		"sshpass -p "+args[1]+" ssh "+args[0]+" StrictHostKeyChecking=no 'echo 1; exit'",
	).Run(); err != nil {
		fmt.Printf("Error connecting to server: %s\n", err)
		os.Exit(1)
	}

	//		fmt.Printf("sshpass -p %s ssh %s StrictHostKeyChecking=no '%sexit'",
	//			args[1],
	//			args[0],
	//			deleteServerBackupsCmd,
	//		)

	createBackupFiles(args)
	fmt.Printf("\n---------------------------------------\n")
	deleteServerBackupFiles(args)
	fmt.Printf("\n---------------------------------------\n")
	scpBackupFiles(args)
	fmt.Printf("\n---------------------------------------\n")
	deleteLocalBackupFiles(args)
	fmt.Printf("\n---------------------------------------\n")
	fmt.Println("Backup completed successfully.")
}

func createBackupFiles(args []string) {
	createCmdText := ""
	for i := 2; i < len(args); i++ {
		createCmdText = createCmdText + fmt.Sprintf(
			"7z a %s.7z; mv %s.7z %s-Backup.7z; ",
			args[i],
			args[i],
			args[i],
		)
	}

	var osOut bytes.Buffer
	createCmd := exec.Command("bash", "-c", createCmdText)
	createCmd.Stdout = &osOut

	if err := createCmd.Run(); err != nil {
		fmt.Printf("Error creating backup files: %s\n", err)
		os.Exit(1)
	}
	fmt.Printf("Command Ran: %s\n", createCmd.String())
	fmt.Printf("Backup files created: %s\n", osOut.String())
}

func scpBackupFiles(args []string) {
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
	var osOut bytes.Buffer
	scpCmd := exec.Command("bash", "-c", scpCmdText)
	scpCmd.Stdout = &osOut

	if err := exec.Command("bash", "-c", scpCmdText).Run(); err != nil {
		fmt.Printf("Error copying backup files to server: %s\n", err)
		os.Exit(1)
	}
	fmt.Printf("Command Ran: %s\n", scpCmd.String())
	fmt.Println("Backup files copied to server: ", osOut.String())
}

func deleteServerBackupFiles(args []string) {
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
	var osOut bytes.Buffer
	deleteServerCmd := exec.Command("bash",
		"-c",
		deleteServerCmdText,
	)
	deleteServerCmd.Stdout = &osOut

	if err := deleteServerCmd.Run(); err != nil {
		fmt.Printf("Error deleting old server backup files: %s\n", err)
		os.Exit(1)
	}

	fmt.Printf("Command Ran: %s\n", deleteServerCmdText)
	fmt.Printf("Old server backup files deleted: %s\n", osOut.String())
}

func deleteLocalBackupFiles(args []string) {
	deleteCmd := ""
	for i := 2; i < len(args); i++ {
		deleteCmd = deleteCmd + fmt.Sprintf("rm %s-Backup.7z; ", args[i])
	}

	if err := exec.Command("bash", "-c", deleteCmd).Run(); err != nil {
		fmt.Printf("Error deleting local backup files: %s\n", err)
		os.Exit(1)
	}

	fmt.Println("Local backup files deleted: ", deleteCmd)
}
