package main

import (
	"fmt"
	"regexp"
	"sort"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/awsutil"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/iam"
	"github.com/dynport/gocli"
)

type listKeys struct {
	Username string `cli:"arg"`
}

func (r *listKeys) Run() error {
	if r.Username == "" {
		return fmt.Errorf("You must specify a username")
	}
	// Create an EC2 service object in the "eu-west-1" region
	// Note that you can also configure your region globally by
	// exporting the AWS_REGION environment variable
	svc := iam.New(session.New(), &aws.Config{})

	// get the keys for user
	listKeyRequest := &iam.ListSSHPublicKeysInput{
		//UserName: &aws.String(d.Id()),
		UserName: aws.String(r.Username),
	}

	listKeyResp, err := svc.ListSSHPublicKeys(listKeyRequest)
	if err != nil {
		if iamerr, ok := err.(awserr.Error); ok && iamerr.Code() == "NoSuchEntity" {
			fmt.Printf("[WARN] No IAM user by name (%s) found", r.Username)
			return nil
		}
		return fmt.Errorf("Error reading IAM User %s: %s", r.Username, err)
	}

	writeLog(fmt.Sprintf("Found %d Public SSH keys", len(listKeyResp.SSHPublicKeys)))

	var keys []*iam.SSHPublicKey
	publicKeys, _ := awsutil.ValuesAtPath(listKeyResp, "SSHPublicKeys[]")
	for _, publicKey := range publicKeys {
		k := publicKey.(*iam.SSHPublicKeyMetadata)
		writeLog(fmt.Sprintf("Requesting Public SSH Key for Key ID: %s", *k.SSHPublicKeyId))

		// request each key
		getKeyRequest := &iam.GetSSHPublicKeyInput{
			Encoding:       aws.String("SSH"), // SSH or PEM
			UserName:       aws.String(r.Username),
			SSHPublicKeyId: aws.String(*k.SSHPublicKeyId),
		}

		getKeyResp, err := svc.GetSSHPublicKey(getKeyRequest)
		if err != nil {
			if iamerr, ok := err.(awserr.Error); ok && iamerr.Code() == "NoSuchEntity" {
				fmt.Printf("[WARN] Could not get SSH Public Key for IAM user by name (%s)", r.Username)
				return nil
			}
			return fmt.Errorf("Error reading public key (%s) for User %s: %s", *k.SSHPublicKeyId, r.Username, err)
		}
		writeLog(fmt.Sprintf("Found keys: %s", getKeyResp))

		keys = append(keys, getKeyResp.SSHPublicKey)
	}

	// format table
	// id launch_time ami name ip type role
	t := gocli.NewTable()
	t.Header("id", "fingerprint", "status", "upload_date", "username")

	// print keys
	if len(keys) > 0 {
		for _, key := range keys {
			t.Add(
				key.SSHPublicKeyId,
				key.Fingerprint,
				key.Status,
				key.UploadDate.Format("2006-01-02T15:04"),
				key.UserName,
			)
		}
	}

	t.SortBy = 1
	sort.Sort(sort.Reverse(t))
	fmt.Println(t)

	return nil
}

func validateAwsIamUserName(v interface{}, k string) (ws []string, errors []error) {
	value := v.(string)
	if !regexp.MustCompile(`^[0-9A-Za-z=,.@\-_+]+$`).MatchString(value) {
		errors = append(errors, fmt.Errorf(
			"only alphanumeric characters, hyphens, underscores, commas, periods, @ symbols, plus and equals signs allowed in %q: %q",
			k, value))
	}
	return
}

func truncate(s string, length int) string {
	if len(s) > length {
		return s[0:length]
	}
	return s
}

func p2s(in *string) string {
	if in == nil {
		return ""
	}
	return *in
}
