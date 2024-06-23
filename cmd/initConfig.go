package cmd

import (
	"fmt"
	"os"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

func InitConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
		return
	}

	home, err := homedir.Dir()

	if err != nil {
		fmt.Println("Cant find home dir: ", err)
		os.Exit(1)
	}
	fmt.Println("home: ", home)

	viper.SetConfigFile(home + "/.config/Backup-Cli/config.yaml")
	viper.SetDefault("test123", "abc")
	if err := viper.SafeWriteConfigAs(home + "/.config/Backup-Cli/config.yaml"); err != nil {
		fmt.Println("File Exists", err)
	}

	c := &conf{}
	if err := viper.Unmarshal(c); err != nil {
		fmt.Println("Error unmarshalling: ", err)
		os.Exit(1)
	}

	fmt.Println("test123: ", viper.Get("test123"))
}
