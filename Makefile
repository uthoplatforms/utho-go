# make publish tag=0.1.30
.PHONY: publish
publish: tidy
	git tag v$(tag)
	git push origin v$(tag)
	GOPROXY=proxy.golang.org go list -m github.com/uthoplatforms/utho-go@v$(tag)

.PHONY: test
test:
	go test -v ./...

.PHONY: test/count
test/count:
	go test -v ./... | grep -c RUN

.PHONY: test/cover
test/cover:
	go test -v -race -buildvcs -coverprofile=/tmp/coverage.out ./...
	go tool cover -html=/tmp/coverage.out

.PHONY: tidy
tidy:
	go fmt ./...
	go mod tidy -v