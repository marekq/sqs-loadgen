sqs-loadgen
===========

Generate random messages to your SQS queue using Go to test SQS batching features. The stack deploys a Lambda SQS generator, a Lambda SQS receiver and an SQS queue using AWS SAM. 


Usage
-----

You can deploy the SQS generator and receiver using AWS SAM. Run the following commands on your local machine to build and deploy the stack to AWS;

```
make init       (run the first time to create the SAM config)
make deploy     (once SAM config is present)
```

After the stack is deployed, you can invoke the 'SQSSender' Lambda to generate traffic. It will use the default settings from environment variables. 


Contact
-------

In case you have any suggestions, questions or remarks, please raise an issue or reach out to @marekq.
