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
	msgcint, _ := strconv.Atoi(msgcstr)

	// print message with to be sent amount of messages and sqs queue url
	fmt.Println("start sending " + msgcstr + " messages to " + urlqueue)

	// setup a session
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	// create a session with sqs and set counter to 0
	svc := sqs.New(sess)
	totalCount := 0

	// start batch write
	for x := 0; totalCount < (msgcint); x++ {

		// set the message payloads
		msgs := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
		var msgBatch []*sqs.SendMessageBatchRequestEntry

		// iterate over the 10 messages
		for i := 0; i < len(msgs); i++ {

			// create an entry per message
			message := &sqs.SendMessageBatchRequestEntry{
				Id:          aws.String(`test_` + strconv.Itoa(i)),
				MessageBody: aws.String(msgs[i]),
			}

			// increase total count
			totalCount++

			// append the batch message
			msgBatch = append(msgBatch, message)
		}

		// set sqs batch entries and queue url from argv
		params := &sqs.SendMessageBatchInput{
			Entries:  msgBatch,
			QueueUrl: &urlqueue,
		}

		// send the batch message
		_, err := svc.SendMessageBatch(params)

		// if error found in sending, print it
		if err != nil {
			fmt.Println(err)
		}

		// print status per 1000 messages
		if totalCount%1000 == 0 {
			fmt.Println("sent " + strconv.Itoa(totalCount) + " messages ")
		}
	}

	// final - print total messages completed
	fmt.Println("finished - sent " + strconv.Itoa(totalCount) + " messages to queue")
}
