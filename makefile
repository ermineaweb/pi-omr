build:
	docker compose build

push:
	docker compose push

all:
	build
	push