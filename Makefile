test-fizzbuzz:
	cd fizzbuzz && go test --cover

default:
	go test -v ./...

test:
	go test --cover -v ./...
