FUNCNAME=.

install-extra:
	go get -u github.com/golang/lint/golint
	go get -u github.com/jstemmer/gotags
	go get -u github.com/k0kubun/pp
	go get -u github.com/monochromegane/the_platinum_searcher/...
	go get -u github.com/motemen/ghq
	go get -u github.com/motemen/gore
	go get -u github.com/nsf/gocode
	go get -u github.com/rogpeppe/godef
	go get -u golang.org/x/tools/cmd/...

## setup tools
setup:
	go get -u github.com/Songmu/make2help/cmd/make2help
	go get -u github.com/golang/dep/cmd/dep
	go get -u github.com/golang/lint/golint

## install dependency
deps: setup
	$(HOME)/go-packages/bin/dep ensure

## run tests - `make test FUNCNAME=PE0034` とする事で `PE0034` を含むテストのみ実行できる
test: deps
	go test ./project/euler -run ${FUNCNAME}

## lint
lint: setup
	go vet ./project/euler
	golint ./project/euler

## show help
help:
	@make2help $(MAKEFILE_LIST)
