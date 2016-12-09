# BBQBerry Makefile

include skel/Makefile

<<<<<<< HEAD
unittest_bbqberry: create_influxdb unittest
	@echo "Done"

create_influxdb:
	go run cmd/influxdb/create_database.go

mock:
	mkdir -p tmp/vendor
	rm -rf mocks && mkdir -p mocks/mock_embd
	ln -fs $(shell pwd)/vendor ./tmp/vendor/src
	    GOPATH=$(shell pwd)/tmp/vendor:$$GOPATH \
	mockgen github.com/kidoman/embd SPIBus > mocks/mock_embd/embd.go
	rm vendor/vendor || true

upload:
	echo "Uploading"
	env OUTBIN=$(OUTBIN) APP_NAME=$(APP_NAME) scripts/upload_ftp.sh
