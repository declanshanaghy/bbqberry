# BBQBerry Makefile


export STUB=yes


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
