#!/bin/bash

# Required for building and testing
go get -u github.com/fzipp/gocyclo                              # Calculates cyclomatic complexities
go get -u github.com/golang/lint                                # Code linter
go get -u github.com/gordonklaus/ineffassign                    # Detects ineffectual assignments
go get -u github.com/github.com/client9/misspell                # Detects common mispellings
go get -u github.com/golang/mock/mockgen                        # Mock generator
go get -u github.com/onsi/ginkgo/ginkgo                         # ginkgo BDD framework
go get -u github.com/onsi/gomega                                # BDD matcher library
go get -u github.com/modocache/gover                            # Code coverage aggregation
go get -u github.com/mattn/goveralls                            # coveralls.io online code coverage viewer
go get -u github.com/kardianos/govendor
go get -u github.com/go-openapi/runtime
go get -u github.com/go-swagger/go-swagger/cmd/swagger

# Required for code quality analysis
go get -u golang.org/x/tools/cmd/goimports                      # Fixes imports
go get -u golang.org/x/tools/cmd/gofmt                          # Reformats go source code
