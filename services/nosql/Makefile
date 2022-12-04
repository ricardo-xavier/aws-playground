all: mycloud_nosql

clean:
	rm -f mycloud_nosql

run:
	./mycloud_nosql

curl:
	curl --request POST http://localhost:9000/nosql/scan --header 'Content-Type: application/json' --data-raw '{ "table": "PAULOBET", "filter": [ { "attr": "sort", "op": "=", "type": "S", "value": "USER" } ] }'

demo:
	cd /home/ricardo/git/mycloud/demo/nosql; mvn clean install; cd target; java -cp nosql-1.0-SNAPSHOT.jar mycloud.demo.nosql.Scan


mycloud_nosql: nosql.go scan.go model/scan_request.go model/scan_response.go
	go build -o mycloud_nosql nosql.go scan.go 
