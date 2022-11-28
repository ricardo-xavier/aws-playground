all: mycloud_nosql

clean:
	rm -f mycloud_nosql


mycloud_nosql: nosql.go scan.go model/scan_request.go model/scan_response.go
	go build -o mycloud_nosql nosql.go scan.go 
