.PHONY: build	

dev:
	@docker run -it -v $(shell pwd):/go/src -p 7001:7000 dahua-netsdk2:dev bash

build:
	@go build -o bin/traffic ./traffic
	@go build -o bin/server ./example/server
	@go build -o bin/example ./example/alarm

sync:
	@scp bin/example devserver2:~/netsdk/bin
	@scp bin/server  devserver2:~/netsdk/bin
	@scp bin/traffic devserver2:~/netsdk/bin
	
dbuild:
	# @docker build -t dahua-netsdk2:dev .
	@docker buildx build --platform=linux/amd64 -t dahua-netsdk2:dev .
