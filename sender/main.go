package main

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/aws/aws-xray-sdk-go/xray"
)

// main function
func handler(ctx context.Context) {

	// create xray handler segment
	_, Seg1 := xray.BeginSubsegment(ctx, "handler")

	// get lambda environment variables for sqs queue
	urlqueue := os.Getenv("sqsqueue")
	os.Setenv("AWS_XRAY_CONTEXT_MISSING", "LOG_ERROR")

	// get message count to send
	msgcstr := os.Getenv("messagecount")
	msgcint, _ := strconv.Atoi(msgcstr)

	// get per message byte count
	bytecountstr := os.Getenv("messagebytes")
	bytecountint, _ := strconv.Atoi(bytecountstr)

	// get aws region and lambda memory setting
	region := os.Getenv("AWS_REGION")
	lambdamemory := os.Getenv("AWS_LAMBDA_FUNCTION_MEMORY_SIZE")

	// print message with to be sent amount of messages and sqs queue url
	fmt.Println("start sending " + msgcstr + " messages to " + urlqueue + " with payload size " + bytecountstr + " bytes\n")

	// get the current timestamp
	startts := int(time.Now().Unix())

	// setup a session
	sess := session.Must(session.NewSession(&aws.Config{
		Region: aws.String(region)},
	))

	// create a session with sqs and set counter to 0
	svc := sqs.New(sess)
	xray.AWS(svc.Client)

	// close handler subsegment
	Seg1.Close(nil)

	totalCount := 0

	// start batch write
	for x := 0; totalCount < (msgcint); x++ {

		// set the message payloads
		msgs := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
		var msgBatch []*sqs.SendMessageBatchRequestEntry

		// iterate over the 10 messages per batch
		for i := 0; i < len(msgs); i++ {

			// create an entry per message
			message := &sqs.SendMessageBatchRequestEntry{
				Id: aws.String(strconv.Itoa(totalCount + i)),

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

		// send the batch message traced with xray
		_, Seg2 := xray.BeginSubsegment(ctx, "handler")
		_, err := svc.SendMessageBatchWithContext(ctx, params)
		Seg2.Close(nil)

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

		// print status per x amount of messages
		if totalCount%modulo == 0 {

			// get current timestamp
			nowts := int(time.Now().Unix())
			diffts := (nowts - startts)

			// skip logs for first two seconds as these can have skewed metrics
			if diffts > 2 {

				sendrate := totalCount / diffts
				fmt.Println("sent " + strconv.Itoa(totalCount) + " messages (" + strconv.Itoa(sendrate) + "/sec) of " + bytecountstr + " bytes with memory size " + lambdamemory + " MB in " + strconv.Itoa(diffts) + " sec")
			}
		}
	}

	// final - print total messages completed
	fmt.Println("finished - sent " + strconv.Itoa(totalCount) + " messages to queue")

}

func main() {
	lambda.Start(handler)
}
