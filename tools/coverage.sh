#!/bin/bash
#
# Code coverage generation
go test -covermode=count -coverprofile=cover.out ./...


# Display the global code coverage
go tool cover -func=cover.out ;

# If needed, generate HTML report
if [ "$1" == "html" ]; then
    go tool cover -html=cover.out -o coverage.html ;
fi
