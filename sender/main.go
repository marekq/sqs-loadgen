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

	// get lambda environment variables for sqs queue
	urlqueue := os.Getenv("sqsqueue")

	// get message count to send
	msgcstr := os.Getenv("messagecount")
	msgcint, _ := strconv.Atoi(msgcstr)

	// get per message byte count
	bytecountstr := os.Getenv("messagebytes")
	bytecountint, _ := strconv.Atoi(bytecountstr)

	// get aws region
	region := os.Getenv("AWS_REGION")

	// print message with to be sent amount of messages and sqs queue url
	fmt.Println("start sending " + msgcstr + " messages to " + urlqueue + " with payload size " + bytecountstr + " bytes\n")

	// setup a session
	sess := session.Must(session.NewSession(&aws.Config{
		Region: aws.String(region)},
	))

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
				Id: aws.String(`test_` + strconv.Itoa(totalCount+i)),

				// repeat the payload to meet the set byte size
				MessageBody: aws.String(strings.Repeat(msgs[i], bytecountint)),
			}

			// append the batch message
			msgBatch = append(msgBatch, message)
		}

		// increase total count by 10
		totalCount = totalCount + 10

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
			fmt.Println("sent " + strconv.Itoa(totalCount) + " messages of bytesize " + bytecountstr)
		}
	}

	// final - print total messages completed
	fmt.Println("finished - sent " + strconv.Itoa(totalCount) + " messages to queue")
}
