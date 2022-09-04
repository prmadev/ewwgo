test:
	richgo test . -vet=all -v -benchtime 10s -vet=all -bench=mem -cover
coverage:
	go test -coverprofile cover.out && go tool cover -func=cover.out