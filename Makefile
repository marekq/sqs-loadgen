.PHONY: build
build:
	sam build --parallel --cached

.PHONY: deploy
deploy:
	sam deploy