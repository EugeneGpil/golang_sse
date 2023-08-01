dev:
	docker compose up --build --remove-orphans --detach --force-recreate &&\
	docker compose exec --user=app golang bash

exec:
	docker compose exec golang bash

exec-root:
	docker compose exec --user=root golang bash

stop:
	docker compose stop
