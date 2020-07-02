.PHONY: build	

dev:
	@docker run -it -v $(shell pwd):/go/src dahua-netsdk2:dev bash

build:
	@go build -o bin/traffic ./traffic
	@go build -o bin/example ./example

sync:
	@scp bin/dahuanet-linux devserver2:~/dahuasdk
	
dbuild:
	@docker build -t dahua-netsdk2:dev .
