
.PHONY: deps
deps:
	go get github.com/golang/mock/gomock
	go install github.com/golang/mock/mockgen
	go mod vendor -v

.PHONY: clean_mocks
clean_mocks:
	rm -rf counterfeiter_mocks
	rm -rf gomock_mocks

.PHONY: mocks
mocks: clean_mocks
	go generate -v ./...

.PHONY: test
test: deps mocks
	go test -v ./...
