tester: *.go
	go get -v
	go test -c -o tester

test: tester
	./tester -test.v
