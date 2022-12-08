all: mycloud

clean:
	rm -f mycloud

run: mycloud
	URL_NOSQL=http://localhost:9000/nosql/ ./mycloud nosql scan --table-name PAULOBET --filter-expression "sort = :sort" --expression-attribute-values "{ \":sort\": { \"S\": \"USER\" } }"

mycloud: mycloud.go nosql.go
	go build -o mycloud mycloud.go nosql.go
