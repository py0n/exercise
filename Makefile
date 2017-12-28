install-extra:
	go get -u github.com/golang/lint/golint
	go get -u github.com/jstemmer/gotags
	go get -u github.com/k0kubun/pp
	go get -u github.com/monochromegane/the_platinum_searcher/...
	go get -u github.com/motemen/ghq
	go get -u github.com/motemen/gore
	go get -u github.com/nsf/gocode
	go get -u github.com/rogpeppe/godef
	go get -u golang.org/x/tools/cmd/godoc
	go get -u golang.org/x/tools/cmd/goimports
	go get -u golang.org/x/tools/cmd/gorename
	go get -u golang.org/x/tools/cmd/guru

## setup tools
setup:
	go get github.com/golang/dep/cmd/dep

## install dependency
deps: setup
	$(HOME)/go-packages/bin/dep ensure

## run tests
test: deps
	go test ./project/euler

## lint
lint:
	go vet ./project/euler
	golint ./project/euler

## show help
help:
	@make2help $(MAKEFILE_LIST)
