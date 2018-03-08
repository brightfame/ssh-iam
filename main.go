package main

import (
	"log"
	"os"

	"github.com/dynport/dgtk/cli"
)

var logger = log.New(os.Stderr, "", 0)

func main() {
	router := cli.NewRouter()
	router.Register("install", &installClient{}, "Install the service")
	router.Register("list-keys", &listKeys{}, "List a Users Public SSH keys")
	router.Register("get-keys", &getKeys{}, "Output a Users Public SSH keys")
	router.Register("sync", &syncUsers{}, "Sync IAM Users")

	//router.Register("sync", &hostsList{}, "Sync IAM Users")
	switch err := router.RunWithArgs(); err {
	case nil, cli.ErrorHelpRequested, cli.ErrorNoRoute:
		// ignore
		return
	default:
		logger.Fatal(err)
	}
}

func writeLog(str string) error {
	if os.Getenv("BF_DEBUG") == "" {
		return nil
	}

	logger.Println(str)
	return nil
}
