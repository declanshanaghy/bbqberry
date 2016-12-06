# Standard Polaris Makefile

include skel/Makefile

APP_NAME = bbqberry

mock:
	mkdir -p tmp/vendor
	rm -rf mocks && mkdir -p mocks/mock_embd
	ln -fs $(shell pwd)/vendor ./tmp/vendor/src
	    GOPATH=$(shell pwd)/tmp/vendor:$$GOPATH \
	mockgen github.com/kidoman/embd SPIBus > mocks/mock_embd/embd.go
	rm vendor/vendor || true