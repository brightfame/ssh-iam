package main

import (
	"os"
)

var configPath = "/home/ubuntu/.ssh-iam"

func initKeysFile() error {
	// create config dir if it doesn't exist
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		err := os.MkdirAll(configPath, 0755)
		if err != nil {
			return err
		}
	}

	configFilePath := configPath
	f, err := os.OpenFile(configFilePath, os.O_RDWR|os.O_CREATE|os.O_EXCL, 0600)
	if err != nil {
		return err
	}
	f.WriteString("foo")

	return nil
}
