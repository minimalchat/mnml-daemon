# Commands
GO_CMD = go
LINT_CMD = $(GOPATH)/bin/golint

# Directories
PACKAGE = github.com/minimalchat/mnml-daemon
SRC = $(GOPATH)/src/$(PACKAGE)
DIST = $(GOPATH)/bin

.PHONY: build lint

build: lint test coverage clean compile

run: lint test go

lint:
	$(LINT_CMD) $(SRC)
	# $(LINT_CMD) $(SRC) $(TEST)

test:
	cd $(SRC)
	$(DIST)/overalls -project=$(PACKAGE) -covermode=count
	# $(GO_CMD) test -v -coverpkg ./... -covermode=count -coverprofile=coverage.out ./...
	# $(GOPATH)/bin/goveralls -coverprofile=coverage.out -service=travis-ci -repotoken $(COVERALLS_TOKEN)

coverage:
	cd $(SRC)
	$(GOPATH)/bin/goveralls -coverprofile=overalls.coverprofile -service=travis-ci -repotoken $(COVERALLS_TOKEN)

clean:
	rm -rf $(DIST)/mnml-daemon

compile:
	mkdir -p $(DIST)
	cd $(SRC)
	$(GO_CMD) build -o $(DIST)/mnml-daemon

go:
	cd $(SRC)
	$(GO_CMD) run main.go
