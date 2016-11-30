# BBQ Berry

Raspberry Pi controlled BBQ

mock:
	rm -rf mocks && mkdir mocks
	mkdir -p tmp/vendor
	if [ ! -s ./tmp/vendor/src ]; then ln -s $(shell pwd)/vendor ./tmp/vendor/src; fi;
	GOPATH=$(shell pwd)/tmp/vendor:$$GOPATH \
	    mockgen github.com/kidoman/embd SPIBus > mocks/embd.go
	#rm vendor/vendor/src