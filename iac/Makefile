add: iac

clean:
	rm -f iac

run: iac
	URL_NOSQL=http://localhost:9000/nosql/ ./iac


iac: iac.go
	go build -o iac iac.go
