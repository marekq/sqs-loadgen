package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

// main function
func main() {

	urlqueue := ""
	msgcstr := ""
	bytecount := 0
	msgcint := 0

	// retrieve sqs queue url from argument 1
	if len(os.Args) > 1 {
		urlqueue = os.Args[1]
		fmt.Println("INFO - found SQS queue argument of " + urlqueue + "\n")

	} else {
		fmt.Println("ERROR - no SQS message queue specifiec in argument, exiting \n")
		os.Exit(3)
	}

	// retrieve message count from argument 2
	if len(os.Args) > 2 {
		msgcstr = os.Args[2]
		msgcint, _ = strconv.Atoi(msgcstr)
		fmt.Println("INFO - found count argument of " + msgcstr + " messages\n")

	} else {
		fmt.Println("WARNING - no SQS message count specified in argument, using default value of 1000\n")
		msgcint = 1000
	}

	if len(os.Args) > 3 {
		bytecount, _ = strconv.Atoi(os.Args[3])
		fmt.Println("INFO - found count per message argument of " + os.Args[3] + " bytes\n")

	} else {
		fmt.Println("WARNING - no byte count specified in argument, using default value of 100 bytes\n")
		bytecount = 100
	}

	// print message with to be sent amount of messages and sqs queue url
	fmt.Println("start sending " + msgcstr + " messages to " + urlqueue + "\n")

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
				Id:          aws.String(`test_` + strconv.Itoa(totalCount)),
				MessageBody: aws.String(strings.Repeat(msgs[i], bytecount)),
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

		modulo := 1000
		if msgcint < 1000 {
			modulo = 10
		} else if msgcint < 5000 {
			modulo = 100
		}

		// print status per 1000 messages
		if totalCount%modulo == 0 {
			fmt.Println("sent " + strconv.Itoa(totalCount) + " messages of bytesize " + strconv.Itoa(bytecount))
		}
	}

	// final - print total messages completed
	fmt.Println("finished - sent " + strconv.Itoa(totalCount) + " messages to queue")
}
