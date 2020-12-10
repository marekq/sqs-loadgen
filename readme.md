sqs-loadgen
===========

Generate random messages to your SQS queue using Go. 


Usage
-----

You can run the code on your local machine or in a Lambda function as follows;

'go run loadgensqs.go https://sqs.eu-west-1.amazonaws.com/1234567890/sqsqueue 10000 100'
'go run loadgensqs.go https://sqs.<region>.amazonaws.com/<account>/<queuename> <optional message count> <optional message byte size>'


Contact
-------

In case you have any suggestions, questions or remarks, please raise an issue or reach out to @marekq.
