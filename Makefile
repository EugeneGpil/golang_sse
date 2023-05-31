dev:
	cd container &&\
		docker compose up --build --remove-orphans --detach &&\
		docker compose exec --user=app golang bash

exec-root:
	cd container && docker compose exec --user=root golang bash

stop:
	cd container && docker compose stop
