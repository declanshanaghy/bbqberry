# Standard Polaris Makefile


unittest: code_report
	ginkgo -r -v -p --progress -trace -cover -coverpkg=./...
	gover
	cat gover.coverprofile | \
	    grep -v vendor | grep -v client | grep -v models | grep -v restapi | \
	    grep -v cmd | grep -v mocks | grep -v example | grep -v test \
	    > gover.coverprofile.sanitized

coverage_local: unittest
	go tool cover -html=gover.coverprofile.sanitized -o cover.html

coverage: unittest
	goveralls -coverprofile=gover.coverprofile.sanitized -service=codeship -repotoken V3p8U7YnvB2xRXYJVmWvrYFsvSXuPSyQx

install: swagger
	go install -v ./...
	cp $$GOPATH/bin/app-server /tmp/bin

build: swagger mock
	go build -o bin/bbqberry cmd/app-server/main.go

build_arm:
	env GOOS=linux GOARCH=arm go build -o bin/bbqberry cmd/app-server/main.go

build_docker:
	docker build -f Dockerfile-app -t polarishq/$(shell basename $(shell pwd)) .

mock:
	mkdir -p tmp/vendor
	rm -rf mocks && mkdir -p mocks/mock_embd
	ln -fs $(shell pwd)/vendor ./tmp/vendor/src
	GOPATH=$(shell pwd)/tmp/vendor:$$GOPATH \
	    mockgen github.com/kidoman/embd SPIBus > mocks/mock_embd/embd.go
	rm vendor/vendor || true

code_report: build
	./scripts/code_analysis.sh; \
	    if [ "$$?" == "0" ]; then \
	        echo "Code Report passed"; \
	    else \
	        echo "Code report failed" && exit 1; \
	    fi

code_quality:
	./scripts/code_quality.sh

# Environment target sets up initial dependencies that are not checked into the repo.
environment:
	./scripts/setup_environment.sh

encrypt:
	jet encrypt dockercfg.secret.json dockercfg.json.encrypted

clean_coverage:
	find . -name "*.coverprofile*" -delete
	find . -name cover.html -delete

clean_vendor:
	find ./vendor -d 1 -type dir | xargs rm -rf

clean_swagger:
	rm -rf cmd/ models/
	rm -rf restapi/operations restapi/doc.go restapi/embedded_spec.go restapi/server.go

clean: clean_coverage clean_swagger
	go clean
	rm -rf tmp/

validate_swagger:
	swagger validate swagger.yml

swagger: validate_swagger clean
	rm -rf client/
	swagger generate server --name app --spec swagger.yml
	swagger generate client --name app --spec swagger.yml
	swagger generate support --name app --spec swagger.yml

dependencies: clean_vendor
	govendor fetch +missing
	govendor add +external
	govendor sync
	govendor remove +unused

codeship: clean
	jet steps $(XARGS)

