dev:
	docker compose up --build --remove-orphans --detach --force-recreate;
	code --folder-uri="vscode-remote://dev-container+$$(pwd | tr -d '\n' | xxd -p)/workspaces/$$(basename "$$(pwd)")"

exec:
	docker compose exec golang bash

exec-root:
	docker compose exec --user=root golang bash

stop:
	docker compose stop

test:
	docker composer run golang go test ./...


# inside container

test:
	go test ./...