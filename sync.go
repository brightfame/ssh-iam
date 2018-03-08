package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/iam"
)

type syncUsers struct {
	IamGroup string `cli:"arg"`
}

func (r *syncUsers) Run() error {
	if r.IamGroup == "" {
		return fmt.Errorf("You must specify an IAM group")
	}

	svc := iam.New(session.New(), &aws.Config{})

	// get a list of users from the specified IAM group
	iamGroupRequest := &iam.GetGroupInput{
		GroupName: aws.String(r.IamGroup),
	}

	iamGroupResp, err := svc.GetGroup(iamGroupRequest)
	if err != nil {
		if iamerr, ok := err.(awserr.Error); ok && iamerr.Code() == "NoSuchEntity" {
			fmt.Printf("[WARN] No IAM group by name (%s) found", r.IamGroup)
			return nil
		}
		return fmt.Errorf("Error reading IAM Group %s: %s", r.IamGroup, err)
	}

	writeLog(fmt.Sprintf("Found IAM Group: %s", iamGroupResp))

	initKeysFile()

	// store the usernames, fingerprints and public ssh keys in memory

	// open the existing keys file and load each fingerprint + key into memory
	// create it if it doesn't exist

	// do a diff to see what to add and what to delete

	return nil
}
