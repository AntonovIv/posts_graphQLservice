graph.gen:
	go run github.com/99designs/gqlgen generate

run:
	go run ./cmd/main.go

build.servicemock:
	mockgen -source=internal\transport\graph\resolver.go -destination=internal\service\posts\mock\service_mock.go

build.repomock:
	mockgen -source=internal\service\posts\service.go -destination=internal\repository\postgre\mock\postgre_mock.go