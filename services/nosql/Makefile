all: mycloud_nosql

clean:
	rm -f mycloud_nosql

run: mycloud_nosql
	./mycloud_nosql 9000

curl:
	curl --request POST http://localhost:9000/nosql/scan --header 'Content-Type: application/json' --data-raw '{ "table": "PAULOBET", "filter": [ { "attr": "sort", "op": "=", "type": "S", "value": "USER" } ] }'

demo:
	cd /home/ricardo/git/mycloud/demo/nosql; mvn clean install; cd target; java -DURL_NOSQL=http://localhost:9000/nosql/ -cp nosql-0.1.0-SNAPSHOT.jar mycloud.demo.nosql.Scan

run_demo:
	cd /home/ricardo/git/mycloud/demo/nosql/target; java -DURL_NOSQL=http://localhost:9000/nosql/ -cp nosql-0.1.0-SNAPSHOT.jar mycloud.demo.nosql.Scan


mycloud_nosql: nosql.go scan.go create.go model/scan_request.go model/scan_response.go
	go build -o mycloud_nosql nosql.go scan.go create.go
