all:
	go run main.go

test:
	go test

coverage:
	go test -coverprofile=c.out
	sed -i "s/_$$(pwd|sed 's/\//\\\//g')/./g" c.out
	go tool cover -html=c.out -o=c.html