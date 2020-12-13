.PHONY: build
build:
	sam build SQSSender 
	sam build -u SQSReceiver

.PHONY: deploy
deploy:
	sam deploy