default: test lint

ci: get lint test

get:
ifeq ($(GO_VERSION), go1.6)
	echo "$(GO_VERSION) is not a supported Go release. Skipping golint."
else ifeq ($(GO_VERSION), go1.7)
	echo "$(GO_VERSION) is not a supported Go release. Skipping golint."
else ifeq ($(GO_VERSION), go1.8)
	echo "$(GO_VERSION) is not a supported Go release. Skipping golint."
else
	go get -u golang.org/x/lint/golint
endif
	go get -t -v ./...

lint:
ifeq ($(GO_VERSION), go1.6)
	echo "$(GO_VERSION) is not a supported Go release. Skipping golint."
else ifeq ($(GO_VERSION), go1.7)
	echo "$(GO_VERSION) is not a supported Go release. Skipping golint."
else ifeq ($(GO_VERSION), go1.8)
	echo "$(GO_VERSION) is not a supported Go release. Skipping golint."
else
	golint ./...
endif

vet:
	go vet ./...

test:
	go test -cover -v ./...

test-v:
	go test -check.vv -cover ./...

