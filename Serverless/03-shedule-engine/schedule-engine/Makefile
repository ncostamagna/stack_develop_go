.PHONY: build clean deploy

build:
	env GOOS=linux go build -ldflags="-s -w" -o bin/schedule schedule/main.go


clean:
	rm -rf ./bin ./vendor Gopkg.lock

deploy: clean build
	sls deploy --verbose
