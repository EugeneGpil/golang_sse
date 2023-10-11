dev:
	code --folder-uri="vscode-remote://dev-container+$$(pwd | tr -d '\n' | xxd -p)/workspaces/$$(basename "$$(pwd)")"

build:
	docker compose up --build --remove-orphans --detach --force-recreate;
	docker compose stop;

test:
	docker composer run golang go test ./...


# inside container

test:
	go test ./...
