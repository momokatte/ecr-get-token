package main

import (
	"bytes"
	"encoding/base64"
	"flag"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ecr"
)

func main() {
	var err error
	var region, accountId string
	flag.StringVar(&region, "region", "", "AWS region")
	flag.StringVar(&accountId, "account", "", "Account/Registry ID")
	flag.Parse()

	if len(region) == 0 {
		flag.Usage()
		os.Exit(1)
	}

	config := aws.Config{
		Region: &region,
		CredentialsChainVerboseErrors: aws.Bool(true),
		HTTPClient: &http.Client{
			Timeout: time.Minute,
		},
	}
	sess := session.Must(session.NewSession(&config))
	client := ecr.New(sess)
	var token *string
	token, err = getEcrToken(client, accountId)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(2)
	}
	fmt.Fprint(os.Stdout, *token)
}

func getEcrToken(c *ecr.ECR, accountId string) (*string, error) {
	req := &ecr.GetAuthorizationTokenInput{}
	if len(accountId) != 0 {
		req.RegistryIds = append(req.RegistryIds, &accountId)
	}
	resp, err := c.GetAuthorizationToken(req)
	if err != nil {
		return nil, err
	}
	if len(resp.AuthorizationData) == 0 {
		return nil, fmt.Errorf("No authorization tokens found")
	}
	var creds []byte
	if creds, err = base64.StdEncoding.DecodeString(*resp.AuthorizationData[0].AuthorizationToken); err != nil {
		return nil, err
	}
	// creds are user:password
	i := bytes.Index(creds, []byte(":"))
	if i < 0 {
		return nil, fmt.Errorf("Error parsing token")
	}
	// return the password
	return aws.String(string(creds[i+1:])), nil
}
