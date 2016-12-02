#!/bin/bash -x

# Required for building and testing
go get -u github.com/go-openapi/runtime                         # Open-API runtime libs
go get -u github.com/go-swagger/go-swagger/cmd/swagger          # Swagger code generator
go get -u github.com/golang/mock/mockgen                        # Mock generator
go get -u github.com/onsi/ginkgo/ginkgo                         # ginkgo BDD framework
go get -u github.com/modocache/gover                            # Code coverage aggregation
go get -u github.com/mattn/goveralls                            # coveralls.io online code coverage viewer
go get -u github.com/kardianos/govendor                         # GoVendor runtime
go get -u github.com/jessevdk/go-flags                          # Runtime flag parser

# Required for code standards reporting (Analysis reporting only, no automatic source changes)
go get -u github.com/fzipp/gocyclo                              # Calculates cyclomatic complexities
go get -u github.com/gordonklaus/ineffassign                    # Detects ineffectual assignments
go get -u github.com/golang/lint/golint                         # Code linter
go get -u github.com/client9/misspell/cmd/misspell              # Detects common mispellings

# Required for code quality analysis (Automatically change files on disk to be committed by dev)
go get -u golang.org/x/tools/cmd/goimports                      # Fixes imports