package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

// main function
func main() {

	// retrieve sqs queue url from argument 1
	urlqueue := os.Args[1]

	// retrieve amount of messages to send from argument 2
	msgcstr := os.Args[2]

	// convert message count string to int
	msgcint, _ := strconv.Atoi(msgcstr)

	// print message with to be sent amount of messages and sqs queue url
	fmt.Println("start sending " + msgcstr + " messages to " + urlqueue)

	// setup a session
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	// create a session with sqs
	svc := sqs.New(sess)

	// create trace for every message group
	for tot := 1; tot < (msgcint); tot++ {

		// send one message to the queue
		_, err := svc.SendMessage(&sqs.SendMessageInput{
			MessageBody: aws.String(msgcstr),
			QueueUrl:    aws.String(urlqueue),
		})

		// print an error if message sending failed
		if err != nil {
			fmt.Println(err)
		}

		if tot%10 == 0 {
			fmt.Println("sent " + strconv.Itoa(tot) + " messages ")
		}
	}

	// print total messages completed
	fmt.Println("finished - sent " + msgcstr + " messages to queue")

}
