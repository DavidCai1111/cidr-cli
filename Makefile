test:
	go test -v -race ./cidr

cover:
	rm -rf *.coverprofile
	go test -coverprofile=cidr.coverprofile -v -race ./cidr
	gover
	go tool cover -html=cidr.coverprofile
	rm -rf *.coverprofile
