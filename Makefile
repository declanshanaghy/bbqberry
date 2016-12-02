# Standard Polaris Makefile

unittest: clean_coverage
	ginkgo -r -v -p --progress -trace -cover -coverpkg=./...
	gover
	cat gover.coverprofile | \
	    grep -v vendor | grep -v client | grep -v models | grep -v restapi | \
	    grep -v cmd | grep -v mocks | grep -v example | grep -v test \
	    > gover.coverprofile.sanitized

coverage_local:
	go tool cover -html=gover.coverprofile.sanitized -o cover.html

coverage:
	goveralls -coverprofile=gover.coverprofile.sanitized -service=codeship -repotoken V3p8U7YnvB2xRXYJVmWvrYFsvSXuPSyQx

install:
	time scp bin/bbqberry pi@pi:~/

build:
	time env GOOS=linux GOARCH=arm go build -o bin/bbqberry cmd/app-server/main.go

run:
	time ssh pi@pi ~pi/bbqberry --host=0.0.0.0 --port=8000

mock:
	mkdir -p tmp/vendor
	rm -rf mocks && mkdir -p mocks/mock_embd
	ln -Fs $(shell pwd)/vendor ./tmp/vendor/src
	GOPATH=$(shell pwd)/tmp/vendor:$$GOPATH \
	    mockgen github.com/kidoman/embd SPIBus > mocks/mock_embd/embd.go
	rm vendor/vendor || true

# Environment target sets up initial dependencies that are not checked into the repo.
environment:
	go get -u github.com/golang/mock/mockgen                        # Mock generator
	go get -u github.com/onsi/ginkgo/ginkgo                         # ginkgo BDD framework
	go get -u github.com/onsi/gomega                                # BDD matcher library
	go get -u github.com/modocache/gover                            # Code coverage aggregation
	go get -u github.com/mattn/goveralls                            # coveralls.io online code coverage viewer
	go get -u github.com/kardianos/govendor
	go get -u github.com/go-openapi/runtime
	go get -u github.com/go-swagger/go-swagger/cmd/swagger

encrypt:
	jet encrypt credentials.env.secret credentials.env.encrypted

clean_coverage:
	find . -name "*.coverprofile*" -delete

clean:
	go clean
	rm -rf tmp/ cmd/ models/
	rm -rf restapi/operations restapi/doc.go restapi/embedded_spec.go restapi/server.go

validate_swagger:
	swagger validate swagger.yml

swagger: validate_swagger clean
	rm -rf client/
	swagger generate server --name app --spec swagger.yml
	swagger generate client --name app --spec swagger.yml
	swagger generate support --name app --spec swagger.yml

dependencies:
	govendor fetch +missing
	govendor add +external
	govendor sync
	govendor remove +unused

codeship: clean
	jet steps $(XARGS)

