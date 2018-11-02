test:
	go test -v -race

cover:
	rm -rf *.coverprofile
	go test -coverprofile=cidr.coverprofile -v -race
	gover
	go tool cover -html=cidr.coverprofile
	rm -rf *.coverprofile
