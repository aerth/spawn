tester: *.go
	go test -c -o tester

test: tester
	./tester -test.v
