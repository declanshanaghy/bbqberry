# BBQBerry Makefile


export INFLUXDB=bbqberry_test
export MONGODB=bbqberry_test
export STUB=yes

APP_NAME := $(shell basename $(PWD))

SWAGGER_YML := swagger.yml
SWAGGER_SERVER ?= true
SWAGGER_CLIENT ?= true
SWAGGER_SUPPORT ?= true

unittest_bbqberry: create_influxdb
	@echo "Running Unit Tests..."
	@ginkgo -r -v -p -randomizeAllSpecs -randomizeSuites \
	    -progress -trace -cover -covermode atomic -skipPackage "./cmd" $(XARGS)

create_influxdb:
	@echo "Creating DB..."
	@go run cmd/influxdb/create_database.go

build_bbqberry:
	@echo "Building BBQBerry..."
	@env GOOS=linux GOARCH=arm go build -o tmp/bin/bbqberry cmd/bbqberry-server/main.go

upload_ftp: stop_bbqberry build_bbqberry
	@echo "Uploading via FTP..."
	@time scripts/upload_ftp.sh
	@echo "Upload complete"
	@make start_remote

upload_scp: stop_bbqberry build_bbqberry
	@echo "Uploading via SCP..."
	@time scp -p tmp/bin/bbqberry pi@bbqberry-gaff:~/deploy/bbqberry
	@echo "Upload complete"
	@make start_remote

upload_config:
	@echo "Uploading config..."
	@scp -p etc/bbqberry-supervisord.conf pi@bbqberry-gaff:~/go/src/github.com/declanshanaghy/bbqberry/etc/bbqberry-supervisord.conf
	@echo "Restarting..."
	@ssh pi@bbqberry-gaff sudo supervisorctl reread
	@ssh pi@bbqberry-gaff sudo supervisorctl update
	@ssh pi@bbqberry-gaff sudo supervisorctl restart bbqberry
	@echo "Upload complete"

stop_bbqberry:
	@echo "Stopping BBQBerry..."
	@ssh pi@bbqberry-gaff sudo supervisorctl stop bbqberry

restart_remote:
	@echo "Restarting BBQBerry remote..."
	@ssh pi@bbqberry-gaff sudo supervisorctl restart bbqberry

start_remote:
	@echo "Restarting BBQBerry remote..."
	@ssh pi@bbqberry-gaff sudo supervisorctl start bbqberry

run_remote: upload_scp
	@echo "Running BBQBerry remote..."
	@ssh pi@bbqberry-gaff ~/deploy/bbqberry --host=0.0.0.0 --port=8080 \
	    --static=/home/pi/go/src/github.com/declanshanaghy/bbqberry/static

run_deployed:
	@echo "Running BBQBerry deployed..."
	@$(OUTBIN)/bbqberry --host=0.0.0.0 --port=8080 \
	    --static=/home/pi/go/src/github.com/declanshanaghy/bbqberry/static

sync_web:
	@echo "Syncing BBQBerry webapp..."
	rsync -rv ./static/ pi@bbqberry-gaff:~/go/src/github.com/declanshanaghy/bbqberry/static/

clean_swagger:
#
#__Deletes all swagger generated files__
#
	@find cmd/$(APP_NAME)-server ! -name 'main_test.go' -type f -exec rm -f {} +
	@rm -rf models/ client/
	@rm -rf restapi/operations restapi/doc.go restapi/embedded_spec.go restapi/server.go
	@printf "Cleaned swagger\n"

validate_swagger:
#
#__Validates swagger.yml__
#
# (if it exists)
#
	@if [ -f $(SWAGGER_YML) ]; then \
	    swagger validate $(SWAGGER_YML); \
	fi

swagger: validate_swagger
#
#__Generates swagger source files__
#
#   The following components are generated:
#* Server
#* API Client library
#* Support files (the main function and the api builder)
#
	@if [ -f $(SWAGGER_YML) ]; then \
	    if [ "$(SWAGGER_SERVER)" = "true" ]; then \
		    swagger generate server --name $(APP_NAME) --spec $(SWAGGER_YML); \
		else \
			printf "Swagger server disabled\n"; \
        fi; \
	    if [ "$(SWAGGER_SUPPORT)" = "true" ]; then \
			swagger generate support --name $(APP_NAME) --spec $(SWAGGER_YML); \
		else \
			printf "Swagger support disabled\n"; \
		fi; \
	    if [ "$(SWAGGER_CLIENT)" = "true" ]; then \
			swagger generate client --name $(APP_NAME) --spec $(SWAGGER_YML) \
		else \
			printf "Swagger client disabled\n"; \
		fi; \
	fi

generate: swagger
#
#__Generates source code__
#
	@printf "Code generation completed\n"

