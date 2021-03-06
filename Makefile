# BBQBerry Makefile


export INFLUXDB=bbqberry_test
export MONGODB=bbqberry_test
export STUB=yes

APP_NAME := $(shell basename $(PWD))

SWAGGER_YML := swagger.yml
SWAGGER_SERVER ?= true
SWAGGER_CLIENT ?= true
SWAGGER_SUPPORT ?= true

unittest_bbqberry:
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

upload_scp: stop_bbqberry build_bbqberry
	@echo "Uploading via SCP..."
	@time scp -p tmp/bin/bbqberry pi@bbqberry-gaff:~/deploy/bbqberry
	@echo "Upload complete"

upload_config:
	@echo "Uploading config..."
	@scp -rp etc/systemd/ pi@bbqberry-gaff:/home/pi/go/src/github.com/declanshanaghy/bbqberry/etc/
	@ssh pi@bbqberry-gaff sudo cp /home/pi/go/src/github.com/declanshanaghy/bbqberry/etc/systemd/bbqberry.service /usr/lib/systemd/system/bbqberry.service
	@ssh pi@bbqberry-gaff sudo cp /home/pi/go/src/github.com/declanshanaghy/bbqberry/etc/systemd/systemd-shutdown/* /usr/lib/systemd/system-shutdown/
	@echo "Upload complete"

reload_config:
	@echo "Reloading config..."
	@ssh pi@bbqberry-gaff sudo systemctl daemon-reload
	@ssh pi@bbqberry-gaff sudo systemctl restart bbqberry
	@echo "Reload complete"

stop_bbqberry:
	@echo "Stopping BBQBerry..."
	@ssh pi@bbqberry-gaff sudo systemctl stop bbqberry

restart_remote:
	@echo "Restarting BBQBerry remote..."
	@ssh pi@bbqberry-gaff sudo systemctl restart bbqberry

start_remote:
	@echo "Restarting BBQBerry remote..."
	@ssh pi@bbqberry-gaff sudo systemctl start bbqberry

run_remote: upload_scp
	@echo "Running BBQBerry remote..."
	@ssh pi@bbqberry-gaff ~/deploy/bbqberry --host=0.0.0.0 --port=8080 \
	    --static=/home/pi/go/src/github.com/declanshanaghy/bbqberry/static

run_deployed:
	@echo "Running BBQBerry deployed..."
	@$(OUTBIN)/bbqberry --host=0.0.0.0 --port=8080 \
	    --static=/home/pi/go/src/github.com/declanshanaghy/bbqberry/static

sync_web:
	@echo "Syncing entire BBQBerry webapp..."
	@rsync -rv ./static/ pi@bbqberry-gaff:~/go/src/github.com/declanshanaghy/bbqberry/static/

sync_views:
	@echo "Syncing BBQBerry views..."
	@scp -r ./static/bbqberry/app/views/ pi@bbqberry-gaff:~/go/src/github.com/declanshanaghy/bbqberry/static/bbqberry/app/

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

swagger: validate_swagger clean_swagger
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

lambda_BBQBerry:
	@cd alexa/BBQBerry && zip -r ../../tmp/BBQBerry.zip .
	@aws lambda update-function-code --function-name BBQBerry --zip-file fileb://tmp/BBQBerry.zip

lambda_BBQBerrySmartHome:
	@cd alexa/BBQBerrySmartHome && zip -r ../../tmp/BBQBerrySmartHome.zip .
	@aws lambda update-function-code --function-name BBQBerrySmartHome --zip-file fileb://tmp/BBQBerrySmartHome.zip

