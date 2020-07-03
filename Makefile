.PHONY: build	

dev:
	@docker run -it -v $(shell pwd):/go/src dahua-netsdk2:dev bash

build:
	@go build -o bin/traffic ./traffic
	@go build -o bin/example ./example

sync:
	@scp bin/example devserver2:~/netsdk/bin
	@scp bin/traffic devserver2:~/netsdk/bin
	
dbuild:
	@docker build -t dahua-netsdk2:dev .
