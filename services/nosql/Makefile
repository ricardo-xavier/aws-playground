all: mycloud_nosql show_idx

clean:
	rm -f mycloud_nosql show_idx

run: mycloud_nosql
	./mycloud_nosql 9000

curl:
	curl --request POST http://localhost:9000/nosql/scan --header 'Content-Type: application/json' --data-raw '{ "table": "PAULOBET", "filter": [ { "attr": "sort", "op": "=", "type": "S", "value": "USER" } ] }'

demo:
	cd /home/ricardo/git/mycloud/demo/nosql; mvn clean install

run_demo_scan:
	cd /home/ricardo/git/mycloud/demo/nosql/target; java -DURL_NOSQL=http://localhost:9000/nosql/ -cp nosql-0.1.0-SNAPSHOT.jar mycloud.demo.nosql.Scan

run_demo_load:
	cd /home/ricardo/git/mycloud/demo/nosql/target; java -DURL_NOSQL=http://localhost:9000/nosql/ -cp nosql-0.1.0-SNAPSHOT.jar mycloud.demo.nosql.Load


mycloud_nosql: nosql.go scan.go create.go put_item.go model/scan_request.go model/scan_response.go btree/create.go btree/open.go btree/put_item.go btree/close.go btree/find.go btree/page.go btree/split.go
	go build -o mycloud_nosql nosql.go scan.go create.go put_item.go

show_idx: show_idx.go
	go build -o show_idx show_idx.go
