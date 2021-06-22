
GOFILES=`find . -name "*.go"`

fmt:
	@gofmt -s -w ${GOFILES}