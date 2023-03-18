GOCMD=go
GOTEST=$(GOCMD) test
EXPORT_RESULT?=FALSE
SOURCES=$(shell find . -name '*.go' -not -name '*_test.go' -not -name "main.go")
BIN_DIR:=bin

.PHONY: clean pre cli goconvey test coverage help
.DEFAULT_GOAL := help

##################################### Binary #####################################

clean: ## Clean the bin directory
	rm -rf $(BIN_DIR)
	rm -f ./checkstyle-report.xml checkstyle-report.xml yamllint-checkstyle.xml

pre: clean ## Create the bin directory
	mkdir -p $(BIN_DIR)

cli: pre $(BIN_DIR)/cdn ## Build the CLI binary

$(BIN_DIR)/%: cmd $(SOURCES)
	$(GOCMD) build -o $@ $</*.go

$(BIN_DIR):
	mkdir -p $@

##################################### Test #####################################

goconvey: ## Run goconvey
	go install github.com/smartystreets/goconvey@latest
	goconvey -port 1234

test: ## Run the tests
ifeq ($(EXPORT_RESULT), TRUE)
	$(GOCMD) install github.com/jstemmer/go-junit-report/v2@latest
	$(eval OUTPUT_OPTIONS = | go-junit-report -iocopy -set-exit-code -out junit-report.xml)
endif
	$(GOCMD) clean -testcache
	$(GOTEST) -v -race ./... $(OUTPUT_OPTIONS)

coverage: $(wildcard **/**/*.go) ## Run the tests and generate the coverage reports
	$(GOTEST) -cover -covermode=count -coverprofile=coverage.out ./...
ifeq ($(EXPORT_RESULT), TRUE)
	$(GOCMD) install github.com/axw/gocov/gocov@latest
	$(GOCMD) install github.com/AlekSi/gocov-xml@latest
	$(GOCMD) tool cover -html=coverage.out -o coverage.html
	gocov convert coverage.out | gocov-xml > coverage.xml
	rm -f coverage.out
endif

##################################### Help #####################################

help: ## Show this help
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' Makefile | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'
