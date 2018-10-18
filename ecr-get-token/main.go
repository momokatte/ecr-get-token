package main

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
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
	var region, accountId, format string
	flag.StringVar(&region, "region", "", "AWS region")
	flag.StringVar(&accountId, "account", "", "Account/Registry ID")
	flag.StringVar(&format, "format", "porcelain", "Output format")
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
	var sess *session.Session
	if sess, err = session.NewSession(&config); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(2)
	}
	client := ecr.New(sess)
	var token *string
	if token, err = getToken(client, accountId); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(2)
	}
	var out string
	if out, err = formatToken(*token, format); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(2)
	}
	fmt.Fprint(os.Stdout, out)
}

func getToken(c *ecr.ECR, accountId string) (*string, error) {
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
	return resp.AuthorizationData[0].AuthorizationToken, nil
}

func parseToken(token string) (user string, pass string, err error) {
	var b []byte
	if b, err = base64.StdEncoding.DecodeString(token); err != nil {
		return
	}
	parts := bytes.Split(b, []byte(":"))
	if len(parts) != 2 {
		err = errors.New("Error parsing credentials")
		return
	}
	return string(parts[0]), string(parts[1]), nil
}

func formatToken(token string, format string) (out string, err error) {
	if format == "base64" {
		return token, nil
	}
	var user, pass string
	if user, pass, err = parseToken(token); err != nil {
		return
	}
	switch format {
	case "porcelain":
		out = fmt.Sprintf("%s %s\n", user, pass)
	case "username":
		out = user
	case "password":
		out = pass
	case "shell":
		out = fmt.Sprintf("USERNAME=%s\nPASSWORD=%s\n", user, pass)
	case "json":
		m := map[string]string{
			"username": user,
			"password": pass,
		}
		var b []byte
		if b, err = json.Marshal(m); err != nil {
			return
		}
		out = string(b)
	case "yaml":
		out = fmt.Sprintf("username: \"%s\"\npassword: \"%s\"\n", user, pass)
	default:
		err = errors.New("Unsupported format")
	}
	return
}
