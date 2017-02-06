# BBQBerry Makefile

include skel/Makefile

OUTBIN=~/deploy

export STUB=yes


unittest_bbqberry: create_influxdb unittest
	@echo "Done"

create_influxdb:
	go run cmd/influxdb/create_database.go

build_bbqberry:
	@echo "Building BBQBerry..."
	env GOOS=linux GOARCH=arm make build OUTBIN=tmp/bin

upload_ftp: kill build_bbqberry
	@echo "Uploading via FTP..."
	@time env OUTBIN=$(OUTBIN) APP_NAME=$(APP_NAME) scripts/upload_ftp.sh
	@echo "Upload complete"

upload_scp: kill build_bbqberry
	@echo "Uploading via SCP..."
	@time scp -p $(OUTBIN)/$(APP_NAME) pi@bbqberry-gaff:~/deploy/bbqberry
	@echo "Upload complete"

kill:
	ssh pi@bbqberry-gaff killall bbqberry

run_remote: upload_scp
	ssh pi@bbqberry-gaff $(OUTBIN)/bbqberry --host=0.0.0.0 --port=8888 --static=/home/pi/go/src/github.com/declanshanaghy/bbqberry/static

run_deployed:
	$(OUTBIN)/bbqberry --host=0.0.0.0 --port=8888 --static=/home/pi/go/src/github.com/declanshanaghy/bbqberry/static
