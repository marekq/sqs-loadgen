.PHONY: build
build:
	sam build SQSSender --cached
	sam build -u SQSReceiver --cached

.PHONY: deploy
deploy:
	sam deploy