dev:
	code --folder-uri="vscode-remote://dev-container+$$(pwd | tr -d '\n' | xxd -p)/workspaces/$$(basename "$$(pwd)")"

build:
	docker-compose --progress plain build >> "docker/logs/build_$$(date "+%Y-%m-%d_%H:%M:%S").log"

build-no-cache:
	docker-compose --progress plain build --no-cache >> "docker/logs/build_$$(date "+%Y-%m-%d_%H:%M:%S").log"

# inside container
test:
	go test ./...
