graph.gen:
	go run github.com/99designs/gqlgen generate

run.lockal.memory:
	go run ./cmd/main.go --container=n --db=memory

run.lockal.postgre:
	go run ./cmd/main.go --container=n 

run.docker.postgre:
	docker compose -f docker-compose.yaml up

run.docker.memory:
	docker compose -f docker-compose.mem.yaml up 

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