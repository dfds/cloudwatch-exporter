package main

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudwatchlogs"

	"fmt"
	"os"
)

func main() {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	svc := cloudwatchlogs.New(sess)

	// Get up to the last 100 log events for LOG-STREAM-NAME
	// in LOG-GROUP-NAME:
	resp, err := svc.GetLogEvents(&cloudwatchlogs.GetLogEventsInput{
		Limit:         aws.Int64(100),
		LogGroupName:  aws.String("/aws/eks/hellman/cluster"),
		LogStreamName: aws.String("kube-apiserver-audit-e331bcb7ff1ce984e4a5bdbe46adf243"),
	})
	if err != nil {
		fmt.Println("Got error getting log events:")
		fmt.Println(err.Error())
		os.Exit(1)
	}

	fmt.Println("Event messages for stream /aws/eks/hellman/cluster in log group kube-apiserver-audit-e331bcb7ff1ce984e4a5bdbe46adf243:")

	// gotToken := ""
	// nextToken := ""

	for _, event := range resp.Events {
		// gotToken = nextToken
		// nextToken = *resp.NextForwardToken

		// if gotToken == nextToken {
		// 	break
		// }

		fmt.Println("  ", *event.Message)
	}
}
