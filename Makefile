graph.gen:
	go run github.com/99designs/gqlgen generate

run.srv.memory:
	go run ./cmd/main.go --container=n --db=memory

run.srv.postgre:
	go run ./cmd/main.go --container=n 

run.all.docker:
	docker compose up

run.db.docker:
	docker compose up db

run.docker.rebuild:
	docker compose up --build posts-app

run.remove.volume:
	docker volume rm posts_graphqlservice_data

build.servicemock:
	mockgen -source=internal\transport\graph\resolver.go -destination=internal\service\posts\mock\service_mock.go

build.repomock:
	mockgen -source=internal\service\posts\service.go -destination=internal\repository\postgre\mock\postgre_mock.go