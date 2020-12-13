sqs-loadgen
===========

Generate random messages to your SQS queue using Go. The stack deploys a Lambda SQS generator, a Lambda SQS receiver and an SQS queue. 


Usage
-----

You can deploy the SQS generator and receiver using AWS SAM. Run the following commands on your local machine to build and deploy the stack to AWS;

```
make
sam deploy -g
```

After the stack is deployed, you can invoke the 'SQSSender' Lambda to generate traffic. 


Contact
-------

In case you have any suggestions, questions or remarks, please raise an issue or reach out to @marekq.
