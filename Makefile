build:
	go build -tags netgo -a -v -o ./bin/rootly ./cmd/rootly/

docker-build:
	docker build -t rootlyhub/cli .

docker-push:
	docker push rootlyhub/cli

clean:
	rm -r ./bin

test:
	go test -count=1 -v ./...

lint:
	golangci-lint run
	hadolint Dockerfile
	goreleaser check
