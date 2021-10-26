.PHONY: init lint

GO111MODULE=on
LINT_OPT := -E gofmt \
            -E golint \
            -E govet \
            -E gosec \
            -E unused \
            -E gosimple \
            -E structcheck \
            -E varcheck \
            -E ineffassign \
            -E deadcode \
            -E typecheck \
            -E misspell \
            -E whitespace \
            -E errcheck \
            --exclude '(comment on exported (method|function|type|const|var)|should have( a package)? comment|comment should be of the form)' \
            --timeout 5m

init:
	go mod download

lint:
	@type golangci-lint > /dev/null || go get -u github.com/golangci/golangci-lint/cmd/golangci-lint
	golangci-lint $(LINT_OPT) run ./...
