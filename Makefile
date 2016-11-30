# Standard Polaris Makefile

unittest:
	ginkgo -r -v -trace -cover -coverpkg=./...
	gover

coverage:
	goveralls -coverprofile=gover.coverprofile -service=codeship -repotoken $COVERALLS_TOKEN

install:
	time scp bin/bbqberry pi@pi:~/

build:
	time env GOOS=linux GOARCH=arm go build -o bin/bbqberry main.go

run:
	time ssh pi@pi ~pi/bbqberry -logtostderr=true

# Environment target sets up initial dependencies that are not checked into the repo.
environment:
	go get github.com/onsi/ginkgo/ginkgo                        # ginkgo BDD framework
	go get github.com/onsi/gomega                               # BDD matcher library
	go get github.com/modocache/gover                           # Code coverage aggregation
	go get github.com/mattn/goveralls                           # coveralls.io online code coverage viewer
	go get -u github.com/kardianos/govendor
	go get -u github.com/go-openapi/runtime
	go get -u github.com/go-swagger/go-swagger/cmd/swagger

encrypt:
	jet encrypt credentials.env.secret credentials.env.encrypted

clean:
	go clean
	find . -name "*.coverprofile" -delete
	rm -rf tmp/ cmd/ models/
	rm -rf restapi/operations restapi/doc.go restapi/embedded_spec.go restapi/server.go

swagger: clean
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

