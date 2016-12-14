# BBQBerry Makefile

include skel/Makefile

export STUB=yes

unittest_bbqberry: create_influxdb unittest
	@echo "Done"

create_influxdb:
	go run cmd/influxdb/create_database.go

upload_ftp:
	@echo "Uploading via FTP..."
	@time env OUTBIN=$(OUTBIN) APP_NAME=$(APP_NAME) scripts/upload_ftp.sh
	@echo "Upload complete"

upload_scp: build
	@echo "Uploading via SCP..."
	@time scp -p $(OUTBIN)/$(APP_NAME) pi@bbqberry-gaff:~/deploy/bbqberry
	@echo "Upload complete"

run_remote: upload_scp
	ssh pi@bbqberry-gaff ./deploy/bbqberry --host=0.0.0.0 --port=8888 --static=/home/pi/go/src/github.com/declanshanaghy/bbqberry/static
